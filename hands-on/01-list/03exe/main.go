// faça um programinha que exiba o dobro de um número informado pelo usuário

package main

import (
	"fmt"
)

func main() {
	var x float64

	fmt.Println("insira um número:")
	fmt.Scanf("%f", &x)

	results := x * 2
	fmt.Println("O dobro do número informado é:", results)
}
