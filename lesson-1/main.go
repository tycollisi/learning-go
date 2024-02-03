package main

// Go Standard Library: fmt formatting strings and printing messages to the console.
import "fmt"

func main() {
  var nameOne string = "mario"
  var nameTwo = "luigi"
  var nameThree string

  fmt.Println(nameOne, nameTwo, nameThree)

  nameOne = "peach"
  nameThree = "bowser"

  fmt.Println(nameOne, nameTwo, nameThree)

  nameFour := "yoshi"
  fmt.Println(nameFour)

  var ageOne int = 20
  var ageTwo = 30
  ageThree := 40

  fmt.Println(ageOne, ageTwo, ageThree)

  var scoreOne float32 = 25.98 
  var scoreTwo float64 = 89248353475853857.23
  scoreThree := 1.5

  fmt.Println(scoreOne, scoreTwo, scoreThree)

}
