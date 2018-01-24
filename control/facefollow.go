package control

import (
	"log"
	"time"

	messages "github.com/nicholasjackson/drone-messages"
)

// StartFollowing puts the drone into autonomous mode
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
	if a.lastFace != nil && a.following {
		a.moveDrone(m)
	}

	a.lastFace = m
}

func (a *AutoPilot) moveDrone(m *messages.FaceDetected) {
	if len(m.Faces) < 1 {
		return
	}

	log.Println("Got Face, moving...", m)

	// calculate the right moves
	centerPointX := (m.Bounds.Max.X) / 2
	centerPointY := (m.Bounds.Max.Y) / 2
	faceCenterX := ((m.Faces[0].Max.X - m.Faces[0].Min.X) / 2) + m.Faces[0].Min.X
	faceCenterY := ((m.Faces[0].Max.Y - m.Faces[0].Min.Y) / 2) + m.Faces[0].Min.Y

	log.Println("Centre:", centerPointX, centerPointY)
	log.Println("Face Center:", faceCenterX, faceCenterY)
	log.Println("Min distance:", a.minDistance)

	if faceCenterX < (centerPointX - a.minDistance) {
		a.drone.Left(a.speed)
	} else if (centerPointX + a.minDistance) < faceCenterX {
		a.drone.Right(a.speed)
	} else if faceCenterY < (centerPointY - a.minDistance) {
		a.drone.Down(a.speed)
	} else if (centerPointY + a.minDistance) < faceCenterY {
		a.drone.Up(a.speed)
	} else {
		a.drone.Stop()
	}

	a.setDeadMansSwitch()
}

// setDeadMansSwitch sets a stop command after timeout incase no further face
// tracking info is received
func (a *AutoPilot) setDeadMansSwitch() {
	if a.deadMansSwitch == nil {
		a.deadMansSwitch = time.AfterFunc(a.timeout, func() {
			log.Println("DMS Stop", a.timeout)
			a.drone.Stop()
		})

		return
	}

	a.deadMansSwitch.Reset(a.timeout)
}
