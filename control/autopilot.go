package control

import (
	"log"
	"time"

	messages "github.com/nicholasjackson/drone-messages"
	"gobot.io/x/gobot/platforms/parrot/minidrone"
)

// AutoPilot implements an autopilot control for Mambo drones
type AutoPilot struct {
	drone          MamboDrone
	lastFace       *messages.FaceDetected
	following      bool
	minDistance    int
	speed          int
	timeout        time.Duration
	deadMansSwitch *time.Timer
}

// NewAutoPilot creates a new autopilot structure
func NewAutoPilot(d MamboDrone) *AutoPilot {
	return &AutoPilot{
		drone:       d,
		following:   false,
		minDistance: 50,
		speed:       30,
		timeout:     1 * time.Second,
	}
}

// Setup sets up all the callbacks for the drone
func (a *AutoPilot) Setup() error {
	log.Println("Started AutoPilot")

	a.drone.On(minidrone.Battery, func(data interface{}) {
		//writeLog("battery: %d\n", data)
	})

	a.drone.On(minidrone.FlightStatus, func(data interface{}) {
		writeLog("flight status: %#v\n", data)
	})

	a.drone.On(minidrone.Takeoff, func(data interface{}) {
		writeLog("taking off...")
	})

	a.drone.On(minidrone.Hovering, func(data interface{}) {
		writeLog("hovering!")
	})

	a.drone.On(minidrone.Landing, func(data interface{}) {
		writeLog("landing...")
	})

	a.drone.On(minidrone.Landed, func(data interface{}) {
		writeLog("landed.")
	})

	return nil
}

// HandleMessage handles a message sent from the stream
func (a *AutoPilot) HandleMessage(m interface{}) error {
	switch m.(type) {
	case *messages.Flight:
		fm := m.(*messages.Flight)

		switch fm.Command {
		case messages.CommandTakeOff:
			return a.drone.TakeOff()
		case messages.CommandLand:
			return a.drone.Land()
		case messages.CommandUp:
			return a.drone.Up(fm.Value)
		case messages.CommandDown:
			return a.drone.Down(fm.Value)
		case messages.CommandLeft:
			return a.drone.Left(fm.Value)
		case messages.CommandRight:
			return a.drone.Right(fm.Value)
		case messages.CommandForward:
			return a.drone.Forward(fm.Value)
		case messages.CommandBackward:
			return a.drone.Backward(fm.Value)
		case messages.CommandClockwise:
			return a.drone.Clockwise(fm.Value)
		case messages.CommandCounterClockwise:
			return a.drone.CounterClockwise(fm.Value)
		case messages.CommandStop:
			a.drone.Left(0)
			a.drone.Right(0)
			a.drone.Forward(0)
			a.drone.Backward(0)
			a.drone.Up(0)
			a.drone.Down(0)
			return a.drone.Stop()
		case messages.CommandFollowFace:
			if fm.Value == 1 {
				a.StartFollowing()
			} else {
				a.StopFollowing()
			}
		}
	case *messages.FaceDetected:
		a.FollowFace(m.(*messages.FaceDetected))
	}

	return nil
}

func writeLog(format string, data ...interface{}) {
	//log.Printf(format, data)
}
