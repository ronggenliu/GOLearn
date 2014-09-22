package main

import (
	"os"
	"text/template"
)

type Person struct {
	Name string
	Age  uint
}

func main() {
	lrg := Person{Name: "ronggenliu", Age: 32}
	a := "123"
	tmpl, err := template.New("test").Parse("{{/* comments */}}{{range $a}{{.Name}} is {{.Age}} old!{{end}}")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, lrg)
	if err != nil {
		panic(err)
	}
}
