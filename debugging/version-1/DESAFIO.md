# 🐞 Debugging — Version 1: Monitor de Servidores

O arquivo `main.go` tem **bugs colocados de propósito**. Sua missão é fazer o programa
rodar e imprimir **exatamente** a saída esperada, **sem usar IA**: só lendo os erros,
o código e usando as ferramentas do Go.

## O que o programa deveria fazer

1. Ter uma lista de 3 servidores (nome, consumo em GB e limite em GB).
2. Adicionar **1 GB** extra de consumo no `srv-web`.
3. Imprimir cada servidor com seu consumo e limite.
4. Calcular o **consumo médio** de todos os servidores.
5. Listar os servidores que passaram de **80%** do limite.

## Saída esperada

```text
=== Monitor de Servidores ===
srv-web      7.0 /  8.0 GB
srv-db      15.0 / 16.0 GB
srv-cache    1.5 /  4.0 GB
Consumo médio: 7.83 GB
Servidores em alerta: 2
- srv-web: 87.5% de uso
- srv-db: 93.8% de uso
```

## Os bugs (3 camadas)

| Camada | Quantidade | Como aparece |
|---|---|---|
| 1. Compilação | 3 | `go run` nem executa: o compilador mostra arquivo:linha:coluna |
| 2. Runtime | 2 | Compila, mas quebra com `panic` e uma stack trace |
| 3. Lógica | 2 | Roda sem erro, mas a saída **não bate** com a esperada |

> Dica: as camadas aparecem uma de cada vez. Corrigir a camada 1 revela a 2, e assim por diante.

## Como rodar

```bash
cd debugging/version-1
go run .
```

---

## 🧰 Ferramentas para investigar

### 1. Ler o erro do compilador
O formato é `./main.go:LINHA:COLUNA: mensagem`. Leia a mensagem com calma: ela
quase sempre diz exatamente o que está errado.

### 2. Ler a stack trace de um `panic`
```text
panic: <motivo do erro>

goroutine 1 [running]:
main.algumaFuncao(...)
        .../main.go:NN      ← linha onde quebrou
main.main()
        .../main.go:NN      ← quem chamou a função
```
Leia de cima para baixo: a primeira linha com `main.go` é onde o problema aconteceu.

### 3. `go vet`
```bash
go vet .
```
Analisa o código e aponta erros comuns que compilam, mas são suspeitos.

### 4. Print debugging
O jeito mais simples: coloque `fmt.Printf("DEBUG x=%v\n", x)` nos pontos
suspeitos e compare o valor real com o que você esperava.
**Lembre de remover depois!**

### 5. Delve (o debugger do Go) 🔍
Instalação (uma vez só):
```bash
go install github.com/go-delve/delve/cmd/dlv@latest
```
> O binário vai para `~/go/bin`. Se o comando `dlv` não for encontrado, adicione
> `export PATH="$PATH:$HOME/go/bin"` no seu `~/.bashrc`.

Uso:
```bash
cd debugging/version-1
dlv debug .
```

Comandos principais dentro do Delve:

| Comando | O que faz |
|---|---|
| `break main.go:20` (ou `b`) | Cria um breakpoint na linha 20 |
| `break main.consumoMedio` | Breakpoint no início de uma função |
| `continue` (ou `c`) | Roda até o próximo breakpoint |
| `next` (ou `n`) | Executa a linha atual e vai para a próxima |
| `step` (ou `s`) | Entra dentro da função chamada na linha |
| `stepout` | Sai da função atual |
| `print x` (ou `p x`) | Mostra o valor de uma variável |
| `locals` | Mostra todas as variáveis locais |
| `args` | Mostra os argumentos da função atual |
| `restart` (ou `r`) | Reinicia o programa |
| `exit` (ou `q`) | Sai do Delve |

**No VS Code:** com a extensão oficial **Go** instalada, clique à esquerda do número
da linha para criar um breakpoint (bolinha vermelha) e aperte **F5**. Ela usa o
Delve por baixo e mostra as variáveis num painel lateral.

> 💡 Os bugs da **camada 3** são os ideais para treinar com o Delve: pare dentro
> das funções e veja se os valores das variáveis são o que você imaginava.

---

## ✅ Checklist

- [ ] Camada 1: o programa compila
- [ ] Camada 2: o programa roda até o fim sem `panic`
- [ ] Camada 3: a saída é **idêntica** à esperada
- [ ] Bônus: para cada bug, anote abaixo **o que era** e **por que acontecia**

## 📝 Minhas anotações

1.
2.
3.
4.
5.
6.
7.
