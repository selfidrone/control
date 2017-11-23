package control

import (
	"testing"

	"github.com/matryer/is"
	messages "github.com/nicholasjackson/drone-messages"
)

var bounds = messages.Rectangle{X: 0, Y: 0, Width: 1000, Height: 800}

var initialMessage = messages.FaceDetected{
	Faces: []messages.Rectangle{
		{
			X:      300,
			Y:      700,
			Width:  200,
			Height: 200,
		},
	},
	Bounds: bounds,
}

func TestFollowFaceDoesNothingOnFirstCall(t *testing.T) {
	ap, md := setupAutopilot(t)
	is := is.New(t)

	ap.HandleMessage(messages.FaceDetected{})

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

func TestMovesDroneToCorrectPosition(t *testing.T) {
	ap, md := setupAutopilot(t)
	is := is.New(t)

	ap.HandleMessage(initialMessage)
	ap.HandleMessage(messages.FaceDetected{
		Faces: []messages.Rectangle{
			{
				X:      500,
				Y:      700,
				Width:  200,
				Height: 200,
			},
		},
		Bounds: bounds,
	})

	lc := md.LeftCalls()[0]
	is.Equal(1, lc.Val) // expected drone to move left by 10
}
