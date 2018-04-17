package stopwatch

import "fmt"

// state represent Stopwatch's current state.
type state int

// HasSplitTime check presents of splitTime state.
func (s state) HasSplitTime() bool { return s&splitTime != 0 }

// WithoutSplitTime returning state without splitTime.
// Its let check main state without bitwise operation:
//
// Examples for check on Initial state:
// Incorrect if state has splitTime:
// 		sw.state() == Initial
//
// Correct but almost unreadable:
// 		(sw.state() | splitTime) ^ splitTime == Initial
//
// Perfect:
//		sw.state().WithoutSplitTime() == Initial
func (s state) WithoutSplitTime() state { return (s | splitTime) ^ splitTime }

// WithSplitTime add splitTime state.
// If state already has splitTime than nothing changes.
func (s state) WithSplitTime() state { return s | splitTime }

// Set change receiver to `to` state without
// changing splitTime of receiver.
func (s *state) Set(to state) {
	if s.HasSplitTime() {
		*s = to.WithSplitTime()
	} else {
		*s = to
	}
}

// Stopwatch can be in several states:
// Initial, Running, Stopped.
// Additionally to Running or Stooped states
// Stopwatch can has the splitTime state.
//
// Initial - initial state that has Stopwatch right after creation or
// after reset.
//
// Running - state in which Stopwatch's displayed elapsed time
// change over time.
//
// Stopped - state in which Stopwatch Display the same time over time.
//
//
// splitTime - when Stopwatch get this state its displayed time will be frozen
// but elapsed time can be still increasing.
// When Stopwatch get out from this state
// its displayed time will be updated and unfrozen.
const (
	Initial state = iota
	Running
	Stopped

	splitTime state = 1 << 2
)

func (s state) String() string {
	var result string

	switch s.WithoutSplitTime() {

	case Initial:
		result = "initial"

	case Running:
		result = "running"

	case Stopped:
		result = "stopped"

	default:
		panic(fmt.Sprintf("impossible number: [%v]", int(s)))
	}

	if s.HasSplitTime() {
		result += " and record split time"
	}

	return result
}
