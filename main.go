package main

import (
	"fmt"
	"time"

	"github.com/king-11/AdventOfCode/day17"
)

func main() {
	then := time.Now()
	val, err := day17.Part2("./day17/data.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("Time Taken: ", time.Since(then))
	fmt.Println("Answer: ", val)
}
