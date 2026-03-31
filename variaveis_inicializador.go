package main 

import "fmt"

var i, j int = 1, 2  			// Inicializadores, 1 por variável
								// As variáveis i e j são inteiros, 1 e 2 respectivamente

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j , c, python, java)
}