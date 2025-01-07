package main

import (
	"fmt"
	"time"
)

/*
	func delayedFunction() {
		fmt.Println("function executed after delay")
	}
*/
func simulateAPICall() <-chan string {
	result := make(chan string)

	go func() {
		time.Sleep(5 * time.Second)
		result <- "API Response : Success"
	}()
	return result
}
func main() {

	timeout := time.NewTimer(3 * time.Second)

	apiResult := simulateAPICall()

	select {
	case res := <-apiResult:
		fmt.Println("received:", res)

	case <-timeout.C:
		fmt.Println("Error API call timed out")
	}
	/*timer := time.NewTimer(time.Second)
	<-timer.C
	delayedFunction() */

	/*timer1 := time.NewTimer(time.Second)
	<-timer1.C
	fmt.Println("timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("timer 2 fired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("timer 2 stopped")
	}
	time.Sleep(2 * time.Second)*/

}
