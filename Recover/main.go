package main

import "fmt"

/*
	func sillySusan() {
		fmt.Println("silly susan called")
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("hackers have been thwarted")
			}
		}()

		panickingPeter()
		fmt.Println("silly susan finished")
	}

	func panickingPeter() {
		fmt.Println("panicking peter called")

		panic("oh no")
		fmt.Println("panicking peter finished")
	}
*/
func main() {
	/*
		fmt.Println("cascading panics")
		sillySusan()
		fmt.Println("end of main func")
	*/
	// cascading Panics

	/*
		// calli a panic
		fmt.Println("Panic! in the Go App")
		defer func() {
			fmt.Println("I am a deferred function")
		}()
		//panic("oh no")
		fmt.Println("end of main")
	*/

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recored Error:\n", r)
		}
	}()
	mayPanic()
	fmt.Println("After mayPanic()")

}
func mayPanic() {
	panic("a problem")
}
