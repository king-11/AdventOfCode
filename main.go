package main

import (
	"fmt"

	"github.com/king-11/AdventOfCode/day7"
)

func main() {
  val, err := day7.Part2("./day7/data.txt")
  if err != nil {
    panic(err)
  }
  fmt.Println(val)
}
