package main

import "net/http"

func main() {
	http.HandleFunc("/", SearchZipCode)
	http.ListenAndServe(":8080", nil)
}

func SearchZipCode(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Hello world!"))
}
