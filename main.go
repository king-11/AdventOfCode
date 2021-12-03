package main

import (
	"fmt"

	"github.com/king-11/AdventOfCode/day1"
)

func main() {
  val, err := day1.Part2("./day2/data.txt")
  if err != nil {
    panic(err)
  }
  fmt.Println(val)
}
