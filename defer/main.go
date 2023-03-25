package main

import "fmt"

// O "defer" ele atrasa a execução dessa linha de código. Depois que o compilador lêr tudo, dai ele executa a linha com o "defer"

func main() {
	// Dessa forma o "defer" vai adiar adiar a execução dessa linha. E só vai executar ela quando finalizar de ler as demais. Ou seja, o "defer" deixa essa execução por ultimo
	defer fmt.Println("Primeira linha")

	fmt.Println("Segunda linha")
	fmt.Println("Terceura linha")
}
