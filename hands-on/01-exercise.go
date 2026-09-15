// STEP 1: Crie um map chamado servidores.

package main

import "fmt"

func adicionarConsumo(consumo *float64, adicional float64) {
	*consumo = *consumo + adicional
}

func main() {
	servidores := map[string]float64{
		"server-01": 7.5,
		"server-02": 9.2,
		"server-03": 5.7,
	}
	fmt.Println("Servidores contratados:")
	alertas := make([]int, 0, 5) // Usando make com capacidade 5

	// 1. Verifica se o servidor "server-01" existe usando o padrão comma ok
	consumoAtual, ok := servidores["server-01"]

	if ok {
		// 2. Passa o endereço de memória (&) de consumoAtual
		adicionarConsumo(&consumoAtual, 2.8)

		// Atualiza o valor alterado de volta no map
		servidores["server-01"] = consumoAtual

		// 3 e 4. Converte para int (arredonda para baixo) e adiciona ao slice com append
		novoValorInt := int(consumoAtual)
		alertas = append(alertas, novoValorInt)

		fmt.Printf("Servidor 'server-01' atualizado para %.1f GB\n", consumoAtual)
		fmt.Printf("Slice de alertas: %v | len: %d | cap: %d\n\n", alertas, len(alertas), cap(alertas))
	} else {
		fmt.Println("O servidor 'server-01' não foi encontrado no map.\n")
	}

	// O loop percorre o map e imprime um por linha
	for chave, valor := range servidores {
		fmt.Printf("- Servidor: %s | Consumo: %.1f GB\n", chave, valor)
	}
}
