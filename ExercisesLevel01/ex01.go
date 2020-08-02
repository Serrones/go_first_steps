/*
Utilizando o operador curto de declaração, atribua estes valores às variáveis
com os identificadores "x", "y", e "z".
    1. 42
    2. "James Bond"
    3. true
Agora demonstre os valores nestas variáveis, com:
    1. Uma única declaração print
    2. Múltiplas declarações print
*/
package main

import (
	"fmt"
)

func main() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Printf("x: %v | y: %v | z: %v\n", x, y, z)

	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}
