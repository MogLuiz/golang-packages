package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Essas Tags servem para mapear as propriedades da struct que muitas vezes estão em maiuscula para ser exportada e ao usar as tags: `json:"n"` fazemos o biding com a sua respectiva propriedade referenciada.
type Account struct {
	Number int `json:"n"`
	Amount int `json:"a"`
}

func main() {
	account := &Account{Number: 1, Amount: 100000}

	// Transaformando Struct em JSON. Quando eu uso o Marshal, eu estou guardando o valor para mim
	res, err := json.Marshal(account)
	if err != nil {
		fmt.Println(err)
		return
	}

	// O JSON é sempre retornado em bytes, por isso  precisamos converte-lo para string
	fmt.Println(string(res))

	// Quando eu uso o Encoder eu pego o valor já faço esse processo de serialização(transformação) entregando para alguem. Ex: stdOUT, algum arquivo, algum webserver
	err = json.NewEncoder(os.Stdout).Encode(account)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Voltando um JSON para o formato de struct
	jsonPuro := []byte(`{"n":1,"a":100000}`)
	// Utilizando o Recurso de tags para vincular os valor "n" e "a" aos valores da struct
	var accountX Account
	err = json.Unmarshal(jsonPuro, &accountX)
	if err != nil {
		fmt.Println(err)
		return
	}
	println(accountX.Amount)
}
