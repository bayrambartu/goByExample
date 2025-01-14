package main

import (
	"fmt"
	"sync"
)

// amaç container içersinde birden fazla counter'i eşzamanlı olarak güncelleyebilmek.
// aynı anda çalışan goroutineler bu sayaçları sync.Mutex kullanrak güvenli hale getirmeli.
// sync.Mutex paylaşılan verilere aynı anda birden fazla goroutin'in erişmesini önleyen bir kilit mekanizmasıdır.
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock() // fonk sonun kilit serbesttir.
	c.counters[name]++
	// veirlen anahtarın sayacaını arttırır.
}

func main() {
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}
	var wg sync.WaitGroup

	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}
	wg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)
	wg.Wait()
	fmt.Println(c.counters)

	/*var ops int
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 10; j++ {
				mu.Lock()
				ops++
				mu.Unlock()
			}
			wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println(ops)

	*/

}
