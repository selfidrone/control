package messages

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"image"
	"io"
	"os"
)

const (
	// MessageFlight is the name of a flight message for the drone
	MessageFlight = "drone.flight"
	// MessageFaceDetection is the name of a message when a new face has been detected
	MessageFaceDetection = "image.facedetection"
	// MessageDroneImage is the name of a messeage when the drone takes a new image
	MessageDroneImage = "image.stream"
	// MessageDronePicture is the name of a messeage when the drone takes a new image
	MessageDronePicture = "image.picture"
)

const (
	// CommandBackward instructs a drone to move backward
	CommandBackward = "backward"
	//CommandClockwise instructs a drone to move clockwise
	CommandClockwise = "clockwise"
	// CommandCounterClockwise instructs a drone to move counter clockwise
	CommandCounterClockwise = "counterclockwise"
	// CommandDown instructs a drone to move down
	CommandDown = "down"
	// CommandForward instructs a drone to move forward
	CommandForward = "forward"
	// CommandLand instructs a drone to land
	CommandLand = "land"
	// CommandLeft instructs a drone to move left
	CommandLeft = "left"
	// CommandRight instructs a drone to move right
	CommandRight = "right"
	// CommandTakeOff instructs a drone taking off
	CommandTakeOff = "takeoff"
	// CommandUp instructs a drone to move up
	CommandUp = "up"
	// CommandStop instructs the drone to stop
	CommandStop = "stop"
	//CommandConnect instructs the application to connect to a drone
	CommandConnect = "connect"
	//CommandDisconnect instructs the application to connect to a drone
	CommandDisconnect = "disconnect"
	// CommandTakePicture isntructs the drone to take a picture
	CommandTakePicture = "takepicture"
	// CommandFollowFace instructs the drone to follow the location of a detected face
	CommandFollowFace = "followface"
)

// DroneImage defines a new image taken from a drone
type DroneImage struct {
	Data string // base64 encoded image
}

// Flight defines a flight instruction message
type Flight struct {
	Command string
	Value   int
}

// FaceDetected defines a face detection message
type FaceDetected struct {
	Faces  []image.Rectangle
	Bounds image.Rectangle
}

// EncodeMessage gob encodes the message and returns a byte slice
func (bm *Flight) EncodeMessage() []byte {
	data, err := json.Marshal(bm)
	if err != nil {
		return nil
	}

	return data
}

// DecodeMessage decodes the messgage from gob byte slice
func (bm *Flight) DecodeMessage(data []byte) {
	json.Unmarshal(data, bm)
}

// EncodeMessage gob encodes the message and returns a byte slice
func (bm *FaceDetected) EncodeMessage() []byte {
	data, err := json.Marshal(bm)
	if err != nil {
		return nil
	}

	return data
}

// DecodeMessage decodes the messgage from gob byte slice
func (bm *FaceDetected) DecodeMessage(data []byte) {
	json.Unmarshal(data, bm)
}

// EncodeMessage gob encodes the message and returns a byte slice
func (bm *DroneImage) EncodeMessage() []byte {
	data, err := json.Marshal(bm)
	if err != nil {
		return nil
	}

	return data
}

// DecodeMessage decodes the messgage from gob byte slice
func (bm *DroneImage) DecodeMessage(data []byte) {
	json.Unmarshal(data, bm)
}

// SaveDataToFile uncompresses the Data field and saves to a file
func (bm *DroneImage) SaveDataToFile(filename string) error {
	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		os.Remove(filename)
	}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	d, err := base64.StdEncoding.DecodeString(bm.Data)
	if err != nil {
		return err
	}

	io.Copy(f, bytes.NewBuffer(d))

	return nil
}
