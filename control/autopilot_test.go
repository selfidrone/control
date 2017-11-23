package control

import (
	"testing"

	"github.com/matryer/is"
	messages "github.com/nicholasjackson/drone-messages"
	"gobot.io/x/gobot/platforms/parrot/minidrone"
)

func setupAutopilot(t *testing.T) (*AutoPilot, *MamboDroneMock) {
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
	}

	ap := NewAutoPilot(mockedMamboDrone)

	return ap, mockedMamboDrone
}

func TestSetupAddsBatteryCallback(t *testing.T) {
	ap, md := setupAutopilot(t)
	is := is.New(t)

	ap.Setup()

	c := fetchCallsForCallback(md, minidrone.Battery)

	is.Equal(len(c), 1) // expected setup to add battery callback
}

func TestSetupAddsFlightStatusCallback(t *testing.T) {
	ap, md := setupAutopilot(t)
	is := is.New(t)

	ap.Setup()

	c := fetchCallsForCallback(md, minidrone.FlightStatus)

	is.Equal(len(c), 1) // expected setup to add flight status callback
}

func TestSetupAddsTakeoffCallback(t *testing.T) {
	ap, md := setupAutopilot(t)
	is := is.New(t)

	ap.Setup()

	c := fetchCallsForCallback(md, minidrone.Takeoff)

	is.Equal(len(c), 1) // expected setup to add takeoff status callback
}

func TestSetupAddsHoveringCallback(t *testing.T) {
	ap, md := setupAutopilot(t)
	is := is.New(t)

	ap.Setup()

	c := fetchCallsForCallback(md, minidrone.Hovering)

	is.Equal(len(c), 1) // expected setup to add hovering status callback
}

func TestSetupAddsLandingCallback(t *testing.T) {
	ap, md := setupAutopilot(t)
	is := is.New(t)

	ap.Setup()

	c := fetchCallsForCallback(md, minidrone.Landing)

	is.Equal(len(c), 1) // expected setup to add landing status callback
}

func TestSetupAddsLandedCallback(t *testing.T) {
	ap, md := setupAutopilot(t)
	is := is.New(t)

	ap.Setup()

	c := fetchCallsForCallback(md, minidrone.Landed)

	is.Equal(len(c), 1) // expected setup to add landed status callback
}

func TestTakeoffMessageSendsCommandToDrone(t *testing.T) {
	ap, md := setupAutopilot(t)
	is := is.New(t)

	ap.HandleMessage(messages.Flight{Command: messages.CommandTakeOff})

	is.Equal(len(md.TakeOffCalls()), 1) // should have called takeoff once
}

func TestLandMessageSendsCommandToDrone(t *testing.T) {
	ap, md := setupAutopilot(t)
	is := is.New(t)

	ap.HandleMessage(messages.Flight{Command: messages.CommandLand})

	is.Equal(len(md.LandCalls()), 1) // should have called takeoff once
}

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
