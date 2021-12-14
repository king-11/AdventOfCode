package main

import (
	"fmt"

	"github.com/king-11/AdventOfCode/day10"
)

func main() {
  val, err := day10.Part2("./day10/data.txt")
  if err != nil {
    panic(err)
  }
  fmt.Println(val)
}
