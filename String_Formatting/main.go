package main

import "fmt"

/*
type point struct {
	x, y int
}*/

func main() {
	// struct

	/*p := point{1, 2}
	fmt.Printf(" Normal: %v\n", p)  // struct icerigi
	fmt.Printf("Detaylı: %+v\n", p) // alan adlari {x:1,y:2}
	fmt.Printf("Go formatı: %#v\n", p)
	fmt.Printf("Türü: %T\n", p)*/

	/*name := "ALice"
	age := 25
	fmt.Printf("My name is %s ! age %d \n", name, age)*/

	/*
		// boolean,integer binary format
		fmt.Printf("boolean: %t\n", false)
		fmt.Printf("Integer: %d\n ", 123)
		fmt.Printf("Binary: %b\n", 14)
		fmt.Printf("ASCII code: %c \n", 33)
		fmt.Printf("Hexadecimal: %x \n", 456)

	*/

	/*
		// Float
		fmt.Printf("normal format: %f \n", 89.4)
		fmt.Printf("Bilimsel format ( kucuk harf) : %e\n", 87.0)
		fmt.Printf("bilimsel format ( buyuk harf): %E \n", 12340000000.0)
	*/
	/*
		// String formatlama ve Pointer adresi
		fmt.Printf("String: %s \n", "Hello")
		fmt.Printf("Tirnak içinde: %q \n", "Hello")
		fmt.Printf("Hex: %x\n", "Range Over Iterators")

		x := 9
		var num int = 42
		fmt.Printf("Pointer: %p\n", &num)
		fmt.Printf("pointer:%p\n", &x)
	*/

	// Width & Precision

	fmt.Printf("Saga hizala: %6d %6d\n", 12, 345)
	fmt.Printf("Float (2 ondalık): |%6.2f|%6.2f|\n", 1.2, 3.45)
	fmt.Printf("Sol hizali: |%-6.2f|%-6.2f|\n", 1.2, 3.45)
	fmt.Printf("Metin saga hizali:|%6s|%6s|\n", "foo", "b")
	fmt.Printf("Metin sola hizali:|%-6s|%-6s|\n", "foo", "b")

	fmt.Printf("6 karakterlik alan, saga kaydir.: %6s %6s", "good", "luck")
}
