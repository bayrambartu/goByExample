package main

import "fmt"

// <- chan t (alma)
// chan <- t (gönderme)
// Bu fonksiyon, yalnızca veri gönderen bir kanalı alır.
// Gönderme yönü: chan<- string (sadece gönderme)
func ping(pings chan<- string, msg string) {
	pings <- msg
}
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings // bir kanaldan veri alınır msag tarafıdnan .. pings bosalır .
	pongs <- msg   // diger kanala veri gönderilir pongs'a veri aktarılır.
}
func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed message")
	pong(pings, pongs)

	fmt.Println(<-pongs)
	//fmt.Println(<-pings) --fatal error: all goroutines are asleep - deadlock!
}
