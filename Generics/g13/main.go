package main

import "fmt"

func SlicesIndex[T comparable](s []T, v T) int {

	for i, val := range s {
		if val == v {
			return i
		}
	}
	return -1
}

func main() {
	// string turu ile calisma
	strings := []string{"apple", "banana", "cherry"}
	fmt.Println(SlicesIndex(strings, "cherry"))

	fmt.Println(SlicesIndex(strings, "apple"))
	fmt.Println(SlicesIndex(strings, "banana"))
	fmt.Println(SlicesIndex(strings, "cherry"))
	fmt.Println(SlicesIndex(strings, "elma")) // -1
	fmt.Println(SlicesIndex(strings, "cherry"))
	fmt.Println("----------------------------------")
	ints := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(SlicesIndex(ints, 1))
	fmt.Println(SlicesIndex(ints, 2))
	fmt.Println(SlicesIndex(ints, 19)) // -1
}
