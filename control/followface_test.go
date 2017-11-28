package control

import (
	"image"
	"testing"
	"time"

	"github.com/matryer/is"
	messages "github.com/nicholasjackson/drone-messages"
)

var speed = 30

var initialMessage = messages.FaceDetected{
	Faces: []image.Rectangle{
		image.Rect(200, 250, 300, 350),
	},
	Bounds: bounds,
}

var bounds = image.Rect(0, 0, 800, 600)

var faces = [][]image.Rectangle{
	[]image.Rectangle{
		image.Rect(200, 250, 300, 350),
	},
	[]image.Rectangle{
		image.Rect(250, 250, 350, 350),
	},
	[]image.Rectangle{
		image.Rect(300, 250, 400, 350),
	},
	[]image.Rectangle{
		image.Rect(350, 250, 450, 350),
	},
	[]image.Rectangle{
		image.Rect(350, 250, 450, 350),
	},
	//		image.Rect(0, 0, 500, 400),
}

func testNoMovement(is *is.I, md *MamboDroneMock) {
	is.Equal(0, len(md.BackwardCalls()))         // expected no backward calls
	is.Equal(0, len(md.ClockwiseCalls()))        // expected no clockwise calls
	is.Equal(0, len(md.CounterClockwiseCalls())) // expected no counter clockwise calls
	is.Equal(0, len(md.DownCalls()))             // expected no down calls
	is.Equal(0, len(md.ForwardCalls()))          // expected no forward calls
	is.Equal(0, len(md.LandCalls()))             // expected no land calls
	is.Equal(0, len(md.LeftCalls()))             // expected no left calls
	is.Equal(0, len(md.RightCalls()))            // expected no right calls
	is.Equal(0, len(md.TakeOffCalls()))          // expecte no take off calls
	is.Equal(0, len(md.UpCalls()))               // expected no up calls
}

func TestStopFollowingStopsDrone(t *testing.T) {
	ap, md, _ := setupAutopilot(t)
	is := is.New(t)

	ap.StopFollowing()

	is.Equal(1, len(md.StopCalls())) // expeced stop to be called
}

func TestFollowFaceDoesNothingOnFirstCall(t *testing.T) {
	ap, md, _ := setupAutopilot(t)
	is := is.New(t)

	ap.HandleMessage(&messages.FaceDetected{})

	testNoMovement(is, md)
}

func TestOnlyMovesWhenFollowingEnabled(t *testing.T) {
	ap, md, _ := setupAutopilot(t)
	is := is.New(t)

	ap.StopFollowing()
	ap.HandleMessage(&initialMessage)
	ap.HandleMessage(&messages.FaceDetected{
		Faces: []image.Rectangle{
			image.Rect(0, 0, 400, 300),
		},
		Bounds: bounds,
	})

	testNoMovement(is, md)
}

func TestMovesDroneLeft(t *testing.T) {
	ap, md, _ := setupAutopilot(t)
	is := is.New(t)

	ap.HandleMessage(&initialMessage)
	ap.HandleMessage(&messages.FaceDetected{
		Faces:  faces[0],
		Bounds: bounds,
	})

	is.Equal(1, len(md.LeftCalls()))       // expected 1 call to move left
	is.Equal(speed, md.LeftCalls()[0].Val) // expected drone to move left at speed
}

func TestDoesNotMovesDroneLeftWhenMinDistance(t *testing.T) {
	ap, md, _ := setupAutopilot(t)
	is := is.New(t)

	ap.HandleMessage(&initialMessage)
	ap.HandleMessage(&messages.FaceDetected{
		Faces:  faces[4],
		Bounds: bounds,
	})

	is.Equal(1, len(md.StopCalls())) // expected 1 call to stop
}

/*
func TestMovesDroneRight(t *testing.T) {
	ap, md, _ := setupAutopilot(t)
	is := is.New(t)

	ap.HandleMessage(initialMessage)
	ap.HandleMessage(messages.FaceDetected{
		Faces: []image.Rectangle{
			image.Rect(0, 0, 200, 300),
		},
		Bounds: bounds,
	})

	is.Equal(1, len(md.RightCalls()))       // expected 1 call to move right
	is.Equal(speed, md.RightCalls()[0].Val) // expected drone to move right at speed
}

func TestDoesNotMovesDroneRightWhenMinDistance(t *testing.T) {
	ap, md, _ := setupAutopilot(t)
	is := is.New(t)

	ap.HandleMessage(initialMessage)
	ap.HandleMessage(messages.FaceDetected{
		Faces: []image.Rectangle{
			image.Rect(0, 0, 299, 300),
		},
		Bounds: bounds,
	})

	is.Equal(1, len(md.RightCalls()))   // expected 1 call to move right
	is.Equal(0, md.RightCalls()[0].Val) // expected to set speed to 0
}
*/

func TestStopsDroneAfterNSecondsAndNoFaceData(t *testing.T) {
	ap, md, _ := setupAutopilot(t)
	is := is.New(t)

	ap.HandleMessage(&initialMessage)
	ap.HandleMessage(&messages.FaceDetected{
		Faces:  faces[0],
		Bounds: bounds,
	})

	time.Sleep(2500 * time.Millisecond)

	is.Equal(1, len(md.LeftCalls()))       // expected 1 call to move right
	is.Equal(speed, md.LeftCalls()[0].Val) // expected drone to move right at speed
	is.Equal(1, len(md.StopCalls()))       // expected 1 call to stop
}

/*
func TestDMSResetsAfterDelayAndNewFaceData(t *testing.T) {
	ap, md, _ := setupAutopilot(t)
	is := is.New(t)

	ap.HandleMessage(initialMessage)
	ap.HandleMessage(messages.FaceDetected{
		Faces: []image.Rectangle{
			image.Rect(0, 0, 200, 300),
		},
		Bounds: bounds,
	})

	time.Sleep(1500 * time.Millisecond)

	ap.HandleMessage(messages.FaceDetected{
		Faces: []image.Rectangle{
			image.Rect(0, 0, 100, 300),
		},
		Bounds: bounds,
	})

	time.Sleep(1500 * time.Millisecond)

	ap.HandleMessage(messages.FaceDetected{
		Faces: []image.Rectangle{
			image.Rect(0, 0, 50, 300),
		},
		Bounds: bounds,
	})

	is.Equal(3, len(md.RightCalls()))       // expected 3 call to move right
	is.Equal(speed, md.RightCalls()[0].Val) // expected drone to move right at speed
	is.Equal(0, len(md.StopCalls()))        // expected 0 call to stop
}
*/
