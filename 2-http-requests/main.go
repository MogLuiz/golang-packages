package main

import (
	"io"
	"net/http"
)

func main() {
	request, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}

	// O "defer" ele atrasa a execução dessa linha de código. Depois que o compilador lêr tudo, dai ele executa a linha com o "defer"
	defer request.Body.Close()

	response, err := io.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}
	println(string(response))

}
