package stopwatch

import "sync"

var instance *Stopwatch
var once sync.Once

// Global single and thread-save instance of Stopwatch.
// It can be used in situation when the same Stopwatch
// needed in a different places and passing around an instance
// is inconvenient.
func Global() *Stopwatch {
	once.Do(func() { instance = New() })

	return instance
}
