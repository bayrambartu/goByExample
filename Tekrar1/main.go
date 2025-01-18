package main

import (
	"fmt"
	"sync"
)

func main() {
	/*ch := make(chan int)
	go func() {
		ch <- 42
	}()
	val := <-ch
	fmt.Println("Kanal üzerinden alınan değer:", val)*/

	/*ch := make(chan int)
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
			fmt.Println("kanala gönderilen değer :", i)
		}
		close(ch)
	}()
	for val := range ch {
		fmt.Println("kanaldan alınan değer", val)
	}
	fmt.Println("tüm veriler alındı kanal kapandı")*/

	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i <= 5; i++ {
			ch <- i
			fmt.Println("sender 1 ", i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 6; i <= 10; i++ {
			ch <- i
			fmt.Println("sender 2 ", i)
		}
	}()

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			val := <-ch
			fmt.Println("receiver 1 ", val)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			val := <-ch
			fmt.Println("receiver 2 ", val)
		}
	}()

	// beklemek için ge goroutine kullanırız.
	go func() {
		wg.Wait()
		close(ch)
	}()

	wg.Wait()
	fmt.Println("done")

}
