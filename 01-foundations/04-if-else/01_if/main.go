package main

import "fmt"

func main() {

	var idade int
	fmt.Println("Entre com a sua idade: ")
	fmt.Scanf("%d", &idade)

	if idade > 18 {
		fmt.Println("Perfeito!")
		fmt.Println("Beba a vontade!")
	}

	if idade < 17 {
		fmt.Println("Vishhhh")
		fmt.Println("Melhor ir tomar leite")
	}

	else {
		fmt.Println("Mostre sua identidade, pra gente verificar?")
	}
