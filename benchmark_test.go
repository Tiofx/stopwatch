package stopwatch_test

import (
	"testing"
	"github.com/Tiofx/stopwatch"
)

var (
	resSw  *stopwatch.Stopwatch
	resStr string
)

func BenchmarkNew(b *testing.B) {
	var sw *stopwatch.Stopwatch

	for i := 0; i < b.N; i++ {
		sw = stopwatch.New()
	}

	resSw = sw

	b.ReportAllocs()
}

func BenchmarkSimple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sw := stopwatch.New()
		sw.PressTopButton()
		sw.PressTopButton()
		sw.PressSecondButton()
	}
}

func BenchmarkStopwatch_String(b *testing.B) {
	var string string

	for i := 0; i < b.N; i++ {
		string = stopwatch.Global().String()
	}

	resStr = string
}
