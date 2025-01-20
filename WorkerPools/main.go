package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	func worker(id int, jobs <-chan int, results chan<- int) {
		for j := range jobs {
			fmt.Println("worker", id, "started  job", j)
			time.Sleep(time.Second)
			fmt.Println("worker", id, "finished job", j)
			results <- j * 2
		}
	}
*/
/*func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}*/
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d, görevi işliyor: %d\n", id, job)
		time.Sleep(time.Second)
		results <- job * 2
	}
}
func main() {

	const numJobs = 10
	const numWorkers = 3

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	var wg sync.WaitGroup

	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)

	}
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)

	}()

	for result := range results {
		fmt.Printf("Sonuç: %d\n", result)
	}
	fmt.Println("tüm görevler tamam")
	/*const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	for j := 1; j <= numJobs; j++ {
		jobs <- j

	}
	close(jobs)
	for a := 1; a <= numJobs; a++ {
		result := <-results
		fmt.Printf("result:  %d\n", result)
	}*/
	/*
		const numJobs = 5
		jobs := make(chan int, numJobs)
		results := make(chan int, numJobs)

		for w := 1; w <= 3; w++ {
			go worker(w, jobs, results)
		}

		for j := 1; j <= numJobs; j++ {
			jobs <- j
		}
		close(jobs)

		for a := 1; a <= numJobs; a++ {
			<-results
		}*/
}
