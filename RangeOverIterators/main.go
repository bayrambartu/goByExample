package main

import (
	"fmt"
	"iter"
)

// Collect: Bir iteratör fonksiyonunu alır ve tüm değerleri bir dilime (slice) toplar
func Collect[T any](seq iter.Seq[T]) []T {
	var result []T
	seq(func(val T) bool {
		result = append(result, val) // Her öğeyi dileme ekle
		return true                  // Devam et
	})
	return result
}

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 1, 1

		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func main() {
	// Listeyi oluştur ve elemanları ekle
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)

	// Listeyi All() fonksiyonu ile iterasyona al ve yazdır
	for e := range lst.All() {
		fmt.Println(e)
	}

	// Listeyi Collect fonksiyonu ile dilime topla
	all := Collect(lst.All()) // Collect fonksiyonunu kullan
	fmt.Println("all:", all)  // [10 13 23]

	// Fibonacci sayıları ile çalış
	for n := range genFib() {
		if n >= 10 {
			break
		}
		fmt.Println(n)
	}
}
