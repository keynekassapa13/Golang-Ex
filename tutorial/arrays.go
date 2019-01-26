// Package Declaration
package main

// Formatting for input and output
// "fmt" string literal -> type of expressions
import "fmt"

func main() {
  var x [5]float64
  x[0] = 98
  x[1] = 93
  x[2] = 77
  x[3] = 82
  x[4] = 83
  var total float64 = 0
  for i := 0; i < 5; i++ {
       total += x[i]
  }
  fmt.Println(total / 5)

  // Straight num declaration
  arr := []float64{1,2,3,4,5}

  fmt.Println(arr)

  // Slicing
  y := make([]float64, 5, 10)
  y = arr[0:3]

  fmt.Println(y)

  // Maps -- similar to dictionary
  z := make(map[string]int)
  z["key"] = 10
  fmt.Println(z["key"])

  // More Maps
  elements := make(map[string]string)
  elements["H"] = "Hydrogen"
  elements["He"] = "Helium"
  elements["Li"] = "Lithium"
  elements["Be"] = "Beryllium"
  elements["B"] = "Boron"
  elements["C"] = "Carbon"
  elements["N"] = "Nitrogen"
  elements["O"] = "Oxygen"
  elements["F"] = "Fluorine"
  elements["Ne"] = "Neon"

  // Delete element
  delete(elements, "He")

  fmt.Println(elements)

  // Nested Dictionary or Maps
  elements2 := map[string]map[string]string{
    "H": map[string]string{
      "name":"Hydrogen",
      "state":"gas",
    },
    "He": map[string]string{
      "name":"Helium",
      "state":"gas",
    },
    "Li": map[string]string{
      "name":"Lithium",
      "state":"solid",
    },
    "Be": map[string]string{
      "name":"Beryllium",
      "state":"solid",
    },
    "B":  map[string]string{
      "name":"Boron",
      "state":"solid",
    },
    "C":  map[string]string{
      "name":"Carbon",
      "state":"solid",
    },
    "N":  map[string]string{
      "name":"Nitrogen",
      "state":"gas",
    },
    "O":  map[string]string{
      "name":"Oxygen",
      "state":"gas",
    },
    "F":  map[string]string{
      "name":"Fluorine",
      "state":"gas",
    },
    "Ne":  map[string]string{
      "name":"Neon",
      "state":"gas",
    },
  }

  // Advance if
  if el_Li, ok := elements2["Li"]; ok {
    fmt.Println(el_Li["name"], el_Li["state"])
  }
}
