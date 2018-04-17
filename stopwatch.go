package stopwatch

import (
	"time"
	"fmt"
)

// The Stopwatch type represents a two button Stopwatch.
type Stopwatch struct {
	start            time.Time
	totalElapsed     time.Duration
	displayedElapsed time.Duration
	state            state
}

// New creates a new reseted Stopwatch.
func New() *Stopwatch {
	stopwatch := Stopwatch{}
	stopwatch.reset()
	return &stopwatch
}

// Display shows time that currently displayed by Stopwatch.
func (s *Stopwatch) Display() time.Duration {
	if s.state == Running {
		return s.displayedElapsed + s.fromStart()
	}

	return s.displayedElapsed
}

// State returns current state of Stopwatch.
func (s *Stopwatch) State() state { return s.state }

// PressTopButton changes stopwatch's state Stopped/Running.
//
// The method can change state of Stopwatch:
//	- from Initial to Running
//	- from Running to Stopped
//	- from Stopped to Running.
func (s *Stopwatch) PressTopButton() {

	switch s.state.WithoutSplitTime() {

	case Initial, Stopped:
		s.state.Set(Running)
		s.updateStart()

	case Running:
		s.state.Set(Stopped)
		s.updateTotalElapsed()

	default:
		panic(s.state)

	}

	if !s.state.HasSplitTime() {
		s.displayedElapsed = s.totalElapsed
		return
	}
}

// PressSecondButton switches splitTime state or reset Stopwatch.
//
// The method changes state of Stopwatch:
//
//  - from Initial to Initial
//  - from Running to Running with splitTime
//  - from Stopped to Initial
//  - from Running with splitTime to Running
//  - from Stopped with splitTime to Stopped.
func (s *Stopwatch) PressSecondButton() {
	if s.state.HasSplitTime() {
		s.state = s.state.WithoutSplitTime()
		s.displayedElapsed = s.totalElapsed
		return
	}

	switch s.state.WithoutSplitTime() {

	case Initial:

	case Running:
		s.displayedElapsed = s.Display()
		s.state |= splitTime
		s.updateTotalElapsed()
		s.updateStart()

	case Stopped:
		s.reset()

	default:
		panic(s.state)
	}
}

func (s Stopwatch) String() string {
	return fmt.Sprintf("Stopwatch is [%v] and display [%v]", s.state, s.Display())
}

func (s *Stopwatch) reset() {
	s.start = time.Time{}
	s.state = Initial
	s.totalElapsed = 0
	s.displayedElapsed = 0
}

func (s *Stopwatch) fromStart() time.Duration { return time.Since(s.start) }
func (s *Stopwatch) updateStart()             { s.start = time.Now() }

func (s *Stopwatch) updateTotalElapsed() { s.totalElapsed += s.fromStart() }
