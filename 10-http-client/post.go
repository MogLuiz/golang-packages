package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	api := http.Client{}

	jsonVar := bytes.NewBuffer([]byte(`{"name": "Luizin"}`))

	res, err := api.Post("http://google.com", "application/json", jsonVar)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	io.CopyBuffer(os.Stdout, res.Body, nil)
}
