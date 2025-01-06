package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		queue := make(chan string, 2)
		queue <- "one"
		queue <- "two"
		close(queue)

		for elem := range queue {
			fmt.Println(elem)
		} */

	/*
		queue := make(chan string, 2)
		queue <- "one"
		queue <- "two"
		//close(queue)
		// eğer bunu tanımlamazsak deadlock hatası alırız

		for elem := range queue {
			fmt.Println(elem)


		}
	*/

	queue := make(chan string, 5)
	go func() {
		queue <- "A"
		queue <- "B"
		queue <- "C"

		time.Sleep(1 * time.Second)
		close(queue)
	}()
	// Ana goroutine -->
	for elem := range queue {
		fmt.Println("Processing:", elem)

	}
	fmt.Println("All tasks processed!")
}
