# 🏋️ Lista de Exercícios — Revisão de Lógica em Go

**76 exercícios** para revisar os módulos **01 · Basics** e **02 · Flow control** deste repositório.

Todos usam **apenas** o que já foi estudado até aqui:

- variáveis, `var` e `:=`
- tipos básicos, valores zero, conversão e inferência de tipo
- constantes (tipadas e não tipadas)
- funções (parâmetros, múltiplos retornos, retornos nomeados)
- `for` (inclusive como `while` e loop infinito com `break`)
- `if` / `else` (inclusive com declaração curta)
- `switch` (com e sem condição, `fallthrough`, ordem de avaliação)
- `defer`

> ⚠️ **Regra de ouro:** nada de slices, maps, structs, ponteiros ou interfaces.
> O objetivo é treinar **lógica** com o básico. Se precisar "guardar" vários valores,
> calcule na hora ou imprima durante o loop.

---

## Como usar

1. Crie a pasta do exercício seguindo o padrão do repositório:

   ```bash
   mkdir -p desafios/01-ola-nome
   ```

2. Crie o `main.go` com `package main` e `func main()`.
3. Rode:

   ```bash
   go run ./desafios/01-ola-nome
   ```

4. No final, confira que nada quebrou:

   ```bash
   go build ./...
   ```

### Legenda de dificuldade

| Símbolo | Nível |
|---------|-------|
| ⭐ | aquecimento |
| ⭐⭐ | intermediário |
| ⭐⭐⭐ | desafio |

> As caixas `[ ]` servem de progresso. Marque `[x]` quando concluir.

---

## Módulo 1 — Fundamentos (variáveis, tipos e constantes)

- [ ] **1. Olá, nome** ⭐ — Crie `saudacao(nome string)` que imprime `Olá, <nome>!`.
- [ ] **2. Troca de valores** ⭐ — Troque duas variáveis `a` e `b` usando atribuição múltipla, sem variável auxiliar.
- [ ] **3. Valores zero** ⭐ — Declare `var` de `int`, `float64`, `bool` e `string` sem inicializar. Imprima com `%v` e `%T` e explique a saída em um comentário.
- [ ] **4. `var` vs `:=`** ⭐ — Reescreva o mesmo bloco usando só `var` e depois só `:=`. Comente por que `:=` não funciona fora de uma função.
- [ ] **5. Celsius → Fahrenheit** ⭐ — Receba `c float64` e retorne `F = C*9/5 + 32`. Cuidado com divisão inteira.
- [ ] **6. Média de três notas** ⭐ — As notas são `int`; retorne a média `float64` com conversão explícita.
- [ ] **7. Constantes de tempo** ⭐ — Defina `const` para segundos, minutos e horas. Calcule quantos segundos existem em `d` dias.
- [ ] **8. Constante grande** ⭐⭐ — Defina `const Big = 1 << 62`, atribua a `int64` e a `float64` e imprima os dois.
- [ ] **9. Inferência de tipos** ⭐ — Imprima com `%T` o tipo de `42`, `3.14`, `"go"`, `true` e `'a'`. Explique cada um em comentário.
- [ ] **10. Overflow** ⭐⭐ — Comece com `var x int64 = 1<<63 - 1`, some 1 e mostre o resultado. Explique o "wrap" em comentário.
- [ ] **11. Distância entre pontos** ⭐⭐ — Dados `x1, y1, x2, y2 float64`, calcule a distância com `math.Sqrt` e `math.Pow`.
- [ ] **12. IMC** ⭐ — Dados `peso` e `altura` (`float64`), retorne `peso / (altura*altura)` formatado com 2 casas.

---

## Módulo 2 — Funções

- [ ] **13. Soma e produto** ⭐ — `somaProduto(a, b int) (int, int)`.
- [ ] **14. Divisão inteira** ⭐ — Retorne o quociente **e** o resto de `a / b`.
- [ ] **15. Maior e menor** ⭐⭐ — `minMax(a, b, c int) (min, max int)` usando retornos nomeados.
- [ ] **16. Área e perímetro** ⭐ — Um retorno nomeado para cada, a partir de largura e altura.
- [ ] **17. Par ou ímpar** ⭐ — Retorne `bool`. Reutilize essa função nos próximos exercícios.
- [ ] **18. Maior de dois** ⭐ — Sem usar `math.Max`.
- [ ] **19. Refatoração** ⭐⭐ — Escreva um `main` com tudo inline e depois quebre a lógica em 3 funções pequenas de responsabilidade única.
- [ ] **20. Parâmetros do mesmo tipo** ⭐ — `retangulo(largura, altura float64)` usando a sintaxe agrupada `(largura, altura float64)`.
- [ ] **21. Saudação formal** ⭐ — `saudacao(titulo, nome string) string`, retornando algo como `"Dr(a). Ana"`.
- [ ] **22. Classificar nota** ⭐⭐ — Receba `nota float64` e retorne `A`, `B`, `C`, `D` ou `F` com `if/else`.
- [ ] **23. km → milhas** ⭐ — Use a constante `1.60934`.
- [ ] **24. Desconto** ⭐⭐ — Retorne `(valorFinal, percentual float64)` a partir de um valor e de um percentual.
- [ ] **25. `divmod` seguro** ⭐⭐⭐ — Retorne `(q, r int, ok bool)`. Se o divisor for `0`, retorne `ok = false` sem entrar em pânico.

---

## Módulo 3 — Laço `for`

- [ ] **26. Contar até 10** ⭐ — Imprima de 1 a 10.
- [ ] **27. Soma 1..N** ⭐ — Some todos os inteiros de 1 até `N`.
- [ ] **28. Tabuada** ⭐ — Imprima a tabuada de `N` de 1 a 10.
- [ ] **29. Fatorial** ⭐ — Calcule `N!` de forma iterativa.
- [ ] **30. Fibonacci** ⭐⭐ — Imprima os `N` primeiros termos. Sem slices: mantenha apenas dois valores (`a, b`) e atualize-os.
- [ ] **31. Collatz** ⭐⭐ — Conte quantos passos um número leva para chegar a 1 (par → `/2`, ímpar → `3n+1`).
- [ ] **32. Menor potência de 2** ⭐⭐ — Encontre a menor potência de 2 maior que `N` usando loop infinito com `break`.
- [ ] **33. Ímpares até 100** ⭐ — Use `continue` para pular os pares.
- [ ] **34. Primos até N** ⭐⭐⭐ — Loop aninhado com flag e `break` ao encontrar divisor.
- [ ] **35. MDC (Euclides)** ⭐⭐ — `for b != 0 { a, b = b, a%b }`.
- [ ] **36. Contar dígitos** ⭐ — Conte os dígitos de um inteiro positivo. Trate o caso `n == 0` (tem 1 dígito).
- [ ] **37. Inverter dígitos** ⭐⭐ — Ex.: `1234` → `4321`, usando `%` e `/`.
- [ ] **38. Soma dos dígitos** ⭐ — Ex.: `1234` → `10`.
- [ ] **39. Triângulo de asteriscos** ⭐⭐ — Desenhe um triângulo retângulo de altura `h` com loops aninhados.
- [ ] **40. Tabuleiro 8×8** ⭐⭐ — Alterne `#` e espaço para simular um tabuleiro de xadrez.
- [ ] **41. FizzBuzz** ⭐ — De 1 a 100: múltiplo de 3 → `Fizz`, de 5 → `Buzz`, de ambos → `FizzBuzz`.
- [ ] **42. Maior sequência de Collatz até N** ⭐⭐⭐ — Para cada número de 1 a `N`, calcule os passos e descubra qual tem a maior sequência.

---

## Módulo 4 — Condicionais (`if` / `else`)

- [ ] **43. Par ou ímpar (main)** ⭐ — Use a função do exercício 17 para decidir a mensagem.
- [ ] **44. Maior de três** ⭐ — Resolva com `if/else` encadeados.
- [ ] **45. Classificação de nota** ⭐ — Mesmas faixas do exercício 22, agora direto no `main`.
- [ ] **46. Ano bissexto** ⭐⭐ — Aplique a regra: divisível por 4, exceto por 100, salvo se por 400.
- [ ] **47. Triângulo válido** ⭐⭐ — Verifique se 3 lados formam um triângulo (cada lado menor que a soma dos outros dois) e classifique em equilátero, isósceles ou escaleno.
- [ ] **48. `if` com declaração curta** ⭐⭐ — Calcule um valor antes do `if` e use a variável apenas dentro dos blocos.
- [ ] **49. Faixas de IMC** ⭐ — `< 18.5` abaixo; `18.5–24.9` ideal; `25–29.9` sobrepeso; `>= 30` obesidade.
- [ ] **50. Desconto progressivo** ⭐⭐ — `< 100`: 0%; `100–499`: 5%; `500–999`: 10%; `>= 1000`: 15%.
- [ ] **51. Sinal do número** ⭐ — Negativo, zero ou positivo.
- [ ] **52. Imposto por faixa** ⭐⭐⭐ — Faixas simplificadas: até 2000 isento; 2000–3000: 10% sobre o excedente; 3000–4500: 20% sobre o excedente **somado** ao das faixas anteriores; acima de 4500: 30%. Cada faixa incide só sobre o que excede o limite anterior.

---

## Módulo 5 — `switch`

- [ ] **53. Dia da semana** ⭐ — `1..7` → nome do dia, com `default` para valor inválido.
- [ ] **54. Estação do ano** ⭐ — Mês → estação (hemisfério sul).
- [ ] **55. `switch` sem condição** ⭐⭐ — Use `switch { case idade < 12: ... }` para classificar faixas de idade.
- [ ] **56. `fallthrough`** ⭐⭐ — Demonstre que ele executa o próximo `case` **sem** testar a condição.
- [ ] **57. Ordem de avaliação** ⭐⭐ — Em um `switch` sem condição, use `case f1():` e `case f2():`, ambas verdadeiras. Prove com `fmt.Println` que só a primeira é executada.
- [ ] **58. Calculadora** ⭐⭐ — `switch op { case "+": ... }` para `+`, `-`, `*`, `/`.
- [ ] **59. Múltiplos valores por case** ⭐ — Classifique uma letra em vogal ou consoante com `case "a", "e", "i", "o", "u":`.
- [ ] **60. Menu simulado** ⭐⭐ — `menu(opcao int)` com 4 opções e `default`. Chame em loop até a opção `0`.

---

## Módulo 6 — `defer`

- [ ] **61. LIFO** ⭐ — Registre 3 `defer`s e observe a ordem de execução (último a entrar, primeiro a sair).
- [ ] **62. Recurso simulado** ⭐ — Crie `abrir()` e `fechar()`; use `defer fechar()` logo após abrir.
- [ ] **63. `defer` + retorno nomeado** ⭐⭐⭐ — Prove que um `defer` pode alterar o valor de retorno nomeado antes da função retornar.
- [ ] **64. `defer` em loop** ⭐⭐ — Coloque um `defer` dentro de um `for` e explique por que as mensagens só aparecem no fim da função.
- [ ] **65. Argumentos avaliados na hora** ⭐⭐ — Use `defer fmt.Println(i)` dentro de um loop e explique a saída.
- [ ] **66. Ordem completa** ⭐⭐ — Misture `defer` e `fmt.Println` e **preveja a saída no papel** antes de rodar.

---

## Desafios integradores (bônus)

- [ ] **67. Calculadora de IMC completa** ⭐⭐ — Funções + `if` + `switch` de categoria + formatação.
- [ ] **68. Conversor de unidades** ⭐⭐ — Menu `switch` para km↔milhas, °C↔°F e kg↔lb.
- [ ] **69. Caixa eletrônico** ⭐⭐⭐ — Decomponha um valor em notas (100, 50, 20, 10, 5, 2) com loop e divisão. Informe o resto que não pôde ser sacado.
- [ ] **70. Segundos → h:m:s** ⭐⭐ — Converta um total de segundos em horas, minutos e segundos com divisão e resto.
- [ ] **71. Juros compostos** ⭐⭐ — Para `meses` meses, imprima a evolução do montante (`montante *= 1 + taxa`).
- [ ] **72. Números perfeitos até N** ⭐⭐⭐ — Um número é perfeito se for igual à soma dos seus divisores próprios (ex.: 6, 28, 496).
- [ ] **73. Dígito verificador** ⭐⭐⭐ — Valide o primeiro dígito de um CPF **sem slices**: extraia os dígitos com `%` e `/`, aplique os pesos 10..2 e a regra `11 - (soma % 11)` (resultado `>= 10` vira `0`).
- [ ] **74. Decimal → binário** ⭐⭐ — Converta com `% 2` e concatenação de string, sem `strconv`.
- [ ] **75. Adivinhação simulada** ⭐⭐ — Sorteie o número secreto com `math/rand` e simule palpites em um loop com no máximo 5 tentativas, usando `break`.
- [ ] **76. Folha de pagamento** ⭐⭐⭐ — Salário bruto → desconto de INSS por faixa (função) + IR por faixa (`switch`) → salário líquido. Use faixas simplificadas e documente as alíquotas no código.

---

## ✅ Autoavaliação

Depois de terminar, você deve conseguir responder sem consultar:

1. Qual a diferença entre `var x int` e `x := 0`?
2. Por que `:=` não funciona no escopo de pacote?
3. O que acontece ao somar 1 em `int64` quando ele já está no valor máximo?
4. Quando usar retorno nomeado? Quais são os riscos?
5. Qual a diferença entre `for` com três cláusulas e `for` como `while`?
6. Quando `continue` é melhor que um `if` aninhado?
7. Como o `switch` decide qual `case` executar? E o que `fallthrough` muda?
8. Por que a ordem dos `defer` é invertida?
9. Quando os argumentos de um `defer` são avaliados?
10. Como um `defer` consegue alterar um retorno nomeado?

## 📅 Sugestão de ritmo

| Semana | Foco | Meta |
|--------|------|------|
| 1 | Módulos 1 e 2 | exercícios 1–25 |
| 1 | Módulo 3 | exercícios 26–42 |
| 2 | Módulos 4, 5 e 6 | exercícios 43–66 |
| 2 | Integradores | exercícios 67–76 |

Se o tempo for curto, priorize os exercícios ⭐ e ⭐⭐ e os integradores 67, 69 e 76.
