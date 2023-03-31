package main

import (
	"os"
	"text/template"
)

type Course struct {
	Name  string
	Hours int
}

type Courses []Course

func main() {
	tmp := template.Must(template.New("template.html").ParseFiles("template.html"))

	err := tmp.Execute(os.Stdout, Courses{
		{"Go", 40},
		{"Java", 10},
		{"Ruby", 90},
	})
	if err != nil {
		panic(err)
	}
}
