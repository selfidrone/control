package recognition

import (
	"fmt"
	"time"

	"gocv.io/x/gocv"
)

// Faces implements facial recognition using GoCV
type Faces struct{}

func (f *Faces) Start(deviceID int) {
	processor := NewFaceProcessor()

	// open webcam
	webcam, err := gocv.VideoCaptureDevice(int(deviceID))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer webcam.Close()

	// prepare image matrix
	img := gocv.NewMat()
	defer img.Close()

	fmt.Printf("start reading camera device: %v\n", deviceID)
	for {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("cannot read device %d\n", deviceID)
			return
		}
		if img.Empty() {
			continue
		}

		st := time.Now()

		fmt.Println("Got image")
		f, b := processor.DetectFaces(img)
		fmt.Printf("faces: %#v, bounds: %#v", f, b)

		fmt.Println("  ")
		fmt.Println("Duration:", time.Now().Sub(st))

		time.Sleep(100 * time.Millisecond)
	}

}

func (f *Faces) Stop() {

}
