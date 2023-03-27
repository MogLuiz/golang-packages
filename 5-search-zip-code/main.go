package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, zipCode := range os.Args[1:] {

		// Buscando na api do viacep o endereço dinamicamente pelo CEP
		req, err := http.Get("http://viacep.com.br/ws/" + zipCode + "/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer a requisição: %v\n", err)
		}

		defer req.Body.Close()

		// Armazenando a resposta da chamada a API do viacep
		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler a resposta: %v\n", err)
		}

		var data ViaCEP
		// Transformando JSON da resposta em uma struct
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer o parse da resposta: %v\n", err)
		}

		// Criando TXT para armazenar os dados
		file, err := os.Create("city.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao criar arquivo: %v\n", err)
		}
		defer file.Close()

		// Armazenando os dados no arquivo TXT criado
		_, err = file.WriteString(fmt.Sprintf("CEP: %s, Localidade: %s, Bairro: %s, UF: %s", data.Cep, data.Localidade, data.Bairro, data.Uf))
		fmt.Printf("CEP: %s armazenado com sucesso!\n", data.Cep)
	}
}
