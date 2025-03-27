package main

import (
	"embed"
	"fmt"
)

// go:embed folder/single_file.txt
var fileString string

//go:embed folder/single_file.txt
var fileByte []byte

//go:embed folder/*
var folder embed.FS

func main() {
	fmt.Println("File as String:")
	fmt.Println(fileString)

	fmt.Println("\nFile as Bytes:")
	fmt.Println(string(fileByte))

	content1, err := folder.ReadFile("folder/file1.hash")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("\nContent of file file1.hash:")
		fmt.Println(content1)

	}
	fmt.Println(string(content1))

	content2, _ := folder.ReadFile("folder/file2.hash")
	fmt.Println("\nContent of folder file2.hash:")
	fmt.Println(string(content2))
}
