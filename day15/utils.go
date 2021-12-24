package day15

import (
	"bufio"
	"math"
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
	return scanner, nil
}

func getMatrix(scanner *bufio.Scanner) ([][]int, error) {
	lines := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		vals := strings.Split(line, "")
		ints := make([]int, len(vals))
		for i, v := range vals {
			vi, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			ints[i] = vi
		}
		lines = append(lines, ints)
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return lines, nil
}

const INF = math.MaxInt64

type node struct {
	coord pair
	val   int
	index int
}

// An priorityQueue is a min-heap of ints.
type priorityQueue []*node

func (h priorityQueue) Len() int           { return len(h) }
func (h priorityQueue) Less(i, j int) bool { return h[i].val < h[j].val }
func (h priorityQueue) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h *priorityQueue) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*node))
}

func (h *priorityQueue) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	old[n-1] = nil
	x.index = -1
	*h = old[0 : n-1]
	return x
}

func expandMatrix(data [][]int) [][]int {
	n, m := len(data), len(data[0])
	expanded := make([][]int, n*5)
	for i := range expanded {
		expanded[i] = make([]int, m*5)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			expanded[i][j] = data[i][j]
		}
	}

	for i := 0; i < n*5; i++ {
		for j := 0; j < m*5; j++ {
			if i < n && j < m {
				continue
			} else if i >= n && j < m {
				expanded[i][j] = expanded[i-n][j] + 1
			} else if i < n && j >= m {
				expanded[i][j] = expanded[i][j-m] + 1
			} else {
				expanded[i][j] = expanded[i-n][j] + 1
			}
			if expanded[i][j] > 9 {
				expanded[i][j] = 1
			}
		}
	}
	return expanded
}
