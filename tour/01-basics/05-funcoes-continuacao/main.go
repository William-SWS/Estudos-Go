// Quando dois ou mais parâmetros nomeados consecutivos de uma função 
// compartilham o mesmo tipo, voce pode omitirt o tipo de todos menos o último

package main 

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(12, 23))
}