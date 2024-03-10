package main

import (
	"bufio"
	"fmt"
	"os"
  "strings"
  "strconv"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
  fmt.Print(prompt)
  input, err := r.ReadString('\n')

  return strings.TrimSpace(input), err
}

func promptOptions(b bill) {
  reader := bufio.NewReader(os.Stdin) 

  opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)

  switch opt {
  case "a":
    name, _ := getInput("Enter item name: ", reader)
    price, _ := getInput("Enter item price: ", reader)

    p, err := strconv.ParseFloat(price, 64)
    if err != nil {
      fmt.Println("The price must be a number.")
      promptOptions(b)
    }

    b.addItem(name, p) 

    fmt.Printf("Item added, name: %v price: %v\n", name, price)
    promptOptions(b)
  case "s":
    b.save()
    fmt.Println("Saved the bill: ", b.name)
  case "t":
    tip, _ := getInput("Enter tip amount ($): ", reader)
    t, err := strconv.ParseFloat(tip, 64)
    if err != nil {
      fmt.Println("The tip must be a number.")
      promptOptions(b)
    }

    b.updateTip(t)
    fmt.Printf("Tip has been addded: %v\n", tip)
    promptOptions(b)
  default:
    fmt.Println("Invalid option, please try again.")
    promptOptions(b)
  }
}

// Create Bill
func createBill() bill {
  reader := bufio.NewReader(os.Stdin)

  name, _ := getInput("Create a new bill name: ", reader)

  b := newBill(name)
  fmt.Println("Created the bill -", b.name)

  return b
}

func main() {
  mybill := createBill()
  promptOptions(mybill)
  //mybill := newBill("Tyler's bill")
  //mybill.addItem("banana", 1.99)
  //mybill.addItem("coffee", 3.25)
  //mybill.updateTip(4.99)
  //fmt.Println(mybill.format())
}




