package main

import "fmt"

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}
	fmt.Printf(" Normal: %v\n", p)  // struct icerigi
	fmt.Printf("Detaylı: %+v\n", p) // alan adlari {x:1,y:2}
	fmt.Printf("Go formatı: %#v\n", p)
	fmt.Printf("Türü: %T\n", p)

	/*name := "ALice"
	age := 25
	fmt.Printf("My name is %s ! age %d \n", name, age)*/

	// struct

}
