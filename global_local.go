package main

import (
	"fmt"
)

// Declares a global variable
var x = 100

func main() {
	// Declares a local variable
	y := 20

	fmt.Printf("global x: %v\n", x)
	fmt.Printf("local y: %v\n\n", y)

	myFunc(y)
}

func myFunc(z int) {
	fmt.Printf("takes local variable y: %v\n", z)
	// Prints global variable
	fmt.Printf("global: %v\n\n", x)
}
