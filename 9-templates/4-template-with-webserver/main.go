package main

import (
	"net/http"
	"text/template"
)

type Course struct {
	Name  string
	Hours int
}

type Courses []Course

func main() {

	http.HandleFunc("/", webServerTemplate)
	http.ListenAndServe(":8282", nil)

}

func webServerTemplate(response http.ResponseWriter, request *http.Request) {
	tmp := template.Must(template.New("template.html").ParseFiles("template.html"))
	err := tmp.Execute(response, Courses{
		{"Go", 40},
		{"Java", 10},
		{"Ruby", 90},
	})
	if err != nil {
		panic(err)
	}
}
