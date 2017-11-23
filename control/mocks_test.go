// Code generated by moq; DO NOT EDIT
// github.com/matryer/moq

package control

import (
	"sync"
)

var (
	lockMamboDroneMockBackward         sync.RWMutex
	lockMamboDroneMockClockwise        sync.RWMutex
	lockMamboDroneMockCounterClockwise sync.RWMutex
	lockMamboDroneMockDown             sync.RWMutex
	lockMamboDroneMockForward          sync.RWMutex
	lockMamboDroneMockLand             sync.RWMutex
	lockMamboDroneMockLeft             sync.RWMutex
	lockMamboDroneMockOn               sync.RWMutex
	lockMamboDroneMockRight            sync.RWMutex
	lockMamboDroneMockTakeOff          sync.RWMutex
	lockMamboDroneMockUp               sync.RWMutex
)

// MamboDroneMock is a mock implementation of MamboDrone.
//
//     func TestSomethingThatUsesMamboDrone(t *testing.T) {
//
//         // make and configure a mocked MamboDrone
//         mockedMamboDrone := &MamboDroneMock{
//             BackwardFunc: func(val int) error {
// 	               panic("TODO: mock out the Backward method")
//             },
//             ClockwiseFunc: func(val int) error {
// 	               panic("TODO: mock out the Clockwise method")
//             },
//             CounterClockwiseFunc: func(val int) error {
// 	               panic("TODO: mock out the CounterClockwise method")
//             },
//             DownFunc: func(val int) error {
// 	               panic("TODO: mock out the Down method")
//             },
//             ForwardFunc: func(val int) error {
// 	               panic("TODO: mock out the Forward method")
//             },
//             LandFunc: func() error {
// 	               panic("TODO: mock out the Land method")
//             },
//             LeftFunc: func(val int) error {
// 	               panic("TODO: mock out the Left method")
//             },
//             OnFunc: func(name string, f func(s interface{})) error {
// 	               panic("TODO: mock out the On method")
//             },
//             RightFunc: func(val int) error {
// 	               panic("TODO: mock out the Right method")
//             },
//             TakeOffFunc: func() error {
// 	               panic("TODO: mock out the TakeOff method")
//             },
//             UpFunc: func(val int) error {
// 	               panic("TODO: mock out the Up method")
//             },
//         }
//
//         // TODO: use mockedMamboDrone in code that requires MamboDrone
//         //       and then make assertions.
//
//     }
type MamboDroneMock struct {
	// BackwardFunc mocks the Backward method.
	BackwardFunc func(val int) error

	// ClockwiseFunc mocks the Clockwise method.
	ClockwiseFunc func(val int) error

	// CounterClockwiseFunc mocks the CounterClockwise method.
	CounterClockwiseFunc func(val int) error

	// DownFunc mocks the Down method.
	DownFunc func(val int) error

	// ForwardFunc mocks the Forward method.
	ForwardFunc func(val int) error

	// LandFunc mocks the Land method.
	LandFunc func() error

	// LeftFunc mocks the Left method.
	LeftFunc func(val int) error

	// OnFunc mocks the On method.
	OnFunc func(name string, f func(s interface{})) error

	// RightFunc mocks the Right method.
	RightFunc func(val int) error

	// TakeOffFunc mocks the TakeOff method.
	TakeOffFunc func() error

	// UpFunc mocks the Up method.
	UpFunc func(val int) error

	// calls tracks calls to the methods.
	calls struct {
		// Backward holds details about calls to the Backward method.
		Backward []struct {
			// Val is the val argument value.
			Val int
		}
		// Clockwise holds details about calls to the Clockwise method.
		Clockwise []struct {
			// Val is the val argument value.
			Val int
		}
		// CounterClockwise holds details about calls to the CounterClockwise method.
		CounterClockwise []struct {
			// Val is the val argument value.
			Val int
		}
		// Down holds details about calls to the Down method.
		Down []struct {
			// Val is the val argument value.
			Val int
		}
		// Forward holds details about calls to the Forward method.
		Forward []struct {
			// Val is the val argument value.
			Val int
		}
		// Land holds details about calls to the Land method.
		Land []struct {
		}
		// Left holds details about calls to the Left method.
		Left []struct {
			// Val is the val argument value.
			Val int
		}
		// On holds details about calls to the On method.
		On []struct {
			// Name is the name argument value.
			Name string
			// F is the f argument value.
			F func(s interface{})
		}
		// Right holds details about calls to the Right method.
		Right []struct {
			// Val is the val argument value.
			Val int
		}
		// TakeOff holds details about calls to the TakeOff method.
		TakeOff []struct {
		}
		// Up holds details about calls to the Up method.
		Up []struct {
			// Val is the val argument value.
			Val int
		}
	}
}

// Backward calls BackwardFunc.
func (mock *MamboDroneMock) Backward(val int) error {
	if mock.BackwardFunc == nil {
		panic("moq: MamboDroneMock.BackwardFunc is nil but MamboDrone.Backward was just called")
	}
	callInfo := struct {
		Val int
	}{
		Val: val,
	}
	lockMamboDroneMockBackward.Lock()
	mock.calls.Backward = append(mock.calls.Backward, callInfo)
	lockMamboDroneMockBackward.Unlock()
	return mock.BackwardFunc(val)
}

// BackwardCalls gets all the calls that were made to Backward.
// Check the length with:
//     len(mockedMamboDrone.BackwardCalls())
func (mock *MamboDroneMock) BackwardCalls() []struct {
	Val int
} {
	var calls []struct {
		Val int
	}
	lockMamboDroneMockBackward.RLock()
	calls = mock.calls.Backward
	lockMamboDroneMockBackward.RUnlock()
	return calls
}

// Clockwise calls ClockwiseFunc.
func (mock *MamboDroneMock) Clockwise(val int) error {
	if mock.ClockwiseFunc == nil {
		panic("moq: MamboDroneMock.ClockwiseFunc is nil but MamboDrone.Clockwise was just called")
	}
	callInfo := struct {
		Val int
	}{
		Val: val,
	}
	lockMamboDroneMockClockwise.Lock()
	mock.calls.Clockwise = append(mock.calls.Clockwise, callInfo)
	lockMamboDroneMockClockwise.Unlock()
	return mock.ClockwiseFunc(val)
}

// ClockwiseCalls gets all the calls that were made to Clockwise.
// Check the length with:
//     len(mockedMamboDrone.ClockwiseCalls())
func (mock *MamboDroneMock) ClockwiseCalls() []struct {
	Val int
} {
	var calls []struct {
		Val int
	}
	lockMamboDroneMockClockwise.RLock()
	calls = mock.calls.Clockwise
	lockMamboDroneMockClockwise.RUnlock()
	return calls
}

// CounterClockwise calls CounterClockwiseFunc.
func (mock *MamboDroneMock) CounterClockwise(val int) error {
	if mock.CounterClockwiseFunc == nil {
		panic("moq: MamboDroneMock.CounterClockwiseFunc is nil but MamboDrone.CounterClockwise was just called")
	}
	callInfo := struct {
		Val int
	}{
		Val: val,
	}
	lockMamboDroneMockCounterClockwise.Lock()
	mock.calls.CounterClockwise = append(mock.calls.CounterClockwise, callInfo)
	lockMamboDroneMockCounterClockwise.Unlock()
	return mock.CounterClockwiseFunc(val)
}

// CounterClockwiseCalls gets all the calls that were made to CounterClockwise.
// Check the length with:
//     len(mockedMamboDrone.CounterClockwiseCalls())
func (mock *MamboDroneMock) CounterClockwiseCalls() []struct {
	Val int
} {
	var calls []struct {
		Val int
	}
	lockMamboDroneMockCounterClockwise.RLock()
	calls = mock.calls.CounterClockwise
	lockMamboDroneMockCounterClockwise.RUnlock()
	return calls
}

// Down calls DownFunc.
func (mock *MamboDroneMock) Down(val int) error {
	if mock.DownFunc == nil {
		panic("moq: MamboDroneMock.DownFunc is nil but MamboDrone.Down was just called")
	}
	callInfo := struct {
		Val int
	}{
		Val: val,
	}
	lockMamboDroneMockDown.Lock()
	mock.calls.Down = append(mock.calls.Down, callInfo)
	lockMamboDroneMockDown.Unlock()
	return mock.DownFunc(val)
}

// DownCalls gets all the calls that were made to Down.
// Check the length with:
//     len(mockedMamboDrone.DownCalls())
func (mock *MamboDroneMock) DownCalls() []struct {
	Val int
} {
	var calls []struct {
		Val int
	}
	lockMamboDroneMockDown.RLock()
	calls = mock.calls.Down
	lockMamboDroneMockDown.RUnlock()
	return calls
}

// Forward calls ForwardFunc.
func (mock *MamboDroneMock) Forward(val int) error {
	if mock.ForwardFunc == nil {
		panic("moq: MamboDroneMock.ForwardFunc is nil but MamboDrone.Forward was just called")
	}
	callInfo := struct {
		Val int
	}{
		Val: val,
	}
	lockMamboDroneMockForward.Lock()
	mock.calls.Forward = append(mock.calls.Forward, callInfo)
	lockMamboDroneMockForward.Unlock()
	return mock.ForwardFunc(val)
}

// ForwardCalls gets all the calls that were made to Forward.
// Check the length with:
//     len(mockedMamboDrone.ForwardCalls())
func (mock *MamboDroneMock) ForwardCalls() []struct {
	Val int
} {
	var calls []struct {
		Val int
	}
	lockMamboDroneMockForward.RLock()
	calls = mock.calls.Forward
	lockMamboDroneMockForward.RUnlock()
	return calls
}

// Land calls LandFunc.
func (mock *MamboDroneMock) Land() error {
	if mock.LandFunc == nil {
		panic("moq: MamboDroneMock.LandFunc is nil but MamboDrone.Land was just called")
	}
	callInfo := struct {
	}{}
	lockMamboDroneMockLand.Lock()
	mock.calls.Land = append(mock.calls.Land, callInfo)
	lockMamboDroneMockLand.Unlock()
	return mock.LandFunc()
}

// LandCalls gets all the calls that were made to Land.
// Check the length with:
//     len(mockedMamboDrone.LandCalls())
func (mock *MamboDroneMock) LandCalls() []struct {
} {
	var calls []struct {
	}
	lockMamboDroneMockLand.RLock()
	calls = mock.calls.Land
	lockMamboDroneMockLand.RUnlock()
	return calls
}

// Left calls LeftFunc.
func (mock *MamboDroneMock) Left(val int) error {
	if mock.LeftFunc == nil {
		panic("moq: MamboDroneMock.LeftFunc is nil but MamboDrone.Left was just called")
	}
	callInfo := struct {
		Val int
	}{
		Val: val,
	}
	lockMamboDroneMockLeft.Lock()
	mock.calls.Left = append(mock.calls.Left, callInfo)
	lockMamboDroneMockLeft.Unlock()
	return mock.LeftFunc(val)
}

// LeftCalls gets all the calls that were made to Left.
// Check the length with:
//     len(mockedMamboDrone.LeftCalls())
func (mock *MamboDroneMock) LeftCalls() []struct {
	Val int
} {
	var calls []struct {
		Val int
	}
	lockMamboDroneMockLeft.RLock()
	calls = mock.calls.Left
	lockMamboDroneMockLeft.RUnlock()
	return calls
}

// On calls OnFunc.
func (mock *MamboDroneMock) On(name string, f func(s interface{})) error {
	if mock.OnFunc == nil {
		panic("moq: MamboDroneMock.OnFunc is nil but MamboDrone.On was just called")
	}
	callInfo := struct {
		Name string
		F    func(s interface{})
	}{
		Name: name,
		F:    f,
	}
	lockMamboDroneMockOn.Lock()
	mock.calls.On = append(mock.calls.On, callInfo)
	lockMamboDroneMockOn.Unlock()
	return mock.OnFunc(name, f)
}

// OnCalls gets all the calls that were made to On.
// Check the length with:
//     len(mockedMamboDrone.OnCalls())
func (mock *MamboDroneMock) OnCalls() []struct {
	Name string
	F    func(s interface{})
} {
	var calls []struct {
		Name string
		F    func(s interface{})
	}
	lockMamboDroneMockOn.RLock()
	calls = mock.calls.On
	lockMamboDroneMockOn.RUnlock()
	return calls
}

// Right calls RightFunc.
func (mock *MamboDroneMock) Right(val int) error {
	if mock.RightFunc == nil {
		panic("moq: MamboDroneMock.RightFunc is nil but MamboDrone.Right was just called")
	}
	callInfo := struct {
		Val int
	}{
		Val: val,
	}
	lockMamboDroneMockRight.Lock()
	mock.calls.Right = append(mock.calls.Right, callInfo)
	lockMamboDroneMockRight.Unlock()
	return mock.RightFunc(val)
}

// RightCalls gets all the calls that were made to Right.
// Check the length with:
//     len(mockedMamboDrone.RightCalls())
func (mock *MamboDroneMock) RightCalls() []struct {
	Val int
} {
	var calls []struct {
		Val int
	}
	lockMamboDroneMockRight.RLock()
	calls = mock.calls.Right
	lockMamboDroneMockRight.RUnlock()
	return calls
}

// TakeOff calls TakeOffFunc.
func (mock *MamboDroneMock) TakeOff() error {
	if mock.TakeOffFunc == nil {
		panic("moq: MamboDroneMock.TakeOffFunc is nil but MamboDrone.TakeOff was just called")
	}
	callInfo := struct {
	}{}
	lockMamboDroneMockTakeOff.Lock()
	mock.calls.TakeOff = append(mock.calls.TakeOff, callInfo)
	lockMamboDroneMockTakeOff.Unlock()
	return mock.TakeOffFunc()
}

// TakeOffCalls gets all the calls that were made to TakeOff.
// Check the length with:
//     len(mockedMamboDrone.TakeOffCalls())
func (mock *MamboDroneMock) TakeOffCalls() []struct {
} {
	var calls []struct {
	}
	lockMamboDroneMockTakeOff.RLock()
	calls = mock.calls.TakeOff
	lockMamboDroneMockTakeOff.RUnlock()
	return calls
}

// Up calls UpFunc.
func (mock *MamboDroneMock) Up(val int) error {
	if mock.UpFunc == nil {
		panic("moq: MamboDroneMock.UpFunc is nil but MamboDrone.Up was just called")
	}
	callInfo := struct {
		Val int
	}{
		Val: val,
	}
	lockMamboDroneMockUp.Lock()
	mock.calls.Up = append(mock.calls.Up, callInfo)
	lockMamboDroneMockUp.Unlock()
	return mock.UpFunc(val)
}

// UpCalls gets all the calls that were made to Up.
// Check the length with:
//     len(mockedMamboDrone.UpCalls())
func (mock *MamboDroneMock) UpCalls() []struct {
	Val int
} {
	var calls []struct {
		Val int
	}
	lockMamboDroneMockUp.RLock()
	calls = mock.calls.Up
	lockMamboDroneMockUp.RUnlock()
	return calls
}
