package day17

import (
	"bufio"
	"math"
	"os"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func triangularNumber(n int) int {
	i := abs(n)
	return (i * (i + 1)) / 2
}

func inverseTriangular(n int) int {
	return int((math.Sqrt(float64(8*n+1)) - 1.0) / 2.0)
}

func getScanner(filename string) (*bufio.Scanner, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	return scanner, nil
}

func createAppend(filename, val string) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.WriteString(val); err != nil {
		return err
	}
	return nil
}

func deleteFile(filename string) error {
	return os.Remove(filename)
}
