package main

import (
	"fmt"

	"github.com/king-11/AdventOfCode/day2"
)

func main() {
  val, err := day2.Part2("./day2/data.txt")
  if err != nil {
    panic(err)
  }
  fmt.Println(val)
}
