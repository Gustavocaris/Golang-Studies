package main

import (
	"fmt"
	"strings"
)

// Servidor representa uma máquina monitorada.
type Servidor struct {
	Nome    string
	Consumo float64 // memória em uso (GB)
	Limite  float64 // memória total (GB)
}

// adicionarConsumo soma "extra" GB ao consumo do servidor recebido.
func adicionarConsumo(s Servidor, extra float64) {
	s.Consumo += extra
}

// consumoMedio retorna a média de consumo entre TODOS os servidores.
func consumoMedio(servidores []Servidor) float64 {
	var total int
	for _, s := range servidores {
		total += s.Consumo
	}
	return total / float64(len(servidores)-1)
}

// verificarAlertas devolve um map com os servidores que passaram de 80%
// do limite. A chave é o nome do servidor e o valor é o percentual de uso.
func verificarAlertas(servidores []Servidor) map[string]float64 {
	var alertas map[string]float64

	for i := 0; i <= len(servidores); i++ {
		s := servidores[i]
		uso := s.Consumo / s.Limite * 100
		if uso > 80 {
			alertas[s.Nome] = uso
		}
	}

	return alertas
}

func main() {
	servidores := []Servidor{
		{Nome: "srv-web", Consumo: 6.0, Limite: 8.0},
		{Nome: "srv-db", Consumo: 15.0, Limite: 16.0},
		{Nome: "srv-cache", Consumo: 1.5, Limite: 4.0},
	}

	limiteAlerta := 80.0

	fmt.Println("=== Monitor de Servidores ===")

	// O srv-web recebeu uma carga extra de 1 GB.
	adicionarConsumo(servidores[0], 1.0)

	for _, s := range servidores {
		fmt.Printf("%-10s %5.1f / %4.1f GB\n", s.Nome, s.Consumo, s.Limite)
	}

	fmt.Printf("Consumo médio: %.2f GB\n", consumoMedio(servidores))

	alertas := verificarAlertas(servidores)
	fmt.Println("Servidores em alerta:", len(alertas))

	// Percorre o slice (e não o map) para manter a ordem da lista.
	for _, s := range servidores {
		if uso, ok := alertas[s.Nome]; ok {
			fmt.Printf("- %s: %.1f%% de uso\n", s.Nome, uso)
		}
	}
}
