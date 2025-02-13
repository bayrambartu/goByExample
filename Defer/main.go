package main

import (
	"fmt"
	"os"
)

func main() {

	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)

}
func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}
func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data") // dosyaya yazma
}
func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()
	if err != nil {
		panic(err)
	}

	/* defer fmt.Println("World")
		defer fmt.Println("One")
		defer fmt.Println("Two")
		fmt.Println("Hello")
		myDefer()

	}
	func myDefer() {
		for i := 0; i < 5; i++ {
			defer fmt.Print(i)

		}
	}*/
}
