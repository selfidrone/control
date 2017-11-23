package control

import (
	"log"
	"os"
	"time"

	"github.com/technomancers/piCamera"
)

// Camera defines the raspberry pi camera
type Camera struct {
	cam       *piCamera.PiCamera
	imageChan chan []byte
	update    time.Duration
}

// NewCamera creates a new camera with the given update duration
func NewCamera(update time.Duration) *Camera {
	c := &Camera{update: update}
	c.imageChan = make(chan []byte)

	args := piCamera.NewArgs()
	args.Mode = 0
	args.Rotation = 180
	args.ExposureMode = piCamera.ExpBacklight

	var err error
	c.cam, err = piCamera.New(nil, args)
	if err != nil {
		log.Fatal(err)
	}

	return c
}

// Start starts the camera
func (c *Camera) Start() {
	err := c.cam.Start()
	if err != nil {
		log.Fatal(err)
	}

	for {
		buffer, err := c.cam.GetFrame()

		if err != nil {
			log.Println("Get Frame", err)
		} else {
			c.writeImage(buffer)
			c.imageChan <- buffer
		}

		time.Sleep(c.update)
	}
}

// Images returns a channel containing jpg images
func (c *Camera) Images() chan []byte {
	return c.imageChan
}

// Stop stops capturing images
func (c *Camera) Stop() {
	c.cam.Stop()
	close(c.imageChan)
}

func (c *Camera) writeImage(buffer []byte) {
	if _, err := os.Stat("./latest.jpg"); !os.IsNotExist(err) {
		os.Remove("./latest.jpg")
	}

	f, err := os.Create("./latest.jpg")
	if err != nil {
		log.Println("Unable to create file", f)
		return
	}
	defer f.Close()

	f.Write(buffer)

	log.Println("Written image to ./latest.jpg")
}
