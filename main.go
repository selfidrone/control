package main

import (
	"bytes"
	"compress/gzip"
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/nats-io/nats"
	"github.com/nicholasjackson/drone-control/control"
	messages "github.com/nicholasjackson/drone-messages"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/ble"
	"gobot.io/x/gobot/platforms/parrot/minidrone"
)

var drone *minidrone.Driver
var nc *nats.Conn

func main() {
	log.Println("Starting drone control")

	setupNATS()

	// create a new camera and update every second
	cam := control.NewCamera(1 * time.Second)
	go cam.Start()

	for i := range cam.Images() {
		var zb bytes.Buffer
		zw, _ := gzip.NewWriterLevel(&zb, gzip.BestCompression)
		zw.Write(i)
		zw.Close()

		m := messages.DroneImage{Data: zb.Bytes()}

		var b bytes.Buffer
		gob.NewEncoder(&b).Encode(m)

		nc.Publish(messages.MessageDroneImage, b.Bytes())
		log.Println("Got image")
	}
}

func startDrone() {

	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	drone = minidrone.NewDriver(bleAdaptor)

	ap := control.NewAutoPilot(drone)

	work := func() {
		ap.Setup()
		nc.Subscribe("done.control", func(m *nats.Msg) {
			ap.HandleMessage(string(m.Data))
		})
	}

	robot := gobot.NewRobot("minidrone",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{drone},
		work,
	)

	robot.Start()
}

func setupNATS() {
	var err error
	nc, err = nats.Connect("nats://192.168.1.113:4222")
	if err != nil {
		log.Fatal("Unable to connect to nats server")
	}
}

func writeLog(message string, args ...interface{}) {
	log.Println(message, args)
	nc.Publish("log", []byte(fmt.Sprintf(message, args)))
}
