package day11

import (
	"bufio"
	"os"
)

func getScanner(filename string) (*bufio.Scanner, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	return scanner, nil
}

func processLines(scanner *bufio.Scanner) ([][]int, error) {
	lines := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		runes := make([]int, 0)
		for _, v := range line {
			val := int(v - '0')
			runes = append(runes, val)
		}
		lines = append(lines, runes)
	}
	return lines, scanner.Err()
}

type point struct {
	x int
	y int
}

type set map[point]bool

func(s set) insertIntoSet(x, y int) {
	s[point{x, y}] = true
}

func(s set) checkInSet(x, y int) bool {
	_, ok := s[point{x, y}]
	return ok
}
