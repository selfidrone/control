package main

import (
	"flag"
	"image"
	"log"
	"time"

	"github.com/nats-io/nats"
	messages "github.com/nicholasjackson/drone-messages"
)

var natsServer = flag.String("nats", "nats://localhost:4222", "location of the nats.io server")
var nc *nats.Conn

var bounds = image.Rect(0, 0, 800, 600)

// centerPoint 400
// moveLeft < 350
// moveLeft > 450

var m = [][]image.Rectangle{
	[]image.Rectangle{
		image.Rect(200, 250, 250, 300), // Initial Pos CP: 225
	},
	[]image.Rectangle{
		image.Rect(200, 250, 300, 300), // Move Left CP: 225
	},
	[]image.Rectangle{
		image.Rect(250, 250, 300, 300), // Move Left CP: 275
	},
	[]image.Rectangle{
		image.Rect(350, 250, 400, 300), // Stop CP: 375
	},
	[]image.Rectangle{
		image.Rect(350, 250, 400, 300), // Stop CP: 375
	},
	[]image.Rectangle{
		image.Rect(500, 250, 550, 300), // Right CP: 525
	},
	[]image.Rectangle{
		image.Rect(450, 250, 500, 300), // Right CP: 475
	},
	[]image.Rectangle{
		image.Rect(400, 250, 450, 300), // Stop CP: 425
	},
}

// This applicaton is used for testing the autopilot behaviour without needing to
// run face detection.
func main() {
	flag.Parse()

	time.Sleep(5 * time.Second)

	var err error
	nc, err = nats.Connect(*natsServer)
	if err != nil {
		log.Fatal("Unable to connect to nats server")
	}

	for _, f := range m {
		log.Println("Message:", f)

		fdm := messages.FaceDetected{
			Faces:  f,
			Bounds: bounds,
		}

		nc.Publish(messages.MessageFaceDetection, fdm.EncodeMessage())
		time.Sleep(1 * time.Second)
	}
}
