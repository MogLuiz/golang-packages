package main

import (
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	api := http.Client{Timeout: time.Second}
	res, err := api.Get("http://google.com")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}
