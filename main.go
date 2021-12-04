package main

import (
	"fmt"

	"github.com/king-11/AdventOfCode/day3"
)

func main() {
  val, err := day3.Part2("./day3/data.txt")
  if err != nil {
    panic(err)
  }
  fmt.Println(val)
}
