package main

import (
	"fmt"
)

func main() {
	a := "a"
	b := "à"
	c := "𐩔"

	fmt.Printf("a --> %v\n", a)
	fmt.Printf("b --> %v\n", b)
	fmt.Printf("c --> %v\n", c)

	d := []byte(a)
	e := []byte(b)
	f := []byte(c)

	fmt.Printf("d --> %v\n", d)
	fmt.Printf("e --> %v\n", e)
	fmt.Printf("f --> %v", f)
}
