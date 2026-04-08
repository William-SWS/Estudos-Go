# Estudos Go

Este repositório contém códigos de exercícios e práticas de lógica de programação em Go, baseados principalmente na documentação oficial da linguagem.

## 📚 Sobre

Repositório de estudos para aprendizado da linguagem Go, contendo exemplos práticos de conceitos fundamentais e exercícios de lógica de programação.

## 📂 Conteúdo

O repositório inclui exemplos práticos sobre:

- **Variáveis e Tipos**
  - [hello.go](hello.go) - Programa básico Hello World
  - [variaveis_inicializador.go](variaveis_inicializador.go) - Declaração de variáveis com inicializadores
  - [declaracao_curta.go](declaracao_curta.go) - Declaração curta de variáveis
  - [inferencia_tipo.go](inferencia_tipo.go) - Inferência de tipos
  - [valores_zero.go](valores_zero.go) - Valores zero das variáveis
  - [conversores_tipo.go](conversores_tipo.go) - Conversão entre tipos
  - [constantes.go](constantes.go) - Uso de constantes

- **Estruturas de Controle**
  - [for.go](for.go) - Versão antiga do exemplo de laço `for`
  - [while.go](while.go) - Versão antiga da simulação de `while`
  - [if_else.go](if_else.go) - Versão antiga do exemplo de `if/else`
  - [condicionais.go](condicionais.go) - Versão antiga de estruturas condicionais
  - [condicionais_curta_curacao.go](condicionais_curta_curacao.go) - Versão antiga de declaração curta em condicionais
  - [go tour/flowcontrol](go%20tour/flowcontrol) - Estrutura nova, organizada por exemplo, para os exercícios de controle de fluxo

- **Funções**
  - [funcoes.go](funcoes.go) - Declaração e uso de funções
  - [funcoes2.go](funcoes2.go) - Exemplos adicionais de funções
  - [funcao_multipla.go](funcao_multipla.go) - Funções com múltiplos retornos
  - [valores_retorno.go](valores_retorno.go) - Retorno de valores
  - [exercicio_funcao.go](exercicio_funcao.go) - Exercícios práticos com funções

- **Pacotes e Bibliotecas**
  - [math.go](math.go) - Uso do pacote `math`

## 🚀 Como Executar

Os arquivos antigos continuam executáveis individualmente:

```bash
go run hello.go
```

A estrutura nova para os exemplos de controle de fluxo fica em subpastas, então o comando recomendado passa a ser:

```bash
go run "./go tour/flowcontrol/for/main.go"
```

Cada pasta nova contém um `main.go` próprio, então você pode executar o exemplo desejado apontando para o arquivo correspondente.

## 📌 Estrutura Nova

Para evitar conflito entre vários `main()` no mesmo diretório, os exemplos do Tour foram reorganizados em pastas separadas:

- [go tour/flowcontrol/for/main.go](go%20tour/flowcontrol/for/main.go)
- [go tour/flowcontrol/while/main.go](go%20tour/flowcontrol/while/main.go)
- [go tour/flowcontrol/if_else/main.go](go%20tour/flowcontrol/if_else/main.go)
- [go tour/flowcontrol/condicionais/main.go](go%20tour/flowcontrol/condicionais/main.go)
- [go tour/flowcontrol/condicionais_curta_curacao/main.go](go%20tour/flowcontrol/condicionais_curta_curacao/main.go)

Essa organização mantém os exemplos na mesma área do workspace, mas separados o suficiente para não quebrar a compilação.

## 📖 Referências

- [Documentação Oficial do Go](https://go.dev/doc/)
- [A Tour of Go](https://go.dev/tour/)

