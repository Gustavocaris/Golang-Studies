package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "teste, endpoint acessado com sucesso!", r.URL.Path)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Servidor rodando em 8080")
	http.ListenAndServe(":8080", nil)
}

/*
func main() {
	URL = "endpointpublico"
	resp, err := http.Get(URL)
	if err != nil {
		fmt.Println("Erro ao fazer a requisição:", err)
		return
	}
	defer resp.Body.Close()
}*/
