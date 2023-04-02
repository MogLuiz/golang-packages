package main

import (
	"context"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	// Inicializando um contexto vazio
	ctx := context.Background()

	// Definindo regra de timeout do meu contexto
	ctx, cancel := context.WithTimeout(ctx, time.Second)

	// Executando o cancelamento do meu contexto sempre que finalizar a execução do programa
	defer cancel()

	// Vinculando a minha request ao meu contexto, dessa forma essa request vai respeitar as regras do contexto.
	// Nesse caso, nosso contexto tem regra de timeout, se exceder esse timeout ele cancela minha chamada http
	request, err := http.NewRequestWithContext(ctx, "GET", "http://google.com", nil)
	if err != nil {
		panic(err)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	println(string(body))
}
