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

	requests := make(chan int, 5) // 5 işlem için tampon kanal
	for i := 1; i <= 5; i++ {
		requests <- i

	}
	close(requests)
	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("işlem %d gerçekleşti \n", req)

	}
}
