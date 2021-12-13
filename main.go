package main

import (
	"fmt"

	"github.com/king-11/AdventOfCode/day9"
)

func main() {
  val, err := day9.Part2("./day9/data.txt")
  if err != nil {
    panic(err)
  }
  fmt.Println(val)
}
