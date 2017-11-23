package control

//go:generate moq -out mocks_test.go . MamboDrone

// MamboDrone defines the interface for our drone
type MamboDrone interface {
	On(name string, f func(s interface{})) error
	Land() error
	TakeOff() error
	Up(val int) error
	Down(val int) error
	Forward(val int) error
	Backward(val int) error
	Right(val int) error
	Left(val int) error
	Clockwise(val int) error
	CounterClockwise(val int) error
}
