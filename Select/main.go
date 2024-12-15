/*
package main

import (

	"fmt"
	"time"

)

	func main() {
		c1 := make(chan string, 1)
		c2 := make(chan string, 1)
		/*go func() {
			//		time.Sleep(500 * time.Millisecond)
			c1 <- "one"
		}()
		go func() {
			//time.Sleep(2 * time.Second)
			c2 <- "two"
		}()
		c1 <- "one"
		c2 <- "two"
		for i := 0; i < 2; i++ { // for dongusu 2 kere calisir ******
			select {
			case msg := <-c1:
				fmt.Println("received", msg)

			case msg := <-c2:
				fmt.Println("received", msg)

			default:
				fmt.Println("default: no message received")
			}

		}
	}
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	// 1 saniye sonra c1'e veri gönderiyoruz
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "message from c1"
	}()

	// Bu select sürekli tekrar ediyor
	for i := 0; i < 5; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		default:
			fmt.Println("no message received")
		}
		time.Sleep(300 * time.Millisecond)
	}
}
