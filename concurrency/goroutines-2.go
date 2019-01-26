// Package Declaration
package main

// Formatting for input and output
// "fmt" string literal -> type of expressions
import (
  "fmt"
  "time"
)

func main() {
  var c chan string = make(chan string)

  go pinger(c)
  go printer(c)

  var input string
  fmt.Scanln(&input)
}

func pinger(c chan string) {
  for i:=0; ; i++ {
    if i == 3 {
      c <- "its three"
    }
    c <- "ping"
  }
}

func printer(c chan string)  {
  for {
    msg := <- c
    fmt.Println(msg)
    time.Sleep(time.Second * 1)
  }
}
