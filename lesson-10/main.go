package main

import (
  "fmt"
)

func updateName(x string) string {
  x = "wedge"
  return x
}

func updateMenu(y map[string]float64) {
  y["coffee"] = 2.99
}

func main() {
  // Non-pointer value var types: strings, ints, bools, floats, arrays, structs
  name := "tifa"

  // Let's say we did want to update the original 'name' variable, we could do
  // name = updateName(name)
  updateName(name) 

  // Every time go passes a variable into a function, Go creates a copy of that variable
  // In other words, our 'updateName' function changed the copied name variable, not the original
  fmt.Println(name)

  // Pointer wrapper value var types: slices, maps, functions
  // A copy of these var types is still made, but the copy points at the same memory block. See example image in memory_example_image directory.
  menu := map[string]float64{
    "pie": 5.95,
    "ice cream": 3.99,
  }

  updateMenu(menu)
  fmt.Println(menu)
}

