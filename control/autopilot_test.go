package control

import (
	"testing"

	"github.com/matryer/is"
	messages "github.com/nicholasjackson/drone-messages"
	"gobot.io/x/gobot/platforms/parrot/minidrone"
)

func setupAutopilot(t *testing.T) (*AutoPilot, *MamboDroneMock, *CameraMock) {
	mockedMamboDrone := &MamboDroneMock{
		BackwardFunc: func(val int) error {
			return nil
		},
		ClockwiseFunc: func(val int) error {
			return nil
		},
		CounterClockwiseFunc: func(val int) error {
			return nil
		},
		DownFunc: func(val int) error {
			return nil
		},
		ForwardFunc: func(val int) error {
			return nil
		},
		LandFunc: func() error {
			return nil
		},
		LeftFunc: func(val int) error {
			return nil
		},
		OnFunc: func(name string, f func(s interface{})) error {
			return nil
		},
		RightFunc: func(val int) error {
			return nil
		},
		TakeOffFunc: func() error {
			return nil
		},
		UpFunc: func(val int) error {
			return nil
		},
		StopFunc: func() error {
			return nil
		},
	}

	mockedCamera := &CameraMock{
		ImagesFunc: func() chan []byte {
			return nil
		},
		StartFunc: func() {
		},
		StopFunc: func() {
		},
	}

	ap := NewAutoPilot(mockedMamboDrone)
	ap.StartFollowing()

	return ap, mockedMamboDrone, mockedCamera
}

func TestSetupAddsBatteryCallback(t *testing.T) {
	ap, md, _ := setupAutopilot(t)
	is := is.New(t)

	ap.Setup()

	c := fetchCallsForCallback(md, minidrone.Battery)

	is.Equal(len(c), 1) // expected setup to add battery callback
}

func TestSetupAddsFlightStatusCallback(t *testing.T) {
	ap, md, _ := setupAutopilot(t)
	is := is.New(t)

	ap.Setup()

	c := fetchCallsForCallback(md, minidrone.FlightStatus)

	is.Equal(len(c), 1) // expected setup to add flight status callback
}

func TestSetupAddsTakeoffCallback(t *testing.T) {
	ap, md, _ := setupAutopilot(t)
	is := is.New(t)

	ap.Setup()

	c := fetchCallsForCallback(md, minidrone.Takeoff)

	is.Equal(len(c), 1) // expected setup to add takeoff status callback
}

func TestSetupAddsHoveringCallback(t *testing.T) {
	ap, md, _ := setupAutopilot(t)
	is := is.New(t)

	ap.Setup()

	c := fetchCallsForCallback(md, minidrone.Hovering)

	is.Equal(len(c), 1) // expected setup to add hovering status callback
}

func TestSetupAddsLandingCallback(t *testing.T) {
	ap, md, _ := setupAutopilot(t)
	is := is.New(t)

	ap.Setup()

	c := fetchCallsForCallback(md, minidrone.Landing)

	is.Equal(len(c), 1) // expected setup to add landing status callback
}

func TestSetupAddsLandedCallback(t *testing.T) {
	ap, md, _ := setupAutopilot(t)
	is := is.New(t)

	ap.Setup()

	c := fetchCallsForCallback(md, minidrone.Landed)

	is.Equal(len(c), 1) // expected setup to add landed status callback
}

func TestTakeoffMessageSendsCommandToDrone(t *testing.T) {
	ap, md, _ := setupAutopilot(t)
	is := is.New(t)

	ap.HandleMessage(&messages.Flight{Command: messages.CommandTakeOff})

	is.Equal(len(md.TakeOffCalls()), 1) // should have called takeoff once
}

func TestLandMessageSendsCommandToDrone(t *testing.T) {
	ap, md, _ := setupAutopilot(t)
	is := is.New(t)

	ap.HandleMessage(&messages.Flight{Command: messages.CommandLand})

	is.Equal(len(md.LandCalls()), 1) // should have called takeoff once
}

func TestUpMessageSendsCommandToDrone(t *testing.T) {
	ap, md, _ := setupAutopilot(t)
	is := is.New(t)

	ap.HandleMessage(&messages.Flight{Command: messages.CommandUp, Value: 1})

	is.Equal(1, len(md.UpCalls()))  // should have called up once
	is.Equal(1, md.calls.Up[0].Val) // should have used an up value of 1
}

func TestDownMessageSendsCommandToDrone(t *testing.T) {
	ap, md, _ := setupAutopilot(t)
	is := is.New(t)

	ap.HandleMessage(&messages.Flight{Command: messages.CommandDown, Value: 1})

	is.Equal(1, len(md.DownCalls()))  // should have called down once
	is.Equal(1, md.calls.Down[0].Val) // should have used an down value of 1
}

func TestLeftMessageSendsCommandToDrone(t *testing.T) {
	ap, md, _ := setupAutopilot(t)
	is := is.New(t)

	ap.HandleMessage(&messages.Flight{Command: messages.CommandLeft, Value: 1})

	is.Equal(1, len(md.LeftCalls()))  // should have called left once
	is.Equal(1, md.calls.Left[0].Val) // should have used an left value of 1
}

func TestRightMessageSendsCommandToDrone(t *testing.T) {
	ap, md, _ := setupAutopilot(t)
	is := is.New(t)

	ap.HandleMessage(&messages.Flight{Command: messages.CommandRight, Value: 1})

	is.Equal(1, len(md.RightCalls()))  // should have called right once
	is.Equal(1, md.calls.Right[0].Val) // should have used an right value of 1
}

func TestForwardMessageSendsCommandToDrone(t *testing.T) {
	ap, md, _ := setupAutopilot(t)
	is := is.New(t)

	ap.HandleMessage(&messages.Flight{Command: messages.CommandForward, Value: 1})

	is.Equal(1, len(md.ForwardCalls()))  // should have called forward once
	is.Equal(1, md.calls.Forward[0].Val) // should have used an forward value of 1
}

func TestBackwardMessageSendsCommandToDrone(t *testing.T) {
	ap, md, _ := setupAutopilot(t)
	is := is.New(t)

	ap.HandleMessage(&messages.Flight{Command: messages.CommandBackward, Value: 1})

	is.Equal(1, len(md.BackwardCalls()))  // should have called backwards once
	is.Equal(1, md.calls.Backward[0].Val) // should have used an backwards value of 1
}

func TestClockwiseMessageSendsCommandToDrone(t *testing.T) {
	ap, md, _ := setupAutopilot(t)
	is := is.New(t)

	ap.HandleMessage(&messages.Flight{Command: messages.CommandClockwise, Value: 1})

	is.Equal(1, len(md.ClockwiseCalls()))  // should have called clockwise once
	is.Equal(1, md.calls.Clockwise[0].Val) // should have used an clockwise value of 1
}

func TestCounterClockwiseMessageSendsCommandToDrone(t *testing.T) {
	ap, md, _ := setupAutopilot(t)
	is := is.New(t)

	ap.HandleMessage(&messages.Flight{Command: messages.CommandCounterClockwise, Value: 1})

	is.Equal(1, len(md.CounterClockwiseCalls()))  // should have called counter clockwise once
	is.Equal(1, md.calls.CounterClockwise[0].Val) // should have used an counter clockwise value of 1
}

func TestStopMessageSendsCommandToDrone(t *testing.T) {
	ap, md, _ := setupAutopilot(t)
	is := is.New(t)

	ap.HandleMessage(&messages.Flight{Command: messages.CommandStop})

	is.Equal(1, len(md.StopCalls())) // should have called stop once
}

/*
func TestFollowFaceStartsFaceFollowing(t *testing.T) {
	ap, _, c := setupAutopilot(t)
	is := is.New(t)

	ap.HandleMessage(&messages.Flight{Command: messages.CommandFollowFace, Value: 1})

	is.Equal(1, len(c.StartCalls())) // should have called start once
}

func TestFollowFaceStopsFaceFollowing(t *testing.T) {
	ap, _, c := setupAutopilot(t)
	is := is.New(t)

	ap.HandleMessage(&messages.Flight{Command: messages.CommandFollowFace, Value: 0})

	is.Equal(1, len(c.StopCalls())) // should have called stop once
}
*/
func fetchCallsForCallback(m *MamboDroneMock, c string) []struct {
	Name string
	F    func(interface{})
} {
	calls := m.OnCalls()
	rc := make([]struct {
		Name string
		F    func(interface{})
	}, 0)

	for _, call := range calls {
		if call.Name == c {
			rc = append(rc, call)
		}
	}

	return rc
}
