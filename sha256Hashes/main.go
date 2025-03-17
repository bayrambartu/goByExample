package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

func main() {
	/*	s := "sha256 this string"
		h := sha256.New()  // create a hash object
		h.Write([]byte(s)) // convert to byte and add to hash
		bs := h.Sum(nil)

		fmt.Println("s-->", s)
		fmt.Println("h -->", h)
		fmt.Printf("%x\n -->", bs) */

	/*
		s := "hello world"
		h := sha256.New()
		h.Write([]byte(s))
		bs := h.Sum(nil)

		//Display in Hexadecimal and Base64 formats
		fmt.Printf("Hex: %x \n", bs)
		fmt.Printf("Base64: %x \n", base64.StdEncoding.EncodeToString(bs))

	*/
	s, err := os.Create("example.txt")
	s.WriteString("Bu bir test dosyasıdır.")
	s.Close()

	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("file not opened", err)
		return
	}
	defer file.Close()

	h := sha256.New()

	if _, err := io.Copy(h, file); err != nil {
		return
	}
	fmt.Printf("SHA-256 Hash: %x\n", h.Sum(nil))

	t := sha256.New()
	t.Write([]byte("hi"))
	t.Write([]byte("golang"))
	fmt.Printf("SHA-256 Hash: %x\n", t.Sum(nil))
}
