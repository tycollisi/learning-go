package main

import "fmt"

func main() {

  menu := map[string]float64{
    "soup": 4.99,
    "pie": 7.99,
    "salad": 6.99,
    "toffee pudding": 3.55,
  }

  fmt.Println(menu)
  fmt.Println(menu["pie"])

  // looping maps
  for k, v := range menu {
    fmt.Println(k, "-", v)
  }

  // ints as key type
  phonebook := map[int]string{
    1111111111: "tyler",
    2222222222: "jon",
    3333333333: "kyle",
  }

  fmt.Println(phonebook)
  fmt.Println(phonebook[1111111111])
  for k, v := range phonebook {
    fmt.Println(k, "-", v)
  }

  phonebook[1111111111] = "bowser"
  fmt.Println(phonebook[1111111111])

}
