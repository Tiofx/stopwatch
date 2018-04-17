package stopwatch_test

import (
	"github.com/Tiofx/stopwatch"
	"fmt"
)

func ExampleState_HasSplitTime() {
	state := stopwatch.Stopped
	fmt.Println(state.HasSplitTime())
	//	Output: false
}

func ExampleState_WithoutSplitTime() {
	state := stopwatch.Stopped.WithSplitTime()
	fmt.Println(state.WithoutSplitTime() == stopwatch.Stopped)
	//	Output: true
}

func ExampleState_String() {
	fmt.Println(stopwatch.Running.WithSplitTime())
	//	Output: running and record split time
}