// Package Declaration
package main

// Formatting for input and output
// "fmt" string literal -> type of expressions
import "fmt"

func main() {
  fmt.Println(factorial(5))
}

// unit cannot store negative numbers (32 bits)
func factorial(x uint) uint {
  if (x == 0) {
    return 1
  }
  return x * factorial(x - 1)
}
