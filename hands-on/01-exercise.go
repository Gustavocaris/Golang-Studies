// STEP 1: Crie um map chamado servidores.

package main

import "fmt"

func main() {
	servidores := map[string]float64{
		"server-01": 7.5,
		"server-02": 9.2,
		"server-03": 5.7,
	}
	fmt.Println("Servidores contratados:")

	// O loop percorre o map e imprime um por linha
	for chave, valor := range servidores {
		fmt.Printf("- Ambiente: %s | Nome: %s\n", chave, valor)
	}
}

// STEP 2: Crie uma função chamada adicionarConsumo que receba:
// o ponteiro de um consumo de memória (*float64);
// um valor adicional (float64).

func adicionarConsumo(consumo *float64, adicional float64) {
	scan("input pela memoria")
}
