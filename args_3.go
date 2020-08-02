package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	/* Here we can use strings.Join
	instead for loop */
	fmt.Println(strings.Join(os.Args[1:], " "))
}
