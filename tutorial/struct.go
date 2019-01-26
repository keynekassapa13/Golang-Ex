// Package Declaration
package main

// Formatting for input and output
// "fmt" string literal -> type of expressions
import (
  "fmt"
  "math"
)

func main()  {
  c := Circle{x: 3, y: 4, r: 5}
  fmt.Println(c.x, c.y, c.r)
  fmt.Println(circlePoint(&c))
  fmt.Println(c.area())
}

type Circle struct {
  x, y, r float64
}

func circlePoint(c *Circle) float64 {
  return math.Pi * c.r * c.r
}

func (c *Circle) area() float64 {
  return math.Pi * c.r * c.r
}
