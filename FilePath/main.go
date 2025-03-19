package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {

	/*
		p := filepath.Join("dir1", "dir2", "filename")
		fmt.Println("p:", p)
		fmt.Println(filepath.Join("dir1//", "filename"))
		fmt.Println(filepath.Join("dir1/.../dir1", "filename"))

	*/
	p := "dir1/dir2/filename"
	fmt.Println("Dir:", filepath.Dir(p))
	fmt.Println("Base:", filepath.Base(p))

	fmt.Println(filepath.IsAbs("dir/file"))  // Relative Path
	fmt.Println(filepath.IsAbs("/dir/file")) // Absolute Paths

	filename := "config.json"
	ext := filepath.Ext(filename)
	fmt.Println(ext)
	fmt.Println(strings.TrimSuffix(filename, ext))

	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	fmt.Println("--")

	path1 := filepath.Join("home", "user", "documents", "file.txt")
	fmt.Println("Path:", path1)

	fmt.Println("Dir:", filepath.Dir(path1))
	fmt.Println("Base:", filepath.Base(path1))

	fmt.Println(filepath.IsAbs(path1))
	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	fmt.Println("Extension:", filepath.Ext(path1))
	fmt.Println("Filename without Extension:", strings.TrimSuffix(path1, filepath.Ext(path1)))

	rel, err = filepath.Rel("/home/user", "/home/user/documents/file.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("Relative Path:", rel)
}
