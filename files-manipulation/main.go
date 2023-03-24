package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("file.txt")
	if err != nil {
		panic(err)
	}

	//  -> Quando não sei os dados que estou escrevendo no arquivo utilizo o []byte
	size, err := f.Write([]byte("Escrevendo dados no arquivo"))

	//  -> Nesse caso eu sei que estou escrevendo uma string dai usei o método "WriteString"
	// size, err := f.WriteString("Escrevendo dados no arquivo")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes\n", size)
	f.Close()

	// Leitura
	file, err := os.ReadFile("file.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(file))

	// Leitura de pouco em pouco abrindo o arquivo (Conceito de stream)
	file2, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file2)
	buffer := make([]byte, 3)
	for {
		index, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:index]))
	}

	// Removendo arquivo
	err = os.Remove("file.txt")
	if err != nil {
		panic(err)
	}
}
