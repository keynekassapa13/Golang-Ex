// Package Declaration
package main

// Formatting for input and output
// "fmt" string literal -> type of expressions
import "fmt"

func main() {
  // Odd Even Function
  // Experimenting with for loop and switch
  for i := 1; i <= 10; i++ {
    if i % 2 == 0 {
      fmt.Println(i, "even")  // Even
    } else {
      fmt.Println(i, "odd")   // Odd
    }
    fmt.Printf("Switch ---")
    switch i {
      case 0: fmt.Printf("Zero")
      case 1: fmt.Printf("One")
      case 2: fmt.Printf("Two")
      case 3: fmt.Printf("Three")
      case 4: fmt.Printf("Four")
      case 5: fmt.Printf("Five")
      default: fmt.Printf("Unknown Number")
    }
    fmt.Printf("\n\n")
  }

}
