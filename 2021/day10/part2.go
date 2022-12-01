package day10

import (
	"sort"
)

func Part2(filename string) (int, error) {
	scanner, err := getScanner(filename)
	if err != nil {
		return 0, err
	}
	data, err := processLines(scanner)
	if err != nil {
		return 0, err
	}

	costs := make([]int, 0)
	for _, val := range data {
		stack := NewStack()
		cost := 0
		for _, v := range val {
			if v == '[' || v == '(' || v == '{' || v == '<' {
				stack.Push(v)
			} else {
				if stack.isMatch(v) {
					stack.Pop()
				} else {
					stack = nil
					break
				}
			}
		}

		if stack != nil {
			for !stack.Empty() {
				val, _ := stack.Pop()
				cost *= 5
				cost += costOfRunePart2(val)
			}
			costs = append(costs, cost)
		}
	}
	n := len(costs)
	sort.Ints(costs)
	return costs[n/2], nil
}
