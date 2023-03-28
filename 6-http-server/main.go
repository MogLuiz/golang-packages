package main

import "net/http"

func main() {
	http.HandleFunc("/", SearchZipCodeHandler)
	http.ListenAndServe(":8080", nil)
}

func SearchZipCodeHandler(response http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	cepParam := request.URL.Query().Get("cep")
	if cepParam == "" {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write([]byte("Hello world!"))
}
