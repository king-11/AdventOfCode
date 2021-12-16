package main

import (
	"fmt"

	"github.com/king-11/AdventOfCode/day12"
)

func main() {
  val, err := day12.Part("./day12/data.txt", 1)
  if err != nil {
    panic(err)
  }
  fmt.Println(val)
}
