package main

import "fmt"

/*
c := make(chan int)
c <- 10 // 10 degerini kanala gönderir

x := <-c // kanaldan veri al ve x'e ata
*/
/*func worker(c chan int) {
	for value := range c {
		fmt.Println(value)
	}
}*/

func main() {
	/*x := make(chan int)
	go worker(x)
	x <- 10
	x <- 20
	close(x)*/
	/*c := make(chan int)

	go func() {
		c <- 3
	}()
	values := <-c
	fmt.Println(values)*/

	/*c := make(chan int)
	c <- 1
	c <- 2
	c <- 3
	fmt.Println(<-c) //1
	fmt.Println(<-c) //2
	fmt.Println(<-c) //3

	fmt.Println(c) //0x140000b6000
	fmt.Println(c) //0x140000b6000
	fmt.Println(c) //0x140000b6000
	// eğer adresler farklı olsaydı bellek yönetimini zorlaştırır ve performansı ciddi şekilde düşürürdü
	// go çalışa zamanı,kanal adreslerini sabit tutarak bu tür bellek sorunlarını önler
	*/
	done := make(chan bool)
	go func() {
		fmt.Println("Goroutine World")
		done <- true
	}()
	<-done // done kanalından bir değer gelene kadar bekler
	fmt.Println("main goroutine  continues")

}
