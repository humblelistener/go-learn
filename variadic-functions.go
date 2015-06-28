package main

import (
  "fmt"
)

func main(){

    highScore := highestScore(334, 324, 345, 424, 234, 123, 352)

    fmt.Println(highScore)
}

func highestScore(scores ...int) int {
  high := scores[0]
  for _, i := range scores {
    if i > high {
      high = i
    }
  }

  return high
}
