package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	/*
		d1 := []byte("hello \ngo\n")
		err := os.WriteFile("/tmp/dat1", d1, 0644)
		check(err)
	*/

	/*
		f, err := os.Create("/tmp/dat2")
		check(err)
		defer f.Close()

		d2 := []byte{115, 111, 109, 101, 10}
		n2, err := f.Write(d2)
		check(err)
		fmt.Printf("wrote %d bytes\n", n2)

		n3, err := f.WriteString("writes\n")
		check(err)
		fmt.Printf("wrote %d bytes\n", n3)

		f.Sync()

		w := bufio.NewWriter(f)
		n4, err := w.WriteString("buffered\n")
		check(err)
		fmt.Printf("wrote %d bytes\n", n4)

		w.Flush()

		fmt.Println()
	*/
	file, err := os.Create("example.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// writing as text

	text := "Hello, this is a test file! \n Let's write some random text"
	_, err = file.WriteString(text)
	check(err)
	fmt.Println("Text succesfully written to file!")

	// Read file contents

	data, err := os.ReadFile("example.txt")
	check(err)
	fmt.Println("File Content: \n", string(data))

	// Use `bufio.Scanner` to read line by line
	file.Seek(0, 0)
	scanner := bufio.NewScanner(file)
	fmt.Println("Reading file line by line:")

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	check(scanner.Err())

	fmt.Println("-----------------------")

	file.Seek(0, 0)
	buffer := make([]byte, 10)
	n, err := file.Read(buffer)
	check(err)

	fmt.Println("First %d bytes %s \n", n, string(buffer))

}
