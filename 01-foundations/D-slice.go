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
