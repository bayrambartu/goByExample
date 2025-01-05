package main

import (
	"fmt"
	"time"
)

// timeouts bir işlem belli sürede tamamlanmazsa o işlemi iptal etme veya alternatif bir işlem yapma mekanizmasıdır.
func main() {
	/* c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1" // bu goroutine 2 sn uyuduktan sonra değeri kanala gönderir.
	}()

	// select'in amacı birden fazla kanalın verilerini dinlemek.
	select {
	case res := <-c1:
		fmt.Println(res)
		// Eğer 1 sn geçerse, time.After kanalı bir sinyal gönderir ve bu blok çalışır.
	case <-time.After(1 * time.Second):
		fmt.Println("timeout -- Zaman doldu")

	} */
	/*c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout")
	}*/
	/*ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	go func() {
		time.Sleep(3 * time.Second)
		ch1 <- "Data from ch1"
	}()
	go func() {
		time.Sleep(5 * time.Second)
		ch2 <- "Data from ch2"
	}()
	select {
	case res := <-ch1:
		fmt.Println("Gelen veri (ch1):", res)
	case res := <-ch2:
		fmt.Println("gelen veri (ch2): ", res)

	case <-time.After(10 * time.Second):

		fmt.Println("timeout-- hiçbir kanal hazır degil")

	} */

	/*ch4 := make(chan string, 1)
	go func() {
		time.Sleep(3 * time.Second)
		ch4 <- "Data 1"
	}()
	select {
	case res := <-ch4:
		fmt.Println(res)
	case <-time.After(2 * time.Second):
		fmt.Println("timeout")
	}*/
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	// Goroutine 1
	go func() {
		time.Sleep(3 * time.Second)
		ch1 <- "Data from ch1"
	}()

	// Goroutine 2
	go func() {
		time.Sleep(5 * time.Second)
		ch2 <- "Data from ch2"
	}()

	// İlk timeout (2 saniye)
	select {
	case res := <-ch1:
		fmt.Println("Gelen veri (ch1):", res)
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout (ch1 2 saniye)")
	}

	// İkinci timeout (4 saniye)
	select {
	case res := <-ch2:
		fmt.Println("Gelen veri (ch2):", res)
	case <-time.After(4 * time.Second):
		fmt.Println("Timeout (ch2 4 saniye)")
	}
}
