// faça um programa que de bom dia,
// pergunte o nome da pessoa e reponde que é um prazer conhece-la sitadno o nome da pessoa.

package main

import "fmt"

func main() {
	fmt.Println("hello, what do you name?")

	var name string

	fmt.Scan("%s", &name)

	fmt.Printf("It's a pleasure to meet you, %s!\n", name)
}
