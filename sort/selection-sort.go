package main

import (
  "fmt"
  "time"
  "math/rand"
)

func main() {
  elements := makeRandomList(2000)
  // fmt.Println(elements)

  start := time.Now()
  elements = selectionSort(elements)
  end := time.Now()
  elapsed := end.Sub(start)

  // fmt.Println(elements)
  fmt.Printf("\n**Time**\n")
  fmt.Println(elapsed)
}

func makeRandomList(rangeNum int) []int {
  elements := make([]int, rangeNum)
  for i := 0; i < rangeNum; i++ {
    elements[i] = (rand.Intn(rangeNum) + 1)
  }
  return elements
}

func selectionSort(elements []int) []int {
  fmt.Printf("**Selection Sort**\n")

  elementsLong := len(elements)
  min := 0
  for x, _ := range elements {
    min = x
    for y := x; y < elementsLong; y++ {
      if elements[y] < elements[min] {
        min = y
      }
    }
    elements[min], elements[x] = elements[x], elements[min]
  }

  return elements
}
