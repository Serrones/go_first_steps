package main

import (
	"fmt"
)

//Go is a static type language

var x int = 10

// Primitive Types
var a int
var b float64
var c string
var d bool

func main() {

	y := 1.5
	x = 5
	fmt.Printf("x: %v -- %T\n", x, x)
	fmt.Printf("y: %v -- %T\n\n", y, y)

	fmt.Printf("a: %v -- %T\n", a, a)
	fmt.Printf("b: %v -- %T\n", b, b)
	fmt.Printf("c: %v -- %T\n", c, c)
	fmt.Printf("d: %v -- %T\n\n", d, d)

	/* Zeros

	int --> 0
	float --> 0.0
	bool --> false
	str --> ""
	pointers, functions, interfaces, slices, channels, maps --> nill
	*/
}
