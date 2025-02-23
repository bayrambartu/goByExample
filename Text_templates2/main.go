package main

import (
	"os"
	"strings"
	"text/template"
)

func main() {
	tmplString := "Hello {{.}}"
	name := "Feury"
	//tmpl, err := template.New("Hello").Parse(tmplString)
	tmpl := template.Must(template.New("Hello").Parse(tmplString))
	/*if err != nil {
		panic("couldn't parse template")
	}*/
	err := tmpl.Execute(os.Stdout, name)
	if err != nil {
		panic("couldn't execute template")
	}
	tmplString = "{{if .}} Hello {{end}}"
	tmpl = template.Must(template.New("hello2").Parse(tmplString))
	_ = tmpl.Execute(os.Stdout, 1)
	names := []string{"Donald", "Bob", "Brodie"}
	tmplString = "{{range .}}Hello {{.}}\n{{end}}"
	tmpl = template.Must(template.New("range").Parse(tmplString))
	_ = tmpl.Execute(os.Stdout, names)
	tmplString = "{{range .}}Hello {{toUpper . }}\n{{end}}"
	FuncMap := map[string]interface{}{
		"toUpper": strings.ToUpper,
	}
	tmpl = template.Must(template.New("funcs").Funcs(FuncMap).Parse(tmplString))
	_ = tmpl.Execute(os.Stdout, names)
}
