// application entry point
package main

// lib import
import "fmt"

//keyword func --> declares a function
func main() {
	/*Println prints a line
	and returns byte_numbers an error_status */
	bytesNumber, errors := fmt.Println("Hello World!!")
	fmt.Println(bytesNumber, errors)
}
