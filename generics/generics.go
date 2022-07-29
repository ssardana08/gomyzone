package main

import "fmt"

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func main() {
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}
	fmt.Println(SumIntsOrFloats(floats))
}
