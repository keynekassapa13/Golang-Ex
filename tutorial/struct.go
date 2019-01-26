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

  c_2 := Circle{x: 3, y: 4, r: 5}
  fmt.Println(c_2.x, c_2.y, c_2.r)

  // * and & as the pointer
  // Call the struct
  fmt.Println(circlePoint(&c))

  // function in struct
  fmt.Println(c.area())

  // call other function in struct
  // interface implementation
  fmt.Println(totalArea(&c, &c_2))
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

// Interface
type Shapes interface {
  area() float64
}

func totalArea(shapes ...Shapes) float64 {
  var area float64
  for _, s := range shapes {
    area += s.area()
  }
  return area
}
