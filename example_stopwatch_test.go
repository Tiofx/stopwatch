package stopwatch_test

import (
	"github.com/Tiofx/stopwatch"
	"fmt"
	"time"
)

func ExampleStopwatch_Display_1() {
	fmt.Println(stopwatch.New().Display())
	//	Output: 0s
}

// Show method invocation time.
func ExampleStopwatch_Display_2() {
	s := stopwatch.New()

	s.PressTopButton()
	s.PressTopButton()

	fmt.Println(s.Display())
}

func ExampleStopwatch_Display_3() {
	delay := 10 * time.Millisecond
	s := stopwatch.New()

	s.PressTopButton()
	time.Sleep(delay)
	s.PressTopButton()

	fmt.Println(isEqual(s.Display(), delay))
	//	Output: true
}

func ExampleStopwatch_State() {
	fmt.Println(stopwatch.New().State() == stopwatch.Initial)
	//	Output: true
}

func ExampleStopwatch_State_1() {
	s := stopwatch.New()
	s.PressTopButton()

	fmt.Println(s.State() == stopwatch.Running)
	//	Output: true
}

func ExampleStopwatch_State_2a() {
	s := stopwatch.New()
	s.PressTopButton()
	s.PressSecondButton()

	fmt.Println(s.State() == stopwatch.Running)
	//	Output: false
}

func ExampleStopwatch_State_2b() {
	s := stopwatch.New()
	s.PressTopButton()
	s.PressSecondButton()

	fmt.Println(
		s.State() == stopwatch.Running.WithSplitTime() &&
			s.State().HasSplitTime() &&
			s.State().WithoutSplitTime() == stopwatch.Running,
	)
	//	Output: true
}

//Show example of representation Stopwatch as string
func ExampleStopwatch_String() {
	s := stopwatch.New()
	s.PressTopButton()
	time.Sleep(1 * time.Millisecond)
	s.PressTopButton()

	fmt.Println(s)
}
