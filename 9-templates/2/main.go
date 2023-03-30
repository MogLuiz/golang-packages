package main

import (
	"os"
	"text/template"
)

type Course struct {
	Name  string
	Hours int
}

func main() {
	course := Course{Name: "GO", Hours: 12}

	tmp := template.Must(template.New("courseTemplate").Parse("Curso: {{.Name}} - Carga horária: {{.Hours}}"))

	err := tmp.Execute(os.Stdout, course)
	if err != nil {
		panic(err)
	}
}
