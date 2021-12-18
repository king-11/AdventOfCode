package main

import (
	"fmt"
	"time"

	"github.com/king-11/AdventOfCode/day16"
)

func main() {
	then := time.Now()
	val, err := day16.Part2("./day16/data.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("Time Taken: ", time.Since(then))
	fmt.Println("Answer: ", val)
}
