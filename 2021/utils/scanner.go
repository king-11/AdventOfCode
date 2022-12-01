package utils

import (
	"bufio"
	"os"
)

func GetScanner(filename string) (*bufio.Scanner, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	return bufio.NewScanner(f), nil
}
