// Um array é uma coleção de elementos do mesmo tipo com tamanho fixo.
//Os índices começam em 0

package main

import "fmt"

func main() {
	servers := [3]string{
		"server-01",
		"server-02",
		"server-03",
	}
	fmt.Println(servers[2])
}
