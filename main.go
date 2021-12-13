package main

import (
	"fmt"

	"github.com/king-11/AdventOfCode/day8"
)

func main() {
  val, err := day8.Part2("./day8/data.txt")
  if err != nil {
    panic(err)
  }
  fmt.Println(val)
}
