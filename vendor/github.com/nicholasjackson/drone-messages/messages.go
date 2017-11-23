package messages

const (
	// MessageFlight is the name of a flight message for the drone
	MessageFlight = "drone.flight"
	// MessageFaceDetection is the name of a message when a new face has been detected
	MessageFaceDetection = "image.facedetection"
	// MessageDroneImage is the name of a messeage when the drone takes a new image
	MessageDroneImage = "image.new"
)

const (
	// CommandTakeOff instructs a drone taking off
	CommandTakeOff = "takeoff"
	// CommandLand instructs a drone to land
	CommandLand = "land"
)

// DroneImage defines a new image taken from a drone
type DroneImage struct {
	Data []byte
}

// Flight defines a flight instruction message
type Flight struct {
	Command string
	Value   int
}

// FaceDetected defines a face detection message
type FaceDetected struct {
	Faces  []Rectangle
	Bounds Rectangle
}

// Rectangle defines a rectangular shape
type Rectangle struct {
	X      int
	Y      int
	Height int
	Width  int
}
