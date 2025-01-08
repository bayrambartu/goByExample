package main

import (
	"fmt"
	"sync"
)

/*
	func worker(id int) {
		fmt.Printf("Worker %d starting\n", id)

		time.Sleep(time.Second)
		fmt.Printf("Worker %d done\n", id)
	}
*/
/*
func worker(id int, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done() // bekleme sayacını 1 azaltır.go routşne tamamlandıgında cağırılır.
	fmt.Println("Worker ", id, " starting")
	time.Sleep(time.Second)
	results <- fmt.Sprintf("Worker %d completed ", id)
	fmt.Printf("Worker %d done \n", id)
} */

func worker(id int, wg *sync.WaitGroup, errors chan<- error) {
	defer wg.Done()

	if id == 2 { // Örneğin, Worker 2 hata veriyor
		errors <- fmt.Errorf("worker %d failed", id)
		return
	}

	fmt.Printf("Worker %d completed successfully\n", id)
}
func main() {
	var wg sync.WaitGroup
	errors := make(chan error, 5)

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg, errors)
	}

	wg.Wait()
	close(errors)

	for err := range errors {
		fmt.Println("Error:", err)
	}
	/*var wg sync.WaitGroup
	results := make(chan string, 5)
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, results, &wg)
	}
	wg.Wait()
	close(results)

	// close(result) olmadan range işlemi yapılamaz.
	for result := range results {
		fmt.Println(result)
	} */

	/*var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	wg.Wait()
	*/
}
