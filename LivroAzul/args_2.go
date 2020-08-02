package main

import (
	"fmt"
	"os"
)

func main() {
	// Declares 2 variables how string, whithout value
	var stringArgs, separator string
	// for loop with range to iterate os.Args
	/* Range returns 2 values, index and value
	Here we want just value */
	for _, arg := range os.Args[1:] {
		stringArgs += separator + arg
		separator = " "
	}
	fmt.Println(stringArgs)
}
