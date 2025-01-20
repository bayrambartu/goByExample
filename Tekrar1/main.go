package main

import (
	"fmt"
	"time"
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

	/*
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
		fmt.Println("done") */

	/*rand.NewSource(time.Now().UnixNano()) // random data
	ch1 := make(chan int)
	ch2 := make(chan int)

	var wg sync.WaitGroup*/

	// <Gondericiler (goroutine'ler)
	/*	wg.Add(2)
		go func() {
			defer wg.Done()
			for i := 0; i < 5; i++ {
				ch1 <- rand.Intn(100)
				time.Sleep(time.Millisecond * 100)
			}
		}()

		go func() {
			defer wg.Done()
			for i := 0; i < 5; i++ {
				ch2 <- rand.Intn(100)
				time.Sleep(time.Millisecond * 100)
			}
		}()

		//<Alıcılar(goroutine'ler)

		go func() {
			wg.Wait()
			close(ch1)
			close(ch2)
		}()

		for {
			select {
			case val, ok := <-ch1:
				if ok {
					fmt.Println("ch1'den alinan veri:", val)

				} else {
					ch1 = nil // kanal kapandi artik dinlenmeyecek.!
				}
			case val, ok := <-ch2:
				if ok {
					fmt.Println("ch2'den alinan veri:", val)
				} else {
					ch2 = nil
				}
			}
			if ch1 == nil && ch2 == nil {
				break
			}
		}
		fmt.Println("DONE") */

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		for i := 1; i <= 5; i++ {
			ch1 <- fmt.Sprintf("Kanal 1: Mesaj %d", i)
			time.Sleep(500 * time.Millisecond) // Daha hızlı gönderim
		}
		close(ch1)
	}()

	go func() {
		for i := 1; i <= 3; i++ {
			ch2 <- fmt.Sprintf("Kanal 2: Mesaj %d", i)
			time.Sleep(1 * time.Second)
		}
		close(ch2)
	}()

	for {
		select {
		case msg1, ok := <-ch1:
			if ok {
				fmt.Println(msg1)
			} else {
				ch1 = nil
			}
		case msg2, ok := <-ch2:
			if ok {
				fmt.Println(msg2)
			} else {
				ch2 = nil
			}
		}

		if ch1 == nil && ch2 == nil {
			break
		}
	}

	fmt.Println("DONE")

}
