package stopwatch

import (
	"fmt"
)

func Examplestate_HasSplitTime() {
	state := Stopped
	fmt.Println(state.HasSplitTime())
	//	Output: false
}

func Examplestate_WithoutSplitTime() {
	state := Stopped.WithSplitTime()
	fmt.Println(state.WithoutSplitTime() == Stopped)
	//	Output: true
}

func Examplestate_String() {
	fmt.Println(Running.WithSplitTime())
	//	Output: running and record split time
}
