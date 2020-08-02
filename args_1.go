package main

import (
	"fmt"
	"os"
)

func main() {
	// Declares 2 variables how string, whithout value
	var stringArgs, separator string
	// for loop to iterate os.Args
	for i := 1; i < len(os.Args); i++ {
		stringArgs += separator + os.Args[i]
		separator = " "
	}
	fmt.Println(stringArgs)
}
