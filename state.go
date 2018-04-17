package stopwatch

import "fmt"

type state int

func (s state) HasSplitTime() bool { return s&splitTime != 0 }

func (s state) WithoutSplitTime() state { return (s | splitTime) ^ splitTime }

func (s *state) Set(to state) {
	if s.HasSplitTime() {
		*s = to.WithSplitTime()
	} else {
		*s = to
	}
}

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
