package main

import (
	"fmt"
	"slices"
)

func main() {
	// slices.Sort ( ordered -> int,float64,string-- artan düzende)

	strs := []string{"c", "a", "b"}
	slices.Sort(strs)
	fmt.Println("Strings:", strs) // O: Strings: [a b c]

	ints := []int{7, 2, 4}
	slices.Sort(ints)
	fmt.Println("Ints:", ints)

	// sıralı mı ? -> IsSorted

	v := slices.IsSorted(ints)
	fmt.Println("Is sorted?", v)

}
