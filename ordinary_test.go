package preserves

import "testing"

// -------
// Testing
// -------

func TestRandInt(t *testing.T) {
	testVar := make(map[int]int64)
	for i := 0; i < 10; i++ {
		testVar[i] = RandInt(0, 10000000) // calling the tested function
	}
	for _, v1 := range testVar {
		count := 0
		for _, v2 := range testVar {
			if v1 == v2 {
				count++
			}
		}
		// work check
		if count > 1 {
			t.Error("RandInt() = Error generating unique random integer.")
		}
	}
}

// ------------
// Benchmarking
// ------------

func Benchmark_RandInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = RandInt(0, int64(i+1)) // calling the tested function
	}
}
