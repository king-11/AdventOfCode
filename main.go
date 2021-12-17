package main

import (
	"fmt"

	"github.com/king-11/AdventOfCode/day15"
)

func main() {
  val, err := day15.Solve("./day15/data.txt", true)
  if err != nil {
    panic(err)
  }
  fmt.Println(val)
}
