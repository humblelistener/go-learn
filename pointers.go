package main

import (
  "fmt"
)

func main(){
  name := "kavin"

  fmt.Println("my name is", name, ". This info is stored in", &name);

  //by default, go passes parameters by value
  addMyLastName(name)

  fmt.Println("but still my name is", name)
}

func addMyLastName(name string) string {
  name = "kavin padmanabhan"
  fmt.Println("my name per that method is", name)

  return name
}
