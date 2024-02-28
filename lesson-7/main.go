package main

import (
  "fmt"
  "strings"
)

func getInitials(n string) (string, string) {
  s := strings.ToUpper(n)
  names := strings.Split(s, " ")

  var initials []string
  for _, v := range names {
    initials = append(initials, v[:1])
  }

  if len(initials) > 1 {
    return initials[0], initials[1]
  }

  return initials[0], "_"
}

func main() {
  firstInitial, secondInitial := getInitials("tyler collisi")
  fmt.Printf("firstInitial = %v secondInitial = %v\n", firstInitial, secondInitial)

  i1, i2 := getInitials("jon")
  fmt.Printf("i1 = %v i2 = %v\n", i1, i2)
}
