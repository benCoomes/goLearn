package main

import "fmt"

func main() {
	ints := map[string]int64{
		"first":  24,
		"second": 2,
	}

	floats := map[string]float64{
		"first":  49.9,
		"second": 1.2,
	}

	fmt.Printf("Non generic sums: %v and %v\n", SumInts(ints), SumFloats(floats))
	fmt.Printf("Generic sums: %v and %v\n", SumFloatsOrInts(ints), SumFloatsOrInts(floats))
	fmt.Printf("Generic sums with custom constraint: %v and %v\n", SumNumbers(ints), SumNumbers(floats))
}

func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

func SumFloatsOrInts[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

type Number interface {
	int64 | int32 | float64 | float32
}
