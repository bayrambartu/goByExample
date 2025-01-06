package main

import "fmt"

func main() {
	/*
		jobs := make(chan int, 5)
		done := make(chan bool)
		go func() {
			for {
				j, more := <-jobs
				if more { // kanal açık ve veri varsa ...
					fmt.Println("received job", j) // kanal açık oldugunda burayı işler
				} else {
					fmt.Println("received all jobs") // kanal kapalı oldugunda ise burayı işler ..
					done <- true
					return

				}
			}
		}()
		for j := 1; j < 3; j++ {
			jobs <- j
			fmt.Println("sent job", j)
		}
		close(jobs) // kanal kapatma
		fmt.Println("sent all jobs")

		<-done // done kanalından sinyal alana kadar ana goroutine'yi bekletir.

		_, ok := <-jobs
		fmt.Println("received more jobs:", ok) // false dönmesibin sebebi kanalın kapalı olması - ok bool değerdir.

	*/
	// use range
	jobs := make(chan int, 5)
	go func() {
		for j := 1; j <= 3; j++ {
			jobs <- j
			fmt.Println("Send job", j)
		}
		close(jobs)

	}()
	for j := range jobs {
		fmt.Println("received job", j)
	}
	fmt.Println("received all jobs")

	_, ok := <-jobs
	fmt.Println("received more jobs:", ok)
}
