
// A expressão T(v) converte o valor v para o tipo T.

package main 

import (
	"fmt"
	"math"
)

func main() {
	var x , y int = 3, 4 
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}