package main

import (
	"fmt"
	"time"
)

func main() {
	/*request := make(chan int, 5)
	for i := 0; i < 5; i++ {
		request <- i

	}
	close(request)

	limiter := time.Tick(time.Millisecond * 200)

	for req := range request {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// en fazla 3 patlama kapasitesine sahip bir sınırlayıcı..
	// Anında işlenir aynı zamanda.
	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}
	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	burstyRequest := make(chan int, 5)
	for i := 0; i < 5; i++ {
		burstyRequest <- i
	}
	close(burstyRequest)

	for req := range burstyRequest {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}*/
	requests := make(chan int, 10)
	for i := 1; i <= 10; i++ {
		requests <- i
	}
	close(requests)

	limiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		limiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(time.Second * 2) {
			limiter <- t
		}
	}()
	for req := range requests {
		<-limiter
		fmt.Println("processing request", req, time.Now())
	}
}
