package main

import (
	"fmt"
	"time"
)

func worker(message chan string) {
	for msg := range message {
		fmt.Println("Worker received :", msg)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	message := make(chan string, 3)
	go worker(message)
	message <- "task 1"
	message <- "task 2"
	message <- "task 3"
	fmt.Println("sent all tasks")
	time.Sleep(5 * time.Second)

	/*message := make(chan string, 1)
	message <- "hello"
	//message <- "World"

	fmt.Println(<-message)
	fmt.Println(<-message)*/
	//fatal error: all goroutines are asleep - deadlock!
}
