# Estudos Go

Este repositório contém códigos de exercícios e práticas de lógica de programação em Go, baseados principalmente na documentação oficial da linguagem.

## 📚 Sobre

Repositório de estudos para aprendizado da linguagem Go, seguindo o [A Tour of Go](https://go.dev/tour/). Cada exemplo é um programa independente, com seu próprio `main.go`.

## 📂 Estrutura

Os exemplos ficam em `tour/`, separados pelas seções do Tour. O número no nome de cada pasta corresponde à página da lição no Tour (ex.: `02-flowcontrol/05-if` → https://go.dev/tour/flowcontrol/5).

```
tour/
├── 01-basics/        # Packages, variables, and functions
└── 02-flowcontrol/   # Flow control statements: for, if, else, switch and defer
```

### 01 · Basics — [go.dev/tour/basics](https://go.dev/tour/basics/1)

| Pasta | Conteúdo |
|-------|----------|
| [00-hello](tour/01-basics/00-hello/main.go) | Hello World |
| [02-imports](tour/01-basics/02-imports/main.go) | Imports e pacotes `math` e `math/rand` |
| [04-funcoes](tour/01-basics/04-funcoes/main.go) | Declaração e uso de funções |
| [05-funcoes-continuacao](tour/01-basics/05-funcoes-continuacao/main.go) | Parâmetros consecutivos do mesmo tipo |
| [06-funcao-multipla](tour/01-basics/06-funcao-multipla/main.go) | Múltiplos retornos |
| [07-valores-retorno](tour/01-basics/07-valores-retorno/main.go) | Retornos nomeados |
| [09-variaveis-inicializador](tour/01-basics/09-variaveis-inicializador/main.go) | Variáveis com inicializadores |
| [10-declaracao-curta](tour/01-basics/10-declaracao-curta/main.go) | Declaração curta (`:=`) |
| [12-valores-zero](tour/01-basics/12-valores-zero/main.go) | Valores zero |
| [13-conversores-tipo](tour/01-basics/13-conversores-tipo/main.go) | Conversão de tipos |
| [14-inferencia-tipo](tour/01-basics/14-inferencia-tipo/main.go) | Inferência de tipos |
| [16-constantes](tour/01-basics/16-constantes/main.go) | Constantes numéricas |
| [exercicio-funcao](tour/01-basics/exercicio-funcao/main.go) | Exercício prático com funções (em andamento) |

### 02 · Flow control — [go.dev/tour/flowcontrol](https://go.dev/tour/flowcontrol/1)

| Pasta | Conteúdo |
|-------|----------|
| [01-for](tour/02-flowcontrol/01-for/main.go) | Laço `for` |
| [03-while](tour/02-flowcontrol/03-while/main.go) | `for` como `while` |
| [05-if](tour/02-flowcontrol/05-if/main.go) | `if` |
| [06-if-declaracao-curta](tour/02-flowcontrol/06-if-declaracao-curta/main.go) | `if` com declaração curta |
| [07-if-else](tour/02-flowcontrol/07-if-else/main.go) | `if` / `else` |
| [09-switch](tour/02-flowcontrol/09-switch/main.go) | `switch` |
| [10-switch-ordem-avaliacao](tour/02-flowcontrol/10-switch-ordem-avaliacao/main.go) | Ordem de avaliação do `switch` |
| [12-defer](tour/02-flowcontrol/12-defer/main.go) | `defer` |

## 🚀 Como Executar

Na raiz do repositório (onde está o `go.mod`), aponte o `go run` para a pasta do exemplo:

```bash
go run ./tour/01-basics/00-hello
go run ./tour/02-flowcontrol/05-if
```

Para compilar todos os exemplos de uma vez e checar se nada quebrou:

```bash
go build ./...
```

Para adicionar uma nova lição, crie uma pasta com o número da página do Tour e um `main.go` com `package main`.

## 📖 Referências

- [Documentação Oficial do Go](https://go.dev/doc/)
- [A Tour of Go](https://go.dev/tour/)
