package main

import (
	"fmt"
	"time"
)

func main() {

	/*
		//limiter := time.Tick(1 * time.Second)
		limiter := time.Tick(500 * time.Millisecond) //  sn'de 2 işlem
		for i := 1; i <= 5; i++ {

			<-limiter
			fmt.Printf("işlem %d yapıldı. Zaman : %s \n", i, time.Now().Format("15:04:05"))
		} */

	/*requests := make(chan int, 10)
	for i := 1; i <= 10; i++ {
		requests <- i
	}
	close(requests)

	limiter := make(chan time.Time, 2)

	for i := 0; i < 2; i++ {
		limiter <- time.Now()
	}
	go func() {
		for t := range time.Tick(2 * time.Second) {
			limiter <- t
			limiter <- t
		}
	}()
	for req := range requests {
		<-limiter
		fmt.Printf("İşlem %d yapıldı. Zaman: %s\n", req, time.Now().Format("15:04:05"))

	}*/

	/*
		requests := make(chan int, 10)
		for i := 1; i <= 10; i++ {
			requests <- i
		}
		close(requests)

		limiter := time.Tick(time.Second * 1)

		for req := range requests {
			<-limiter // izin al
			fmt.Printf("İşlem %d yapıldı. Zaman: %s\n", req, time.Now().Format("15:04:05"))

		}

	*/
	/*	requests := make(chan int, 10)
		for i := 1; i <= 10; i++ {
			requests <- i
		}
		close(requests)

		limiter := make(chan time.Time, 5)

		for i := 0; i < 5; i++ {
			limiter <- time.Now()
		}

		go func() {
			for t := range time.Tick(1 * time.Second) {
				limiter <- t
			}

		}()
		for req := range requests {
			<-limiter
			fmt.Printf("request: %d , time : %s \n", req, time.Now().Format("2006-01-02 15:04:05"))
		}

	*/

	requests := make(chan int, 10)

	// 10 işlem ekleyelim
	for i := 1; i <= 10; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(3 * time.Second)

	for req := range requests {
		select {
		case <-limiter: // İzin varsa işlemi gerçekleştir
			fmt.Printf("İşlem %d yapıldı. Zaman: %s\n", req, time.Now().Format("15:04:05"))
		case <-time.After(2 * time.Second): // 2 saniye içinde işlem yapılmazsa zaman aşımı
			fmt.Printf("İşlem %d zaman aşımına uğradı!\n", req)

		}
	}
}
