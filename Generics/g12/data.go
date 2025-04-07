package main

// Generics ile veri yapıları

import "fmt"

type element[T any] struct {
	value T
	next  *element[T]
}

type LinkedList[T any] struct {
	head *element[T]
}

func (l *LinkedList[T]) Push(value T) {
	newElement := &element[T]{value: value}
	if l.head == nil {
		l.head = newElement
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = newElement
	}
}

func (l *LinkedList[T]) Print() {
	current := l.head
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
}
func main() {
	listInt := LinkedList[int]{}
	listInt.Push(1)
	listInt.Push(2)
	listInt.Push(3)
	listInt.Print()

	listString := LinkedList[string]{}
	listString.Push("hello")
	listString.Push("world")
	listString.Print()
}
