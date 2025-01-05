package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		messages := make(chan string)
		//signals := make(chan bool)
		msg := "hi"
		select {
		case messages <- msg:
			fmt.Println("received message", msg)
		default:
			fmt.Println("no message received")
	*/
	/*messages := make(chan string)
	select {
	case msg := <-messages:
		fmt.Println("received message:", msg)
	default:
		fmt.Println("no message received")

	msg := make(chan string)
	select {
	case msg <- "Hello":
		fmt.Println("Message sent")
	default:
		fmt.Println("Message not sent")
	}}*/

	messages := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		messages <- "ping"
	}()
	for i := 0; i < 3; i++ {
		select {
		case msg := <-messages:
			fmt.Println("received:", msg)
		default:
			fmt.Println("no message yet")
		}
		time.Sleep(1 * time.Second)
	}
}
