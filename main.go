package main

import (
	"fmt"

	"github.com/king-11/AdventOfCode/day13"
)

func main() {
  val, err := day13.Part("./day13/data.txt")
  if err != nil {
    panic(err)
  }
  fmt.Println(val)
}
