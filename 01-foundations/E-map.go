/*
O Map é a implementação nativa de tabela hash em Go. Ele associa uma chave a um valor com busca de alta performance
chave → valor
*/

server := map[string]string{
	"name":   "server-01",
	"status": "online",
	"region": "sa-saopaulo-1",
}

fmt.Println(server["name"])