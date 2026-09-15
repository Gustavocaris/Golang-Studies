//Go é uma linguagem estaticamente tipada e fortemente tipada. Existem duas formas principais de declarar variáveis:

// Primeira forma VAR: Usada quando você precisa declarar uma variável sem inicializá-la imediatamente (ela assume o Zero Value) ou no escopo de pacote (fora de funções).

var age int // Declaração de variável do tipo int, sem inicialização (Zero Value é 0)

var name string // Declaração de variável do tipo string, sem inicialização (Zero Value é "")

var isActive bool // Declaração de variável do tipo bool, sem inicialização (Zero Value é false)



// Segunda forma, seria a curta (:=) — O Padrão em Funções
//Usada dentro de funções. O compilador infere o tipo com base no valor atribuído à direita.

count := 10        // O compilador infere como int
message := "Golang" // O compilador infere como string







