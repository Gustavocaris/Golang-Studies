/*
Uma struct permite criar um tipo próprio agrupando informações relacionadas.

*/

// Imagine um servidor:
type Server struct {
	Name   string
	IP     string
	CPU    float64
	Online bool
}

// Agora Server virou um tipo. Podemos criar:

server := Server{
	Name:   "server-01",
	IP:     "10.0.0.10",
	CPU:    45.5,
	Online: true,
}

// e acessar ;

fmt.Println(server.Name)
fmt.Println(server.IP)
fmt.Println(server.CPU)
fmt.Println(server.Online)