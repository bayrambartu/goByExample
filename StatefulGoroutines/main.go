package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)


	type readOp struct {
		key  int      // okunacak anahtar
		resp chan int // deger döndürmek için kanal
	}

	type writeOp struct {
		key  int       // güncellenecek olan anahtar
		val  int       // yeni deger
		resp chan bool // yazma islemi basarili mi ?
	}


func main() {

	var readOps uint64
	var writeOps uint64

	reads := make(chan readOp)
	writes := make(chan writeOp)

	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				write.resp <- true
			}
		}
	}()

	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	time.Sleep(time.Second)
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)

	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)

	/*
		var counter int
		var wg sync.WaitGroup
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				for j := 0; j < 1000; j++ {
					counter++ // DATA RACE

				}
				wg.Done()
			}()
		}
		wg.Wait()
		fmt.Println("COUNTER ..:", counter) */
/*
	var counter int64
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&counter, 1)

			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("COUNTER ..:", counter)
}

*/