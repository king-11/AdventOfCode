package main

import (
	"fmt"
	"time"

	"github.com/king-11/AdventOfCode/day18"
)

func main() {
	then := time.Now()
	val, err := day18.Part2("./day18/data.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("Time Taken: ", time.Since(then))
	fmt.Println("Answer: ", val)
}
