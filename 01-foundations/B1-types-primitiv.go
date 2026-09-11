// 🔵Para fechar o primeiro degrau, vamos olhar todos os tipos nativos que o Go oferece para armazenar dados brutos na memória.

/*
1. Números Inteiros (Com e Sem Sinal)

Em Go, você escolhe o tamanho do inteiro com base no limite de dados que precisa guardar (e quanta memória quer economizar):

    Com Sinal (aceitam positivos e negativos):

        int8: de -128 até 127

        int16: de -32.768 até 32.767

        int32: de -2.147.483.648 até 2.147.483.647

        int64: para números massivos (ex: timestamps em milissegundos)

        int: O padrão da linguagem (assume 32 ou 64 bits a depender da arquitetura do seu sistema operacional).

    Sem Sinal / Unsigned (apenas números >= 0):

        uint8, uint16, uint32, uint64, uint

        byte: É exatamente um apelido (alias) para uint8. Usado para manipular dados brutos, arquivos ou pacotes de rede.

2. Números Decimais (Ponto Flutuante)

    float32: Precisão simples.

    float64: O padrão absoluto em Go. Use sempre float64 para cálculos de precisão, coordenadas ou valores monetários simples.

3. Booleanos

    bool: Aceita estritamente true ou false. (Em Go, 1 ou 0 não convertem para booleano automaticamente).

4. Textos e Caracteres

    string: Sequência de texto imutável codificada em UTF-8.

    rune: É um alias para int32. Representa um único caractere Unicode (ex: 'A', '1', '🚀').
*/

// 🔵 Mão na Massa: Declarando os Tipos Primitivos





package main

import "fmt"

func main() {
	// Inteiros com e sem sinal
	var idade int = 25
	var portaRede uint16 = 8080
	var bufferByte byte = 255 // uint8

	// Decimais
	var preco float64 = 99.90

	// Booleano
	var ativo bool = true

	// Textos e Caracteres
	var nome string = "Gustavo"
	var emoji rune = '🚀' // Aspas simples para rune/caractere único

	fmt.Println("--- Mapeamento de Tipos Primitivos ---")
	fmt.Printf("Nome: %s (Tipo: %T)\n", nome, nome)
	fmt.Printf("Idade: %d (Tipo: %T)\n", idade, idade)
	fmt.Printf("Porta: %d (Tipo: %T)\n", portaRede, portaRede)
	fmt.Printf("Buffer: %d (Tipo: %T)\n", bufferByte, bufferByte)
	fmt.Printf("Preço: %.2f (Tipo: %T)\n", preco, preco)
	fmt.Printf("Status: %t (Tipo: %T)\n", ativo, ativo)
	fmt.Printf("Emoji Rune: %c - Código Unicode: %d (Tipo: %T)\n", emoji, emoji, emoji)
}

// O %T no fmt.Printf é uma instrução do Go que imprime o tipo exato daquela variável na tela.



























