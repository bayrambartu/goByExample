package main

import (
	"os"
	"text/template"
)

func Create(name, t string) *template.Template {
	return template.Must(template.New(name).Parse(t))
}
func main() {
	/*
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
	*/

	/*
		// template.Must kullanimi
		// must kullanimi hataları otomatik olarak panic yapar ve programın durmasını sağlar.
		t2 := template.New("t2)")
		t2 = template.Must(t2.Parse("Value is {{.}} \n"))

		//	t2,err := t2.Parse("Value is {{.}} \n")
		//	if err != nil{
		//		panic(err)

		t2.Execute(os.Stdout, "Hello")
	*/

	/* BAK
	// struct map

	// Helper fonksiyon (sablon olusturma)
	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}
	t3 := Create("t3", "Name: {{.Name}}\n")

	// struct ile
	t3.Execute(os.Stdout, struct {
		Name string
	}{"Jane Doe"})

	// Map ile
	t3.Execute(os.Stdout, map[string]string{
		"Name": "Mick Mouse",
	}) */

	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	t3 := Create("t3",
		"{{if . -}} Yes {{else -}} No {{end}}\n")
	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "")

	t4 := Create("t4",
		"Languages: {{range .}}{{.}} {{end}}\n")

	//
	t4.Execute(os.Stdout, []string{"Go", "Rust", "C++", "C#"})

}
