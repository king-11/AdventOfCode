package day9

import (
	"bufio"
	"os"
)

func getScaner(filename string) (*bufio.Scanner, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	return scanner, nil
}

func splitToInt(s string) []int {
	var result []int
	for _, v := range s {
		result = append(result, int(v)-int('0'))
	}
	return result
}

func processLines(scanner *bufio.Scanner) ([][]int, error) {
	var result [][]int
	for scanner.Scan() {
		line := scanner.Text()
		currentResult := splitToInt(line)
		result = append(result, currentResult)
	}
	return result, scanner.Err()
}
