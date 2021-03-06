package stopwatch_test

import (
	. "github.com/Tiofx/stopwatch"
	"testing"
	"time"
)

const epsilon = 5 * time.Millisecond

func isEqual(a, b time.Duration) bool { return time.Duration(uint(a-b)) <= epsilon }

func TestStopwatch_1(t *testing.T) {
	sw := New()
	if !(sw.Display() == 0 && sw.State().WithoutSplitTime() == Initial) {
		t.Fail()
	}
}

func TestStopwatch_StatesMove(t *testing.T) {
	sw := New()

	emptyMethod,
		topButton,
		secondButton,
		runSplit :=
		func() {},
		sw.PressTopButton,
		sw.PressSecondButton,
		func() { sw.PressTopButton(); sw.PressSecondButton() }

	testTable := []struct {
		method   func()
		expected interface{}
	}{
		{emptyMethod, Initial},
		{secondButton, Initial},
		{secondButton, Initial},
		{topButton, Running},
		{topButton, Stopped},
		{secondButton, Initial},
		{runSplit, Running.WithSplitTime()},
		{secondButton, Running},
		{secondButton, Running.WithSplitTime()},
		{topButton, Stopped.WithSplitTime()},
		{topButton, Running.WithSplitTime()},
		{topButton, Stopped.WithSplitTime()},
		{secondButton, Stopped},
		{secondButton, Initial},
	}

	for i, test := range testTable {
		test.method()
		if sw.State() != test.expected {
			t.Errorf("test [%v], exptected [%v] != actual [%v]", i, test.expected, sw.State())
		}
	}
}

func TestStopwatch_Time(t *testing.T) {
	sw := New()

	emptyMethod,
		topButton,
		secondButton :=
		func() {},
		sw.PressTopButton,
		sw.PressSecondButton

	testTable := [][]struct {
		method                 func()
		expectedTime, waitTime time.Duration
	}{
		{
			{emptyMethod, 0, 4 * time.Millisecond},
			{topButton, 0, 10 * time.Millisecond},
			{topButton, 10 * time.Millisecond, 7 * time.Millisecond},
			{emptyMethod, 10 * time.Millisecond, 8 * time.Millisecond},
			{secondButton, 0, 3 * time.Millisecond},
		},
		{
			{topButton, 0, 8 * time.Millisecond},
			{topButton, 8 * time.Millisecond, 4 * time.Millisecond},
			{topButton, 8 * time.Millisecond, 12 * time.Millisecond},
			{secondButton, 20 * time.Millisecond, 6 * time.Millisecond},
			{topButton, 20 * time.Millisecond, 10 * time.Millisecond},
			{topButton, 20 * time.Millisecond, 13 * time.Millisecond},
			{topButton, 20 * time.Millisecond, 13 * time.Millisecond},
			{secondButton, 46 * time.Millisecond, 13 * time.Millisecond},
			{secondButton, 0, 0},
		},
	}

	tester := func(t *testing.T, tests []struct {
		method                 func()
		expectedTime, waitTime time.Duration
	}) {
		for i, test := range tests {
			test.method()
			if !isEqual(sw.Display(), test.expectedTime) {
				t.Errorf("tester [%v], exptected [%v] != actual [%v]", i, test.expectedTime, sw.Display())
			}

			time.Sleep(test.waitTime)
		}
	}

	t.Run("start/idle/stop/reset", func(t *testing.T) { tester(t, testTable[0]) })
	t.Run("split time", func(t *testing.T) { tester(t, testTable[1]) })
}
