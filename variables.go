package main

import (
	"fmt"
)

func main() {
	//gopher := declares a variable
	x := 10
	y := "local variable"
	z := 10.5

	// Printf prints dynamic variables
	// %v is the values's variable
	// %T is the type's variable
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n", y, y)
	fmt.Printf("z: %v, %T\n\n", z, z)

	/* When you just just need modify
	the values's variable, use = */
	x = 20
	y = "tchau"

	/* you can modify a variable with :=
	but it needs declare at least one variable */
	z, w := 23.98, "gopher declares variable"

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n", y, y)
	fmt.Printf("w: %v, %T\n", w, w)
	fmt.Printf("z: %v, %T\n", z, z)

	/* you can declare a variable with
	bool value */
	valueTrue := 10 == 10
	valueFalse := 10 != 10

	fmt.Printf("valueTrue: %v, %T\n", valueTrue, valueTrue)
	fmt.Printf("valueFalse: %v, %T\n\n", valueFalse, valueFalse)

}
