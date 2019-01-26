// Package Declaration
package main

// Formatting for input and output
// "fmt" string literal -> type of expressions
import "fmt"

func main()  {
  defer func()  {
    str := recover()
    fmt.Println(str)
  }()
  // Panic indicate program error
  panic("PANIC")

  // defer will run the function after the function completes
  defer second()
  first()
}

func first()  {
  fmt.Println("first")
}

func second()  {
  fmt.Println("second")
}
