package control

import "log"

// SimulatedDrone implements the MamboDrone interface and is used to simulate the interaction with a real drone
// useful for testing things like autopilot
type SimulatedDrone struct{}

// On implements the MamboDrone interface on
func (s *SimulatedDrone) On(name string, f func(s interface{})) error {
	log.Println("On name:", name)
	return nil
}

// Land implements the MamboDrone interface on
func (s *SimulatedDrone) Land() error {
	log.Println("Simulated: Land")
	return nil
}

// TakeOff implements the MamboDrone interface on
func (s *SimulatedDrone) TakeOff() error {
	log.Println("Simulated: TakeOff")
	return nil
}

// FlatTrim implements the MamboDrone interface on
func (s *SimulatedDrone) FlatTrim() error {
	log.Println("Simulated: FlatTrim")
	return nil
}

// Up implements the MamboDrone interface on
func (s *SimulatedDrone) Up(val int) error {
	log.Println("Simulated: Up")
	return nil
}

// Down implements the MamboDrone interface on
func (s *SimulatedDrone) Down(val int) error {
	log.Println("Simulated: Down")
	return nil
}

// Forward implements the MamboDrone interface on
func (s *SimulatedDrone) Forward(val int) error {
	log.Println("Simulated: Forward")
	return nil
}

// Backward implements the MamboDrone interface on
func (s *SimulatedDrone) Backward(val int) error {
	log.Println("Simulated: Backward")
	return nil
}

// Right implements the MamboDrone interface on
func (s *SimulatedDrone) Right(val int) error {
	log.Println("Simulated: Right")
	return nil
}

// Left implements the MamboDrone interface on
func (s *SimulatedDrone) Left(val int) error {
	log.Println("Simulated: Left")
	return nil
}

// Clockwise implements the MamboDrone interface on
func (s *SimulatedDrone) Clockwise(val int) error {
	log.Println("Simulated: Clockwise")
	return nil
}

// CounterClockwise implements the MamboDrone interface on
func (s *SimulatedDrone) CounterClockwise(val int) error {
	log.Println("Simulated: CounterClockwise")
	return nil
}

// Stop implements the MamboDrone interface on
func (s *SimulatedDrone) Stop() error {
	log.Println("Simulated: Stop")
	return nil
}
