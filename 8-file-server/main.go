package main

import (
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./public"))

	mux := http.NewServeMux()
	mux.Handle("/", fileServer)
	mux.HandleFunc("/blog", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Hello bloguinho!"))
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}
