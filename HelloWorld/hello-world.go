// Package Declaration
package main

// Formatting for input and output
// "fmt" string literal -> type of expressions
import "fmt"

// Declare the type
type fn func(string)

const author string = "Keyne Oei"

// Init Function
func main() {
	fmt.Printf("Hello, World\n")

  CallOtherFunction(printHelloWorld, "Keyne")

  StringEx()

  printAuthor()
}

// Function Exercise
func printHelloWorld(name string) {
  var sentence string = "Hello, " + name + "\n"
  /*
  OR
  sentence := "Hello, " + name + "\n"
  no type is specified when using :=
  */
  fmt.Printf(sentence)
}

// Parameter of the function must be followed by the type
func CallOtherFunction(f1 fn, name string) {
  f1(name)
}

// Testing Constant Variable
func printAuthor() {
  fmt.Printf("\nAuthor --- " + author + "\n")
}

// String Exercise
func StringEx() {
  fmt.Println(len("Hello World")) // Print the length of the string
  fmt.Println("Hello World"[1]) // Print decimal of "e"
  fmt.Println("Hello " + "World")
}
