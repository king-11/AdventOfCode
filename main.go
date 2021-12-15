package main

import (
	"fmt"

	"github.com/king-11/AdventOfCode/day11"
)

func main() {
  val, err := day11.Part2("./day11/data.txt")
  if err != nil {
    panic(err)
  }
  fmt.Println(val)
}
