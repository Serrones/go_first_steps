package main

import (
	"fmt"
)

type hotdog int

var b hotdog = 101

func main() {
	x := 10
	fmt.Printf("x: %v -- %T\n", x, x)

	fmt.Printf("b: %v -- %T\n", b, b)
	// We can attribute int(b) for x
	x = int(b)
	fmt.Printf("x: %v -- %T\n", x, x)
}
