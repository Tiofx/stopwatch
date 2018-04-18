package stopwatch_test

import (
	"fmt"
	. "github.com/Tiofx/stopwatch"
	"time"
)

func ExampleStopwatch_Display_first() {
	fmt.Println(New().Display())
	//	Output: 0s
}

// Show method invocation time.
func ExampleStopwatch_Display_second() {
	s := New()

	s.PressTopButton()
	s.PressTopButton()

	fmt.Println(s.Display())
}

func ExampleStopwatch_Display_third() {
	delay := 10 * time.Millisecond
	s := New()

	s.PressTopButton()
	time.Sleep(delay)
	s.PressTopButton()

	fmt.Println(isEqual(s.Display(), delay))
	//	Output: true
}

func ExampleStopwatch_State() {
	fmt.Println(New().State() == Initial)
	//	Output: true
}

func ExampleStopwatch_State_first() {
	s := New()
	s.PressTopButton()

	fmt.Println(s.State() == Running)
	//	Output: true
}

func ExampleStopwatch_State_secondA() {
	s := New()
	s.PressTopButton()
	s.PressSecondButton()

	fmt.Println(s.State() == Running)
	//	Output: false
}

func ExampleStopwatch_State_secondB() {
	s := New()
	s.PressTopButton()
	s.PressSecondButton()

	fmt.Println(
		s.State() == Running.WithSplitTime() &&
			s.State().HasSplitTime() &&
			s.State().WithoutSplitTime() == Running,
	)
	//	Output: true
}

//Show example of representation Stopwatch as string
func ExampleStopwatch_String() {
	s := New()
	s.PressTopButton()
	time.Sleep(1 * time.Millisecond)
	s.PressTopButton()

	fmt.Println(s)
}
