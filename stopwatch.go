package stopwatch

import (
	"time"
	"fmt"
)

type Stopwatch struct {
	start            time.Time
	totalElapsed     time.Duration
	displayedElapsed time.Duration
	state            state
}

func New() *Stopwatch {
	stopwatch := Stopwatch{}
	stopwatch.reset()
	return &stopwatch
}

func (s *Stopwatch) Display() time.Duration {
	if s.state == Running {
		return s.displayedElapsed + s.fromStart()
	}

	return s.displayedElapsed
}

func (s *Stopwatch) State() state { return s.state }

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
