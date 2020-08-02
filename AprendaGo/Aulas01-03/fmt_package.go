package main

import (
	"fmt"
)

func main() {
	// Interpreted string literals
	x := "oi, bom dia.\ncomo vai?\n\t\"Tudo bem\", disse ele"
	fmt.Println(x)

	// Raw string literals
	y := `"oi, bom dia.\ncomo vai?\n\t\"Tudo bem\", disse ele"`
	fmt.Println(y)

	a := "oi"

	// Println adds new line
	fmt.Println(a)
	fmt.Println(a)

	// SPrint returns string instead print it
	b := fmt.Sprint(a)
	fmt.Println("b: ", b)
}
