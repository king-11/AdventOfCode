package main

import (
	"fmt"

	"github.com/king-11/AdventOfCode/day4"
)

func main() {
  val, err := day4.Part2("./day4/data.txt")
  if err != nil {
    panic(err)
  }
  fmt.Println(val)
}
