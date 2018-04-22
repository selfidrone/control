package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"

	hclog "github.com/hashicorp/go-hclog"
	stan "github.com/nats-io/go-nats-streaming"
	"github.com/selfidrone/control/control"
	"github.com/selfidrone/control/recognition"
	messages "github.com/selfidrone/messages"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/ble"
	"gobot.io/x/gobot/platforms/parrot/minidrone"
)

var (
	robot     *gobot.Robot
	drone     control.MamboDrone
	autoPilot *control.AutoPilot
	nc        stan.Conn
	cam       control.Camera
)

var (
	imageSize  = flag.Int("size", 0, "recording image size")
	exposure   = flag.String("exposure", "auto", "exposure mode")
	height     = flag.Int("height", 600, "height")
	width      = flag.Int("width", 800, "width")
	natsServer = flag.String("nats", "nats://localhost:4222", "location of the nats.io server")
	simulate   = flag.Bool("simulate", false, "Simulate sending commands to drone")
	logLevel   = flag.String("log_level", "INFO", "log level should be set to a vault INFO, WARN, DEBUG, TRACE")
	camera     = flag.Int("camera", 0, "device id for the camera to use")
)

var latestImage []byte
var imageMutex sync.Mutex
var messageExpiration = (1000 * time.Millisecond)
var log hclog.Logger

func main() {
	flag.Parse()

	log = hclog.New(&hclog.LoggerOptions{Level: hclog.LevelFromString(*logLevel)})
	log.Info("Starting drone control", "log_level", *logLevel)

	setupNATS()
	//go initDroneCamera()
	go startGoCV()

	handleExit()
}

func handleExit() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func startGoCV() {
	faces := recognition.Faces{}
	faces.Start(*camera)
}

func initDroneCamera() {
	log.Info("Init Camera")

	// create a new camera and update every second
	cam = control.NewCamera(500*time.Millisecond, *height, *width, *imageSize, *exposure)

	for i := range cam.Images() {
		imageMutex.Lock()
		latestImage = i
		imageMutex.Unlock()

		di := messages.DroneImage{}
		di.Data = base64.StdEncoding.EncodeToString(i)

		log.Debug("New image from camera")
		nc.Publish(messages.MessageDroneImage, di.EncodeMessage())
	}
}

func sendLatestPicture() {
	log.Info("Send latest camera image")
	imageMutex.Lock()

	di := messages.DroneImage{}
	di.Data = base64.StdEncoding.EncodeToString(latestImage)
	imageMutex.Unlock()
	nc.Publish(messages.MessageDronePicture, di.EncodeMessage())
}

func setupNATS() {
	var err error
	clientID := fmt.Sprintf("server-%d", time.Now().UnixNano())
	nc, err = stan.Connect("test-cluster", clientID, stan.NatsURL(*natsServer))
	if err != nil {
		log.Error("Unable to connect to nats server: ", err)
		os.Exit(1)
	}

	_, err = nc.Subscribe(messages.MessageFlight, handleMessage)
	if err != nil {
		log.Error("Unable to subscribe to queue", "error", err)
		os.Exit(1)
	}

	_, err = nc.Subscribe(messages.MessageFaceDetection, handleFaceDetection)
	if err != nil {
		log.Error("Unable to subscribe to queue", "error", err)
		os.Exit(1)
	}
}

func handleMessage(m *stan.Msg) {
	if isExpiredMessage(m) {
		log.Debug("message expired", "message", m.String())
		return
	}

	msg := messages.Flight{}
	msg.DecodeMessage(m.Data)

	log.Debug("Message received", "msg", fmt.Sprintf("%#v", msg))

	switch msg.Command {
	case messages.CommandConnect:
		log.Info("Attempting to connect to drone")
		go startCamera()
		go startDrone("Mambo_586378")
	case messages.CommandDisconnect:
		log.Info("Attempting to disconnect to drone")
		stopDrone()
		stopCamera()
	case messages.CommandTakePicture:
		log.Info("Handle take picture message")
		go func() {
			sendLatestPicture()
		}()
	default:
		go func(msg messages.Flight) {
			log.Info("Handle flight message")
			autoPilot.HandleMessage(&msg)
		}(msg)
	}
}

func handleFaceDetection(m *stan.Msg) {
	if isExpiredMessage(m) {
		log.Debug("message expired", m.String())
		return
	}

	fm := messages.FaceDetected{}
	fm.DecodeMessage(m.Data)

	autoPilot.HandleMessage(&fm)
}

func isExpiredMessage(m *stan.Msg) bool {
	return (time.Now().Sub(time.Unix(0, m.Timestamp)) > messageExpiration)
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
		log.Info("Drone Ready...")
	}

	if *simulate {
		log.Info("Running in simulated mode")

		drone = control.NewSimulatedDrone(log.Named("SimulatedDrone"))
		autoPilot = control.NewAutoPilot(drone)

		work()
	} else {
		log.Info("Connecting to minidrone", "name", name)

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
			log.Error("Error stopping drone", "err", r)
		}
	}()

	robot.Stop()
}

func writeLog(message string, args ...interface{}) {
	log.Info(message, args)
	nc.Publish("log", []byte(fmt.Sprintf(message, args)))
}
