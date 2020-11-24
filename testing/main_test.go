package main

import (
	"fmt"
	"testing"
)

// Run with go test
// can also add -cover flag for coverage measurement
// use go test -coverprofile <filename>.out to output cover analysis.
// then run go tool cover -html:=<filename>.out to view visual representation of code coverage.

func TestAverage(t *testing.T) {

	type test struct {
		data   []float64
		answer float64
	}

	table := []test{
		test{[]float64{1.0, 2.0}, 1.5},
		test{[]float64{3.0, 3.0, 3.0}, 3.0},
		test{[]float64{9.0, 6.0, 12.0}, 9.0},
		test{[]float64{100.0, 68.0}, 84.0},
		test{[]float64{7.0}, 7.0},
	}

	for _, v := range table {
		ans := Average(v.data)
		if ans != v.answer {
			t.Error("Expected ", v.answer, " got ", ans)
		}
	}
}

// Example tests
// You can add some comments in the example function which show as examples in the compiled godoc.

func ExampleAverage() {
	fmt.Println(Average([]float64{1.0, 2.0}))
	// Output:
	// 1.5
}

//Benchmarking
// run with `go test -bench <regex or . for all>`

func BenchmarkAverage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Average([]float64{9.0, 6.0, 12.0})
	}
}
