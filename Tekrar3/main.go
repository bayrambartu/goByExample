package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		ticker := time.Tick(time.Second * 2)

		for i := 1; i <= 5; i++ {
			<-ticker // // 1 saniyede bir işlem yapılır.
			fmt.Printf("işlem %d gerçekleşti \n", i)
		}
		fmt.Println("tüm işlemler tamamlandı")
	*/
	// Rate limiter ile sıra yönetimi

	/*
		requests := make(chan int, 5) // 5 işlem için tampon kanal
		for i := 1; i <= 5; i++ {
			requests <- i

		}
		close(requests)
		limiter := time.Tick(1 * time.Second)

		for req := range requests {
			<-limiter
			fmt.Println("işlem %d gerçekleşti \n", req)

		} */

	// Patlama (Burst) Limitleri

	/*
		requests := make(chan int, 10)
		for i := 1; i <= 10; i++ {
			requests <- i

		}
		close(requests)
		limiter := make(chan time.Time, 3) // Burst limit : 3 işlem birden yapılabilir.

		for i := 0; i < 3; i++ {
			limiter <- time.Now()
		}
		go func() {
			for t := range time.Tick(1 * time.Second) {
				limiter <- t
			}
		}()

		// işlemlerin sırayla yapılmasına izin ver ..:
		for req := range requests {
			<-limiter
			fmt.Printf("işlem %d gerçekleşti \n", req)
		}
	*/

	// Patlama Limiti ve Sıralı Yönetim

	requests := make(chan int, 15)

	for i := 1; i <= 15; i++ {
		requests <- i
	}
	close(requests)

	limiter := make(chan time.Time, 5)

	for i := 0; i < 5; i++ {
		limiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(time.Second * 2) {
			limiter <- t
		}
	}()

	for req := range requests {
		<-limiter // tampondan işlem izni al
		fmt.Printf("İşlem %d gerçekleşti \n", req)
	}
}
