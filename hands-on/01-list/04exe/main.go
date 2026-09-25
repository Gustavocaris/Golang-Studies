// faça um programa que vende potes de Whey.
// Se o cliente escolher whey normal 1, o preço é 100 reais.
// Se o cliente escolher whey isolado 2, o preço é 150 reais.
// Se o cliente escolher whey hidrolisado 3, o preço é 200 reais.

package main

import "fmt"

func main() {
	var option int
	fmt.Print("Qual Whey voce gostaria de escolher ? 1 -> Normal 2 -> Isolado 3 -> hidrolisado")

	fmt.Scanf("%d", &option)

	switch option {
	case 1:
		fmt.Println("O preço do Whey Normal é 100 reais")
	case 2:
		fmt.Println("O whey Isolado custa 150 reais")
	case 3:
		fmt.Println("O whey hidrolisado, é mais carinho. Fica em 200 reais")
	default:
		fmt.Println("Opção invalida. Digite uma Opção valida")
	}

}
