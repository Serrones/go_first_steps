package main

import (
	"fmt"
)

func main() {
	// Common string
	string1 := "Olá"

	// Respecting format
	string2 := `Olá
		meu
			precioso
					  mundo
	!!!`
	fmt.Println(string1)
	fmt.Println(string2)

	// Converting string to bytes
	stringBytes := []byte(string1)
	fmt.Println(stringBytes)

	// Here we iterate char by char
	for _, v := range string1 {
		fmt.Printf("%v = %T, %#U --- %#X\n", v, v, v, v)
	}

	fmt.Println(" ")

	// Here we iterate byte by byte
	for i := 0; i < len(string1); i++ {
		fmt.Printf("%v = %T, %#U --- %#X\n", string1[i], string1[i], string1[i], string1[i])
	}
}
