package stopwatch_test

import (
	. "github.com/Tiofx/stopwatch"
	"testing"
)

var (
	resSw  *Stopwatch
	resStr string
)

func BenchmarkNew(b *testing.B) {
	var sw *Stopwatch

	for i := 0; i < b.N; i++ {
		sw = New()
	}

	resSw = sw

	b.ReportAllocs()
}

func BenchmarkSimple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sw := New()
		sw.PressTopButton()
		sw.PressTopButton()
		sw.PressSecondButton()
	}
}

func BenchmarkStopwatch_String(b *testing.B) {
	var string string

	for i := 0; i < b.N; i++ {
		string = Global().String()
	}

	resStr = string
}
