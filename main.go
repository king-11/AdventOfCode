package main

import (
	"fmt"

	"github.com/king-11/AdventOfCode/day14"
)

func main() {
  val, err := day14.Part("./day14/data.txt", 40)
  if err != nil {
    panic(err)
  }
  fmt.Println(val)
}
