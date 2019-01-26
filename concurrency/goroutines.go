// Package Declaration
package main

// Formatting for input and output
// "fmt" string literal -> type of expressions
import (
  "fmt"
  "time"
  "math/rand"
)

func main() {
  // call goroutines
  for i:=0; i<10; i++ {
    go f(i)
  }
  // goroutine doesnt wait for the function to complete
  // input to wait
  var input string
  fmt.Scanln(&input)
}

func f(n int) {
  for i:=0; i<10; i++ {
    fmt.Println(n, ":", i)
    amt := time.Duration(rand.Intn(250))
    time.Sleep(time.Millisecond * amt)
  }
}
