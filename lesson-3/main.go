package main

import "fmt"

func main() {
  // arrays have fixed lengths in go
  //var ages [3]int = [3]int{20, 25, 30}
  var ages = [3]int{20, 25, 30}

  names := [4]string{"john", "monica", "charlotte", "rodrigo"}
  names[0] = "tyler"
  fmt.Println(ages, len(ages))
  fmt.Println(names, len(names))

  // slices (use arrays under the hood)
  var scores = []int{100, 50, 60}
  fmt.Println(scores, len(scores))
  scores = append(scores, 85)
  fmt.Println(scores, len(scores))

  // slice ranges
  rangeOne := names[1:3]
  rangeTwo := names[2:]
  rangeThree := names[:3]
  fmt.Println(rangeOne)
  fmt.Println(rangeTwo)
  fmt.Println(rangeThree)

}
