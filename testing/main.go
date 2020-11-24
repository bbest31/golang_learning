// Package main
package main

// Average takes in a slice of 64-bit floats and returns the average.
func Average(fs []float64) float64 {
	var s float64
	for _, v := range fs {
		s += v
	}

	return s / float64(len(fs))
}
