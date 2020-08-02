package main

import (
	"fmt"
)

// Create a hotdog type
type hotdog int

var b hotdog

func main() {
	fmt.Printf("b type is %T", b)
}
