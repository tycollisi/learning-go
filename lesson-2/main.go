package main

import "fmt"

func main() {
  age := 35
  name := "shaun"

  fmt.Print("Hello, ")
  fmt.Print("world!\n")

  fmt.Println("Hello, ")
  fmt.Println("world!")

  fmt.Println("my age is", age, "and my name is", name)

  fmt.Printf("my age is %v and my name is %v\n", age, name)
  fmt.Printf("my age is %v and my name is %q\n", age, name)
  fmt.Printf("age is of type %T\n", age)
  fmt.Printf("you scored %0.2f points!\n", 225.55)

  // Sprintf (save formatted strings)
  savedString := fmt.Sprintf("my age is %v and my name is %v\n", age, name)
  fmt.Println("The saved string is:", savedString)

}
