package stopwatch

import (
	"fmt"
)

func ExampleState_HasSplitTime() {
	state := Stopped
	fmt.Println(state.HasSplitTime())
	//	Output: false
}

func ExampleState_WithoutSplitTime() {
	state := Stopped.WithSplitTime()
	fmt.Println(state.WithoutSplitTime() == Stopped)
	//	Output: true
}

func ExampleState_String() {
	fmt.Println(Running.WithSplitTime())
	//	Output: running and record split time
}
