package main

import (
  "fmt"
) 

func main() {
  x := 0
  // while x is less than five, run this code
  for x < 5 {
    fmt.Println("value of x:", x)     
    x++
  }

  // while i is less than five, run this code
  for i := 0; i < 5; i++ {
    fmt.Println("value of i:", i)     
  }

  names := []string{"mario", "luigi", "yoshi", "peach"}
  fmt.Println("Length of names:",len(names))
  // while a is less than the length of names, run this code
  for a := 0; a < len(names); a++ {
    fmt.Println(names[a])     
  }

  for index, value := range names {
    fmt.Printf("The value at index %v is %v\n", index, value)
  }

  for _, value := range names {
    fmt.Printf("The value is %v\n", value)
  }

  for index, _ := range names {
    fmt.Printf("The index is %v\n", index)
  }
}

 

