package main

import "fmt"

/*
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func sumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}
*/
// add Generic func:

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

type Number interface {
	int64 | float64
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func main() {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}
	floats := map[string]float64{
		"first":  35.09,
		"second": 26.99,
	}
	fmt.Printf("Generics ssums : %v ve %v\n",
		SumNumbers(ints),
		SumNumbers(floats),
	)
	fmt.Printf("Sums Ints Or Float : %v ve %v\n", SumIntsOrFloats(ints), SumIntsOrFloats(floats))

	/*
		fmt.Printf("Non-generic sums: %v ve %v\n", SumInts(ints), sumFloats(floats))
	*/
}
