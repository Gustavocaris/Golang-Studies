
// O switch substitui encadeamentos longos de if/else. Em Go, o break é automático ao final de cada case

status := 404

switch status {
case 200, 201: // Aceita múltiplos valores na mesma linha
	fmt.Println("Sucesso")
case 404:
	fmt.Println("Não Encontrado")
default:
	fmt.Println("Erro Interno")
}