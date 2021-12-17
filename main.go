package main

import (
	"fmt"
	"time"

	"github.com/king-11/AdventOfCode/day15"
)

func main() {
	then := time.Now()
	val, err := day15.Solve("./day15/data.txt", true)
	if err != nil {
		panic(err)
	}
	fmt.Println("Time Taken: ", time.Since(then))
	fmt.Println("Answer: ", val)
}
