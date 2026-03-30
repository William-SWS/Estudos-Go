// AO declarar uma variável sem especificar se tipo, o tipo da variável é inferida a partir do valor do lado direito

// var i int
// j := i // j é um int

// No entanto 

// i := 42           // int
// f := 3.142        // float64
// g := 0.867 + 0.5i // complex128
package main

import "fmt"

func main() {
	v := 3.142  // change me!
	fmt.Printf("v is of type %T\n", v)
}
