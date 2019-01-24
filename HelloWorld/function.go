// Package Declaration
package main

// Formatting for input and output
// "fmt" string literal -> type of expressions
import "fmt"

func main() {
  xs := []float64{98,93,77,82,83}
  fmt.Println(average(xs))

  x, y := f()
  fmt.Println(x)
  fmt.Println(y)
}

// parameter of float64
// return float64
func average(xs []float64) float64 {
  total := 0.0
  for _, v := range xs {
    total += v
  }
  return total / float64(len(xs))
}

// parameter of none
// return two integers
func f() (int, int) {
  return 5, 6
}
