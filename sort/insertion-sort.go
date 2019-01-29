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
  elements = insertionSort(elements)
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

func insertionSort(elements []int) []int {
  fmt.Printf("**Insertion Sort**\n")

  key, y := 0, 0
  elementsLong := len(elements)
  for x:=1; x < elementsLong; x++ {
    key = elements[x]
    y = x-1
    for y >= 0 && key < elements[y] {
      elements[y+1] = elements[y]
      y -= 1
    }
    elements[y+1] = key
  }

  return elements
}
