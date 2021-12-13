package day8

import (
	"bufio"
	"os"
	"strings"
)

type Entry struct {
	signals []string
	digits []string
}

func getScanner(filename string) (*bufio.Scanner, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	return scanner, nil
}

func processEntries(scanner *bufio.Scanner) ([]*Entry, error) {
	entries := make([]*Entry, 0)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "|")
		entry := &Entry{
			signals: strings.Fields(strings.TrimSpace(line[0])),
			digits: strings.Fields(strings.TrimSpace(line[1])),
		}
		entries = append(entries, entry)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return entries, nil
}
