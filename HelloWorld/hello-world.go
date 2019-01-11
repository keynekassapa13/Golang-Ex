// Package Declaration
package main

// Formatting for input and output
// "fmt" string literal -> type of expressions
import "fmt"

// Declare the type
type fn func(string)

// Init Function
func main() {
	fmt.Printf("Hello, World\n")
  CallOtherFunction(printHelloWorld, "Keyne")
}

func printHelloWorld(name string) {
  fmt.Printf("Hello, " + name + "\n")
}

// Parameter of the function must be followed by the type
func CallOtherFunction(f1 fn, name string) {
  f1(name)
}
