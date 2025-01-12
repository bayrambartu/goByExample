package main

import (
	"fmt"
	"sync"
)

func main() {
	// Data Race---
	/*var ops int
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				ops++
			}
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("ops:", ops) */

	// ----
	/*var ops atomic.Uint64 // bu değişken atomik işlemler için kullanılacak ve tüm goroutine'ler tarafından paylaşılacak.

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ { // her goroutine ops sayacını 1000 kere arttırır.
				ops.Add(1) // atomik bir işlem olduğu için hiçbir iki işlem karışmaz.
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("ops:", ops.Load())
	*/

	// MUTEX

	var ops int
	var wg sync.WaitGroup
	var mu sync.Mutex
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				mu.Lock()
				ops++
				mu.Unlock()

			}
			wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println("ops:", ops)
}
