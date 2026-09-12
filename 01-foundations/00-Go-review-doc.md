# Go Foundations

Resumo dos fundamentos da linguagem Go.

## 1. Basic Structure

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

- `package main` → define o pacote principal.
- `import` → importa pacotes.
- `func main()` → ponto de entrada do programa.

---

## 2. Variables

### Explicit declaration

```go
var name string = "Gustavo"
var age int = 24
```

### Type inference

```go
var name = "Gustavo"
```

### Short declaration

Forma mais comum dentro de funções:

```go
name := "Gustavo"
age := 24
```

> `:=` pode ser usado apenas dentro de funções.

---

## 3. Basic Types

| Type | Example |
|---|---|
| `string` | `"Gustavo"` |
| `int` | `24` |
| `float64` | `1.80` |
| `bool` | `true` / `false` |

Example:

```go
name := "Gustavo"
age := 24
height := 1.80
isWorking := true
```

---

## 4. Constants

Use `const` for values that should not change.

```go
const port = 8080
```

Multiple constants:

```go
const (
    httpPort  = 80
    httpsPort = 443
)
```

---

## 5. Zero Values

Go automatically assigns a default value when a variable is declared without an initial value.

```go
var name string
var age int
var online bool
```

Results:

```text
string → ""
int    → 0
bool   → false
```

---

## 6. Type Conversion

Go does not automatically convert between numeric types.

```go
age := 24

height := float64(age)
```

Common conversions:

```go
int(value)
float64(value)
string(value)
```

> Converting a number to `string` is not the same as formatting the number as text. For that, packages such as `strconv` are commonly used.

---

## 7. Arrays

Fixed-size collection:

```go
servers := [3]string{
    "server-01",
    "server-02",
    "server-03",
}
```

Access by index:

```go
fmt.Println(servers[0])
```

Indexes start at `0`.

---

## 8. Slices

Dynamic collection. Very common in Go.

```go
servers := []string{
    "server-01",
    "server-02",
}
```

Add elements:

```go
servers = append(servers, "server-03")
```

---

## 9. Maps

Key-value structure:

```go
server := map[string]string{
    "name":   "server-01",
    "status": "online",
    "region": "sa-saopaulo-1",
}
```

Access a value:

```go
fmt.Println(server["name"])
```

---

## 10. Structs

Custom data structure used to group related fields.

```go
type Server struct {
    Name   string
    IP     string
    CPU    float64
    Online bool
}
```

Create a struct:

```go
server := Server{
    Name:   "server-01",
    IP:     "10.0.0.10",
    CPU:    45.5,
    Online: true,
}
```

Access fields:

```go
fmt.Println(server.Name)
fmt.Println(server.IP)
```

---

## Quick Reference

```text
Variable
    ↓
Basic Types
    ↓
Arrays / Slices / Maps
    ↓
Structs
    ↓
Functions
    ↓
Methods
    ↓
Interfaces
    ↓
Error Handling
    ↓
Goroutines / Channels
    ↓
HTTP / APIs
    ↓
Cloud / DevOps
```

### Useful Commands

```bash
# Check Go version
go version

# Initialize a module
go mod init github.com/USERNAME/REPOSITORY

# Run a Go program
go run main.go

# Format code
gofmt -w main.go

# Build the program
go build

# Run tests
go test
```