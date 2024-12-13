/*package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("worker starting...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}
func main() {
	done := make(chan bool, 1)
	go worker(done)
	<-done // done kanalından bir şey gelene kadar bekler "blocikng"---kanaldan veri almayı bekler
	// kanaldan veri almadan programın bitirmesini önler !!
}
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

/*func worker(id int, done chan bool) {
fmt.Printf("Worker %d: working...\n", id)
time.Sleep(time.Second) // 1 saniye boyunca bekliyor
fmt.Printf("Worker %d: done\n", id)
done <- true // İşin bittiğini bildir */

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("worker %d done\n", id)

}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)

	}
	wg.Wait()
	/* done1 := make(chan bool, 1)
	done2 := make(chan bool, 1)
	done3 := make(chan bool, 1)

	go worker(1, done1)
	go worker(2, done2)
	go worker(3, done3)

	<-done1 // Worker 1'in bitmesini bekle
	<-done2 // Worker 2'nin bitmesini bekle
	<-done3 // Worker 3'ün bitmesini bekle

	fmt.Println("All workers are done!") */
}
