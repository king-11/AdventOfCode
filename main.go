package main

import (
	"fmt"

	"github.com/king-11/AdventOfCode/day5"
)

func main() {
  val, err := day5.Part2("./day5/data.txt")
  if err != nil {
    panic(err)
  }
  fmt.Println(val)
}
