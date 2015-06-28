package main

import (
  "fmt"
  "reflect"
)

// this is global variable declaration
var (
  name, favGame string
  age int
  schoolName = "McKinnon kinder garden"
)

func main() {
  name = "kavin"
  age = 4

  //now this local variable declaration
  bestFriendsName := "frankie"

  fmt.Println("My name is", name, "and my age is", age)
  fmt.Println("type of name is", reflect.TypeOf(name))
  fmt.Println("type of age is", reflect.TypeOf(age))

  fmt.Println(name, "'s best friends name is", bestFriendsName)
}
