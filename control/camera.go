package control

import (
	"log"
	"os"
	"time"

	"github.com/technomancers/piCamera"
)

//go:generate moq -out camera_mocks_test.go . Camera

// Camera defines a new control interface
type Camera interface {
	Start()
	Stop()
	Images() chan []byte
}

// CameraImpl defines the raspberry pi camera
type CameraImpl struct {
	cam       *piCamera.PiCamera
	imageChan chan []byte
	update    time.Duration
	running   bool
}

// NewCamera creates a new camera with the given update duration
func NewCamera(update time.Duration, height, width int, imageSize int, exposure string) *CameraImpl {
	c := &CameraImpl{update: update}
	c.imageChan = make(chan []byte)

	args := piCamera.NewArgs()
	args.Mode = imageSize
	//	args.Rotation = 180
	args.Height = height
	args.Width = width

	switch exposure {
	case "auto":
		args.ExposureMode = piCamera.ExpAuto
	case "night":
		args.ExposureMode = piCamera.ExpNight
	case "backlight":
		args.ExposureMode = piCamera.ExpBacklight
		args.Metering = piCamera.MeterBacklit
	case "spotlight":
		args.ExposureMode = piCamera.ExpSpotlight
	}

	var err error
	c.cam, err = piCamera.New(nil, args)
	if err != nil {
		log.Fatal(err)
	}

	return c
}

// Start starts the camera
func (c *CameraImpl) Start() {
	log.Println("Starting Camera")
	c.running = true

	err := c.cam.Start()
	if err != nil {
		log.Fatal(err)
	}

	for c.running == true {
		buffer, err := c.cam.GetFrame()

		if err != nil {
			log.Println("Get Frame", err)
		} else {
			//c.writeImage(buffer)
			c.imageChan <- buffer
		}

		time.Sleep(c.update)
	}
}

// Images returns a channel containing jpg images
func (c *CameraImpl) Images() chan []byte {
	return c.imageChan
}

// Stop stops capturing images
func (c *CameraImpl) Stop() {
	log.Println("Stopping Camera")

	c.running = false
	c.cam.Stop()
}

func (c *CameraImpl) writeImage(buffer []byte) {
	if _, err := os.Stat("/tmp/latest.jpg"); !os.IsNotExist(err) {
		os.Remove("/tmp/latest.jpg")
	}

	f, err := os.Create("/tmp/latest.jpg")
	if err != nil {
		log.Println("Unable to create file", f)
		return
	}
	defer f.Close()

	f.Write(buffer)

	log.Println("Written image to /tmp/latest.jpg")
}
