package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"

	stan "github.com/nats-io/go-nats-streaming"
	"github.com/nicholasjackson/drone-control/control"
	messages "github.com/nicholasjackson/drone-messages"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/ble"
	"gobot.io/x/gobot/platforms/parrot/minidrone"
)

var robot *gobot.Robot
var drone control.MamboDrone
var autoPilot *control.AutoPilot
var nc stan.Conn
var cam control.Camera

var imageSize = flag.Int("size", 0, "recording image size")
var exposure = flag.String("exposure", "auto", "exposure mode")
var height = flag.Int("height", 600, "height")
var width = flag.Int("width", 800, "width")
var natsServer = flag.String("nats", "nats://localhost:4222", "location of the nats.io server")
var simulate = flag.Bool("simulate", false, "Simulate sending commands to drone")
var latestImage []byte
var imageMutex sync.Mutex

func main() {
	log.Println("Starting drone control")
	flag.Parse()

	setupNATS()
	go initDroneCamera()

	handleExit()
}

func handleExit() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func initDroneCamera() {
	fmt.Println("Init Camera")

	// create a new camera and update every second
	cam = control.NewCamera(1000*time.Millisecond, *height, *width, *imageSize, *exposure)

	for i := range cam.Images() {
		imageMutex.Lock()
		latestImage = i
		imageMutex.Unlock()

		di := messages.DroneImage{}
		di.Data = base64.StdEncoding.EncodeToString(i)

		nc.Publish(messages.MessageDroneImage, di.EncodeMessage())
	}
}

func sendLatestPicture() {
	log.Println("Taking picture")
	imageMutex.Lock()

	di := messages.DroneImage{}
	di.Data = base64.StdEncoding.EncodeToString(latestImage)
	imageMutex.Unlock()
	nc.Publish(messages.MessageDronePicture, di.EncodeMessage())
}

func startCamera() {
	cam.Start()
}

func stopCamera() {
	cam.Stop()
}

func startDrone(name string) {
	work := func() {
		autoPilot.Setup()
		fmt.Println("Ready...")
	}

	if *simulate {
		drone = &control.SimulatedDrone{}
		autoPilot = control.NewAutoPilot(drone)

		work()
	} else {
		bleAdaptor := ble.NewClientAdaptor(name)
		drone = minidrone.NewDriver(bleAdaptor)
		autoPilot = control.NewAutoPilot(drone)

		robot = gobot.NewRobot("minidrone",
			[]gobot.Connection{bleAdaptor},
			[]gobot.Device{drone.(*minidrone.Driver)},
			work,
		)
		robot.Start()
	}

}

func stopDrone() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error stopping drone", r)
		}
	}()

	robot.Stop()
}

func setupNATS() {
	var err error
	clientID := fmt.Sprintf("server-%d", time.Now().UnixNano())
	nc, err = stan.Connect("test-cluster", clientID, stan.NatsURL(*natsServer))
	if err != nil {
		log.Fatal("Unable to connect to nats server: ", err)
	}

	nc.Subscribe(messages.MessageFlight, func(m *stan.Msg) {
		msg := messages.Flight{}
		msg.DecodeMessage(m.Data)

		switch msg.Command {
		case messages.CommandConnect:
			log.Println("Attempting to connect to drone")
			go startCamera()
			go startDrone("Mambo_586378")
		case messages.CommandDisconnect:
			stopDrone()
			stopCamera()
		case messages.CommandTakePicture:
			go func() {
				sendLatestPicture()
			}()
		default:
			go func(msg messages.Flight) {
				autoPilot.HandleMessage(&msg)
			}(msg)
		}
	})

	nc.Subscribe(messages.MessageFaceDetection, func(m *stan.Msg) {
		fm := messages.FaceDetected{}
		fm.DecodeMessage(m.Data)

		autoPilot.HandleMessage(&fm)
	})
}

func writeLog(message string, args ...interface{}) {
	log.Println(message, args)
	nc.Publish("log", []byte(fmt.Sprintf(message, args)))
}
