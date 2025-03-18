package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	err := os.WriteFile("/tmp/dat", []byte("hello this is a test file"), 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("File created and file written")

	dat, err := os.ReadFile("/tmp/dat")
	check(err)
	fmt.Println(string(dat))

	f, err := os.Open("/tmp/dat")
	check(err)
	defer f.Close()

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Print(string(b1[:n1]))

	o2, err := f.Seek(6, io.SeekStart)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2[:n2]))

	_, err = f.Seek(-4, io.SeekEnd)
	check(err)
	b3 := make([]byte, 4)
	n3, err := f.Read(b3)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o2, string(b3[:n3]))

	/*
		o4, err := f.Seek(6, io.SeekStart)
		check(err)

		b4 := make([]byte, 2)
		n4, err := io.ReadAtLeast(f, b4, 2)
		check(err)
		fmt.Printf("%d bytes @ %d: %s\n", n4, o4, string(b4))
	*/
	// bufio
	f.Seek(0, io.SeekStart)
	r5 := bufio.NewReader(f)
	b5, _ := r5.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b5))
}
