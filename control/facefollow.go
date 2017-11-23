package control

import (
	messages "github.com/nicholasjackson/drone-messages"
)

// FollowFace moves to drone in the direction of a detected face
func (a *AutoPilot) FollowFace(m messages.FaceDetected) {
	if a.lastFace != nil {
		a.moveDrone(m)
	}

	a.lastFace = &m
}

func (a *AutoPilot) moveDrone(m messages.FaceDetected) {
	// calculate the right moves
	xDist := m.Faces[0].X - a.lastFace.Faces[0].X
	if xDist > 1 {
		a.drone.Left(1)
	}
}
