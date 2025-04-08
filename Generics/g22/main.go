package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

/*
	func Min(x, y float64) float64 {
		if x < y {
			return x
		}
		return y
	}
*/
func GMin[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}

// if wouldn't use package (go 1.18 +)
type Ordered interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64 | ~string
}

func GMinimum[T Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}
func main() {
	y := GMinimum(1, 2)
	fmt.Println(y)
	//fmt.Println(Min(1, 2))

	x := GMin[float64](2.232, 3.43)
	fmt.Println(x)

}
