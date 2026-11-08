package main

import (
	"fmt"
	"os"
)

func main() {
	// criando um arquivo
	arquivo, err := os.Create("arquivo.txt")
	if err != nil {
		fmt.Println("Erro ao criar o arquivo:", err)
		return
	}

	defer arquivo.Close()
}
