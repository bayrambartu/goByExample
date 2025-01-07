package main

import (
	"fmt"
	"time"
)

func main() {
	// belli aralıkta düzenli olarak bir işlemi gerçekleştirmek için kullanırız. Timer'lardan farkı -> sürekli sinyal göndermesidir.
	/*ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped") */

	ticker := time.NewTicker(2 * time.Second)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Health check stopped.")
				return
			case d := <-ticker.C:
				fmt.Println("Performing health check at", d)
			}
		}
	}()
	time.Sleep(10 * time.Second)
	done <- true
	fmt.Println("Ticker stopped.")
}
