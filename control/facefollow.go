package control

import (
	"log"
	"time"

	messages "github.com/nicholasjackson/drone-messages"
)

// StartFollowing allows the drone to start tracking a face and follow it
func (a *AutoPilot) StartFollowing() {
	log.Println("Start Following")
	a.following = true
	a.lastFace = nil
}

// StopFollowing stops tracking a face
func (a *AutoPilot) StopFollowing() {
	log.Println("Stop Following")
	a.following = false
	a.lastFace = nil
	a.drone.Stop()
}

// FollowFace moves to drone in the direction of a detected face
func (a *AutoPilot) FollowFace(m *messages.FaceDetected) {
	if a.lastFace != nil && a.following && a.lastFace != m {
		a.moveDrone(m)
	}

	a.lastFace = m
}

func (a *AutoPilot) moveDrone(m *messages.FaceDetected) {
	log.Println("Got Face, moving...", m)

	// calculate the right moves
	centerPoint := (m.Bounds.Max.X) / 2
	faceCenter := ((m.Faces[0].Max.X - m.Faces[0].Min.X) / 2) + m.Faces[0].Min.X

	log.Println("Centre:", centerPoint)
	log.Println("Face Center:", faceCenter)

	if faceCenter < (centerPoint - a.minDistance) {
		log.Println("Left")
		a.drone.Left(a.speed)
	} else if (centerPoint + a.minDistance) < faceCenter {
		log.Println("Right")
		a.drone.Right(a.speed)
	} else {
		log.Println("Stop")
		a.drone.Stop()
	}

	a.setDeadMansSwitch()
}

// setDeadMansSwitch sets a stop command after timeout incase no further face
// tracking info is received
func (a *AutoPilot) setDeadMansSwitch() {
	if a.deadMansSwitch == nil {
		log.Println("DMS Stop", a.timeout)
		a.deadMansSwitch = time.AfterFunc(a.timeout, func() {
			a.drone.Stop()
		})

		return
	}

	if !a.deadMansSwitch.Stop() {
		<-a.deadMansSwitch.C
	}
	a.deadMansSwitch.Reset(a.timeout)
}
