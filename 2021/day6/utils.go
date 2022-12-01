package day6

import (
	"bufio"
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

func initialState(scanner *bufio.Scanner) (map[int]int, error) {
	if scanner.Scan() {
		vals := scanner.Text()
		state := make(map[int]int)
		for _, v := range strings.Split(string(vals), ",") {
			val, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			state[val] += 1
		}
		return state, nil
	}

	if err := scanner.Err(); err != nil{
		return nil, err
	}
	return nil, scanner.Err()
}

func nextState(state map[int]int) map[int]int {
	new_state := make(map[int]int)
	for key, val := range state {
		if key == 0 {
			new_state[6] += val
			new_state[8] += val
		} else {
			new_state[key-1] += val
		}
	}

	return new_state
}
