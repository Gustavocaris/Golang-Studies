/*

O Slice é a estrutura de dados mais utilizada em Go. Ele funciona como um array dinâmico, mas por baixo dos panos é uma estrutura leve que aponta para um Array subjacente.

Um Slice é composto internamente por 3 componentes:

    Ponteiro: Indica onde começa o primeiro elemento do Array subjacente.

    Comprimento (len): O número de elementos atualmente presentes no Slice.

    Capacidade (cap): O número máximo de elementos que o Slice pode conter antes de precisar realocar memória.



Dica de Senior: Quando você usa append() e a capacidade (cap) estoura, o Go aloca automaticamente um novo Array com o dobro do tamanho e copia os dados.
Usar make([]tipo, len, cap) previne realocações desnecessárias quando você já sabe o tamanho aproximado dos dados.

*/

package main

import "fmt"

func main1() {
	servers := []string{
		"server-01",
		"server-02",
		"server-03",
	}
	fmt.Println(servers[2])

	// O slice não tem um tamanho fixo.
	servers = append(servers, "server-04")

}

// Slice de números
// Não precisa ser somente string.
func main2() {
	ports := []int{
		80,
		443,
		8080,
	}
	fmt.Println(ports[1])
}

// Até mesmo booleano

func main3() {
	status := []bool{
		true,
		true,
		false,
	}
	fmt.Println(status[0])
}
