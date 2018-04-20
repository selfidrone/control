package control

import (
	hclog "github.com/hashicorp/go-hclog"
)

// SimulatedDrone implements the MamboDrone interface and is used to simulate the interaction with a real drone
// useful for testing things like autopilot
type SimulatedDrone struct {
	log hclog.Logger
}

func NewSimulatedDrone(log hclog.Logger) *SimulatedDrone {
	return &SimulatedDrone{log}
}

// On implements the MamboDrone interface on
func (s *SimulatedDrone) On(name string, f func(s interface{})) error {
	s.log.Info("On", "name", name)
	return nil
}

// Land implements the MamboDrone interface on
func (s *SimulatedDrone) Land() error {
	s.log.Info("Command", "name", "Land")
	return nil
}

// TakeOff implements the MamboDrone interface on
func (s *SimulatedDrone) TakeOff() error {
	s.log.Info("Command", "name", "TakeOff")
	return nil
}

// FlatTrim implements the MamboDrone interface on
func (s *SimulatedDrone) FlatTrim() error {
	s.log.Info("Command", "name", "FlatTrim")
	return nil
}

// Up implements the MamboDrone interface on
func (s *SimulatedDrone) Up(val int) error {
	s.log.Info("Command", "name", "Up", "val", val)
	return nil
}

// Down implements the MamboDrone interface on
func (s *SimulatedDrone) Down(val int) error {
	s.log.Info("Command", "name", "Down", "val", val)
	return nil
}

// Forward implements the MamboDrone interface on
func (s *SimulatedDrone) Forward(val int) error {
	s.log.Info("Command", "name", "Forward", "val", val)
	return nil
}

// Backward implements the MamboDrone interface on
func (s *SimulatedDrone) Backward(val int) error {
	s.log.Info("Command", "name", "Backward", "val", val)
	return nil
}

// Right implements the MamboDrone interface on
func (s *SimulatedDrone) Right(val int) error {
	s.log.Info("Command", "name", "Right", "val", val)
	return nil
}

// Left implements the MamboDrone interface on
func (s *SimulatedDrone) Left(val int) error {
	s.log.Info("Command", "name", "Left", "val", val)
	return nil
}

// Clockwise implements the MamboDrone interface on
func (s *SimulatedDrone) Clockwise(val int) error {
	s.log.Info("Command", "name", "Clockwise", "val", val)
	return nil
}

// CounterClockwise implements the MamboDrone interface on
func (s *SimulatedDrone) CounterClockwise(val int) error {
	s.log.Info("Command", "name", "CounterClockwise", "val", val)
	return nil
}

// Stop implements the MamboDrone interface on
func (s *SimulatedDrone) Stop() error {
	s.log.Info("Command", "name", "Stop")
	return nil
}
