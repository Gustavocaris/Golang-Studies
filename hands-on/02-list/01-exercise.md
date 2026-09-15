# 📋 O Desafio: Validador de Capacidade de Servidores

Você foi encarregado de criar um pequeno script de monitoramento para a infraestrutura.

Crie um arquivo chamado:

```text
01-exercise.go
```

e implemente os seguintes pontos.

## Requisitos do Código

### 1. Estrutura de Dados — Maps e Slices

* Crie um `map` chamado `servidores`.
* A **chave** deve ser o nome do servidor (`string`).
* O **valor** deve ser o consumo de memória em GB (`float64`).
* Adicione pelo menos **3 servidores** ao seu map.

Exemplo:

```go
servidores := map[string]float64{
	"srv-web":   4.5,
	"srv-db":    14.2,
	"srv-cache": 1.8,
}
```

---

### 2. Ponteiro e Atualização de Memória

Crie uma função chamada `adicionarConsumo` que receba:

* o **ponteiro** de um consumo de memória (`*float64`);
* um **valor adicional** (`float64`).

A função deve somar o valor adicional **diretamente no endereço de memória original**.

Exemplo de assinatura:

```go
func adicionarConsumo(consumo *float64, adicional float64) {
	// implementar
}
```

---

### 3. Conversão de Tipos e Slice

Crie um `slice` de inteiros chamado `alertas`, utilizando `make` com **capacidade inicial de 5**.

Exemplo:

```go
alertas := make([]int, 0, 5)
```

Em seguida:

1. Verifique se o servidor `"srv-db"` existe no seu `map`.
2. Se ele existir, utilize a função `adicionarConsumo` para adicionar **2.8 GB** ao consumo de memória.
3. Converta o novo valor de memória do `"srv-db"` de `float64` para `int`.

   * A conversão deve resultar no arredondamento para baixo, conforme o comportamento da conversão de `float64` para `int` em Go.
4. Adicione esse valor inteiro ao slice `alertas` utilizando `append`.

---

### 4. Exibição dos Resultados

Ao final do programa, imprima:

* o novo valor de memória do `"srv-db"`;
* o slice `alertas`;
* a capacidade (`cap`) do slice `alertas`.

Exemplo do que deve ser exibido:

```text
Novo consumo do srv-db: 17.0 GB
Alertas: [17]
Capacidade dos alertas: 5
```

> **Desafio:** tente implementar o exercício sem consultar uma solução pronta. O objetivo é praticar `maps`, `slices`, `make`, `append`, ponteiros, funções, verificação de existência em maps e conversão de tipos.
