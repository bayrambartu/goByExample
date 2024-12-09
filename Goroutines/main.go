package main

import (
	"fmt"
	"sync"
)

/*
	func f(from string) {
		for i := 0; i < 3; i++ {
			fmt.Println(from, ":", i)
		}
	}
*/
func f(wg *sync.WaitGroup, from string) {
	defer wg.Done() // fonksiyon bitince bir işlem çalıştır
	for i := 0; i < 5; i++ {
		fmt.Println(from, ":", i)
	}
}
func main() {
	/*f("direct")       // bu kodu calistirmak için main fonksiyonunu bekler
	go f("goroutine") // goroutine ile fonksiyon cagrisi(asenkron)

	go func(msg string) { // goroutine ile anonim func (Lambda/Closure)
		fmt.Println(msg)
	}("going") // msg tanımlanır anonim func()

	time.Sleep(time.Second * 2)
	fmt.Println("done")*/

	var wg sync.WaitGroup
	wg.Add(2)
	go f(&wg, "Goroutine1")
	go f(&wg, "Goroutine2")
	// f(&wg, "Goroutine3")

	wg.Wait()
	fmt.Println("done")
}
