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
  elements = bubbleSort(elements)
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

func bubbleSort(elements []int) []int {
  fmt.Printf("**Bubble Sort**\n")

  elementsLong := len(elements)
  for x, _ := range elements {
    for y := 0; y < (elementsLong-x-1); y++ {
      if elements[y] > elements[y+1] {
        elements[y], elements[y+1] = elements[y+1], elements[y]
      }
    }
  }
  return elements
}
