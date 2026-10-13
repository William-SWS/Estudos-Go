package main

import "fmt"

func main() { 
	var i, j int = 1, 2
	k := 3 									// Aqui a variável k é declarada implicatamente como int 3 
	c, python, java := true, false, "no"   

	fmt.Println(i, j,k, c, python, java)
}