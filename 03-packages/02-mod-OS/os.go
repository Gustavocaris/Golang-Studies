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

	// escrevendo no arquivo
	_, err = arquivo.WriteString("Olá, este é um arquivo de teste do GUGU.\n")
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo:", err)
		return
	}

	fmt.Println("Arquivo criado e escrito com sucesso!")

}
