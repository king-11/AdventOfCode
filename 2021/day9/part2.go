package day9

import (
	"sort"
	"sync"
)

func Part2(filename string) (int, error) {
	scanner, err := getScaner(filename)
	if err != nil {
		return 0, err
	}

	vals, err := processLines(scanner)
	if err != nil {
		return 0, err
	}

	marked := make([][]bool, len(vals))
	for i := range marked {
		marked[i] = make([]bool, len(vals[i]))
	}

	var wg sync.WaitGroup
	var func1 func(i , j int, sendto chan int) int
	func1 = func(i, j int, sendto chan int) int {
		marked[i][j] = true
		bsize := 1
		if i > 0 && !marked[i-1][j] && vals[i-1][j] > vals[i][j] && vals[i-1][j] != 9 {
			bsize += func1(i-1, j, nil)
		}
		if i < len(vals)-1 && !marked[i+1][j] && vals[i+1][j] > vals[i][j] && vals[i+1][j] != 9 {
			bsize += func1(i+1, j, nil)
		}
		if j > 0 && !marked[i][j-1] && vals[i][j-1] > vals[i][j] && vals[i][j-1] != 9 {
			bsize += func1(i, j-1, nil)
		}
		if j < len(vals[i])-1 && !marked[i][j+1] && vals[i][j+1] > vals[i][j] && vals[i][j+1] != 9 {
			bsize += func1(i, j+1, nil)
		}

		if sendto != nil {
			sendto <- bsize
			wg.Done()
		}

		return bsize
	}

	sizes := make(chan int, len(vals)*len(vals[0]))
	for i, arr := range vals {
		for j, val := range arr {
			check := true
			if i > 0 && vals[i-1][j] <= val {
				check = false
			}
			if i < len(vals)-1 && vals[i+1][j] <= val {
				check = false
			}
			if j > 0 && vals[i][j-1] <= val {
				check = false
			}
			if j < len(vals[i])-1 && vals[i][j+1] <= val {
				check = false
			}
			if check {
				wg.Add(1)
				go func1(i, j, sizes)
			}
		}
	}
	wg.Wait()
	close(sizes)
	values := make([]int, 0)
	for v := range sizes {
		values = append(values, v)
	}
	sort.Ints(values)
	n := len(values)
	return values[n-1]*values[n-2]*values[n-3], nil
}
