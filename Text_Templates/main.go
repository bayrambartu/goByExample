package main

import (
	"os"
	"text/template"
)

func main() {
	t1 := template.New("t1")
	t1, err := t1.Parse("Value is {{.}} \n")
	if err != nil {
		panic(err)
	}
	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, map[string]string{"name": "John"})
	x := false
	t1.Execute(os.Stdout, x)
}
