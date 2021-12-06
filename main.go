package main

import (
	"fmt"

	"github.com/king-11/AdventOfCode/day6"
)

func main() {
  val, err := day6.Part2("./day6/data.txt")
  if err != nil {
    panic(err)
  }
  fmt.Println(val)
}
