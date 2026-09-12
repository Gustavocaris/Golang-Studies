// Em Go não existem while ou do-while. O for assume todos os papéis de repetição.

// 1. Loop Tradicional com Contador (C-Style)
for i := 0; i < 3; i++ {
	fmt.Println("Contador:", i)
}

// 2. Loop estilo "while" (apenas com condição)
contador := 1
for contador <= 3 {
	fmt.Println("Rodando no estilo while:", contador)
	contador++
}

// 3. Loop Infinito (usado em daemons, workers ou escutas de rede)
// for {
//     // Roda infinitamente até encontrar um break ou return
// }

// FOR RANGE

// Para percorrer Slices, Arrays, Maps e Strings, o Go disponibiliza o operador range. Ele simplifica drasticamente a iteração trazendo índice/chave e valor.



