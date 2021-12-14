package day10

import (
	"bufio"
	"errors"
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

func processLines(scanner *bufio.Scanner) ([][]rune, error) {
	lines := make([][]rune, 0)
	for scanner.Scan() {
		line := scanner.Text()
		runes := make([]rune, 0)
		for _, val := range line {
			runes = append(runes, val)
		}
		lines = append(lines, runes)
	}
	return lines, scanner.Err()
}

type Stack struct {
	s []rune
}

func NewStack() *Stack {
	return &Stack{make([]rune, 0)}
}

func (s *Stack) Push(v rune) {
	s.s = append(s.s, v)
}

func (s *Stack) Pop() (rune, error) {
	if s.Empty() {
		return 0, errors.New("empty Stack")
	}

	res := s.s[s.Size()-1]
	s.s = s.s[:s.Size()-1]
	return res, nil
}

func (s *Stack) Peek() (rune, error) {
	if s.Empty() {
		return 0, errors.New("empty Stack")
	}

	return s.s[s.Size()-1], nil
}

func (s *Stack) Size() int {
	return len(s.s)
}

func (s *Stack) Empty() bool {
	return len(s.s) == 0
}


func (s *Stack) isMatch(r rune) bool {
	top, err := s.Peek()
	if err != nil {
		return false
	}
	if r == ')' && top == '(' ||
		r == ']' && top == '[' ||
		r == '}' && top == '{' ||
		r == '>' && top == '<' {
		return true
	}

	return false
}

func costOfRunePart1(r rune) int {
	switch r {
	case ')':
		return 3
	case ']':
		return 57
	case '}':
		return 1197
	case '>':
		return 25137
	default:
	}
	return 0
}

func costOfRunePart2(r rune) int {
	switch r {
	case '(':
		return 1
	case '[':
		return 2
	case '{':
		return 3
	case '<':
		return 4
	default:
	}
	return 0
}
