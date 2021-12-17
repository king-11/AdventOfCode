package day14

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getScanner(filename string) (*bufio.Scanner, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	return scanner, nil
}

func getInsertionPair(scanner *bufio.Scanner) (map[string]string, error) {
	lines := make(map[string]string)
	for scanner.Scan() {
		line := scanner.Text()
		var key string
		var val rune
		_, err := fmt.Sscanf(line, "%s -> %c", &key, &val)
		if err != nil {
			return nil, err
		}
		lines[key] = strings.TrimSpace(string(val))
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return lines, nil
}
