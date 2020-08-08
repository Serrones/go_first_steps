package main

import (
	"fmt"
)

func main() {
	x := 1
	y := x << 1
	z := x >> 1
	fmt.Printf("%d\t%b\n%d\t%b\n%d\t%b\n", x, x, y, y, z, z)
}
