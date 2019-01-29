package main

import (
  "fmt"
  "time"
  "math/rand"
  // "math"
)

func main() {
  elements := makeRandomList(2000)
  // fmt.Println(elements)
  fmt.Printf("**Merge Sort**\n")

  start := time.Now()
  elements = mergeSort(elements, 0, 0)
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

func mergeSort(elements []int, left int, right int) []int {

  if right == 0 {
    right = len(elements)
  }

  if left < (right - 1) {
    center := (left + right ) / 2

    mergeSort(elements, left, center)
    mergeSort(elements, center, right)
    temp_elements := merge(elements[left:center], elements[center:right])
    for x := left; x < right; x++ {
      elements[x] = temp_elements[x-left]
    }
  }

  return elements
}

func merge(left []int, right []int) []int {

  a, b := 0, 0
  result := make([]int, 0)
  for x := 0; x < (len(left) + len(right)); x++ {
    if a == len(left) {
      result = append(result, right[b:len(right)]...)
    } else if b == len(right) {
      result = append(result, left[a:len(left)]...)
    } else if left[a] < right[b] {
      result = append(result, left[a])
      a += 1
    } else {
      result = append(result, right[b])
      b += 1
    }
  }

  return result
}
