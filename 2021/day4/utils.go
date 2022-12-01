package day4

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
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

func convertToInt(s string) (int, error) {
	val, err := strconv.Atoi(s)
	return val, err
}

func readOrder(scanner *bufio.Scanner) ([]int, error) {
	var err error
	if scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			return nil, errors.New("empty line")
		}

		val := strings.Split(line, ",")
		order := make([]int, len(val))
		for i, v := range val {
			order[i], err = convertToInt(v)
			if err != nil {
				return nil, err
			}
		}
		return order, nil
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return nil, nil
}

type Board [5][5]Data

type Data struct {
	Val int
	Marked bool
}

func readBoard(scanner *bufio.Scanner) (*Board, error) {
	board := new(Board)
	var err error
	for i := 0; i < 5 && scanner.Scan(); i++ {
		line := scanner.Text()
		if line == "" {
			return &Board{}, errors.New("empty line")
		}

		val := strings.Fields(line)
		for j, v := range val {
			board[i][j].Val, err = convertToInt(v)
			if err != nil {
				return &Board{}, err
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return &Board{}, err
	}

	return board, nil
}

func(board *Board) checkForColOrRow() (idx int, isRow bool) {
	// row
	for i := 0; i < 5; i++ {
		marked := true
		// column
		for j := 0; j < 5; j++ {
			if !board[i][j].Marked {
				marked = false
				break
			}
		}
		if marked {
			return i, true
		}
	}

	// column
	for i := 0; i < 5; i++ {
		marked := true
		// row
		for j := 0; j < 5; j++ {
			if !board[j][i].Marked {
				marked = false
				break
			}
		}
		if marked {
			return i, false
		}
	}

	return -1, false
}

func(board *Board) markNumber(val int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if board[i][j].Val == val {
				board[i][j].Marked = true
				return
			}
		}
	}
}

func(board *Board) unmarkedMultiplication() int {
	mult := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !board[i][j].Marked {
				mult += board[i][j].Val
			}
		}
	}

	return mult
}
