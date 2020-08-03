package main

import (
	"fmt"
)

func main() {
	var i uint16
	// 65535 is the limit number of uint16 type
	// If we put 65536, it returns an error
	i = 65535
	// If we increment by one, it returns zero
	i++
	fmt.Println(i)
}
