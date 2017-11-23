package control

import (
	"log"

	messages "github.com/nicholasjackson/drone-messages"
	"gobot.io/x/gobot/platforms/parrot/minidrone"
)

// AutoPilot implements an autopilot control for Mambo drones
type AutoPilot struct {
	drone    MamboDrone
	lastFace *messages.FaceDetected
}

// NewAutoPilot creates a new autopilot structure
func NewAutoPilot(d MamboDrone) *AutoPilot {
	return &AutoPilot{drone: d}
}

// Setup sets up all the callbacks for the drone
func (a *AutoPilot) Setup() error {
	a.drone.On(minidrone.Battery, func(data interface{}) {
		writeLog("battery: %d\n", data)
	})

	a.drone.On(minidrone.FlightStatus, func(data interface{}) {
		writeLog("flight status: %d\n", data)
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
	case messages.Flight:
		switch m.(messages.Flight).Command {
		case "takeoff":
			return a.drone.TakeOff()

		case "land":
			return a.drone.Land()
		}
	case messages.FaceDetected:
		a.FollowFace(m.(messages.FaceDetected))
	}

	return nil
}

func writeLog(m string, data ...interface{}) {
	log.Println(m, data)
}
