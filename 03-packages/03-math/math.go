package main

import (
	"fmt"
	"math"
)

func acessPi() {
	fmt.Printf("PI arredondado (8 casas decimais): %.8f\n", math.Pi)
}

func acessEuler() {
	fmt.Printf("Euler arredondado (8 casas decimais): %.8f\n", math.E)
}

func main() {
	acessPi()
	acessEuler()
}
