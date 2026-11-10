package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Country string `json:"country"`
}

func main() {

	// 1. Criar uma pessoa
	person := Person{
		Name:    "Gustavo",
		Age:     53,
		Country: "Brazil",
	}

	// 2 - > cconverter Struct para JSON
	jsonData, err := json.MarshalIndent(person, "", "    ")
	if err != nil {
		fmt.Println("Erro ao converter para JSON:", err)
		return
	}

	fmt.Println("JSON gerado:")
	fmt.Println(string(jsonData))

	// 3 Converter JSON para Struct
	var newPerson Person

	err = json.Unmarshal(jsonData, &newPerson)
	if err != nil {
		fmt.Println("Erro ao ler JSON:", err)
		return
	}

	fmt.Println("\nDados recuperados:")
	fmt.Println("Nome:", newPerson.Name)
	fmt.Println("Idade:", newPerson.Age)
	fmt.Println("País:", newPerson.Country)

}
