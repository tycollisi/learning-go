package main

import (
  "fmt"
)

// Rather than making a copy of the variable passed in, here we make a copy of the pointer passed in.
// In other words, we can use pointers to access / change the original data in functions
func updateName(x *string) {
  *x = "wedge"
}

func main() {
  name := "tifa"

  m := &name

  fmt.Println("memory address of name is: ", m)  
  fmt.Println("value at memory address:", *m)
  fmt.Println(name)

  updateName(m)

  fmt.Println("memory address of name is: ", m)  
  fmt.Println("value at memory address:", *m)
  fmt.Println(name)
}
