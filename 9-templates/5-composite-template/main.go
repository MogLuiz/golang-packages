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

	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		tmp := template.Must(template.New("content.html").ParseFiles(templates...))
		err := tmp.Execute(response, Courses{
			{"Go", 40},
			{"Java", 10},
			{"Ruby", 90},
		})
		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":8282", nil)
}
