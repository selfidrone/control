package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/nats-io/nats"
	"github.com/nicholasjackson/drone-control/control"
	messages "github.com/nicholasjackson/drone-messages"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/ble"
	"gobot.io/x/gobot/platforms/parrot/minidrone"
)

var robot *gobot.Robot
var drone *minidrone.Driver
var autoPilot *control.AutoPilot
var nc *nats.Conn
var cam control.Camera

var imageSize = flag.Int("size", 0, "recording image size")
var exposure = flag.String("exposure", "auto", "exposure mode")
var height = flag.Int("height", 600, "height")
var width = flag.Int("width", 800, "width")
var natsServer = flag.String("nats", "nats://localhost:4222", "location of the nats.io server")

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
	cam = control.NewCamera(500*time.Millisecond, *height, *width, *imageSize, *exposure)

	for i := range cam.Images() {
		di := messages.DroneImage{}
		di.SetZippedData(i)

		nc.Publish(messages.MessageDroneImage, di.EncodeMessage())
	}
}

func startCamera() {
	cam.Start()
}

func stopCamera() {
	cam.Stop()
}

func startDrone(name string) {
	bleAdaptor := ble.NewClientAdaptor(name)
	drone = minidrone.NewDriver(bleAdaptor)

	autoPilot = control.NewAutoPilot(drone)

	work := func() {
		autoPilot.Setup()
	}

	robot = gobot.NewRobot("minidrone",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{drone},
		work,
	)

	robot.Start()
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
	nc, err = nats.Connect(*natsServer)
	if err != nil {
		log.Fatal("Unable to connect to nats server")
	}

	nc.Subscribe(messages.MessageFlight, func(m *nats.Msg) {
		msg := messages.Flight{}
		msg.DecodeMessage(m.Data)

		switch msg.Command {
		case messages.CommandConnect:
			log.Println("Attempting to connect to drone")
			//go startCamera()
			go startDrone("Mambo_586378")
		case messages.CommandDisconnect:
			stopDrone()
			//stopCamera()
		default:
			go func(msg messages.Flight) {
				autoPilot.HandleMessage(&msg)
			}(msg)
		}
	})

	nc.Subscribe(messages.MessageFaceDetection, func(m *nats.Msg) {
		fm := messages.FaceDetected{}
		fm.DecodeMessage(m.Data)

		autoPilot.HandleMessage(&fm)
	})
}

func writeLog(message string, args ...interface{}) {
	log.Println(message, args)
	nc.Publish("log", []byte(fmt.Sprintf(message, args)))
}
