package day18

import (
	"sync"

	"github.com/king-11/AdventOfCode/utils"
)

func Part1(filename string) (int, error) {
	scanner, err := utils.GetScanner(filename)
	if err != nil {
		return 0, err
	}

	for scanner.Scan() {
		if HEAD == nil {
			HEAD, err = newNode(scanner.Text())
			if err != nil {
				return 0, err
			}
		} else {
			cur_node, err := newNode(scanner.Text())
			if err != nil {
				return 0, err
			}
			HEAD = HEAD.join(cur_node)
			HEAD.explore()
		}
	}

	val := HEAD.value()
	HEAD = nil
	return val, nil
}

func Part2(filename string) (int, error) {
	scanner, err := utils.GetScanner(filename)
	if err != nil {
		return 0, err
	}

	nodes := make([]string, 0)
	for scanner.Scan() {
		node := scanner.Text()
		if err != nil {
			return 0, err
		}
		nodes = append(nodes, node)
	}

	if err = scanner.Err(); err != nil {
		return 0, err
	}

	max := new(int)
	sem := make(chan int, 8)
	errChan := make(chan error, 1)
	var wg sync.WaitGroup
	var mutex sync.Mutex
	for i := range nodes {
		for j := i + 1; j < len(nodes); j++ {
			wg.Add(1)
			go worker(nodes[i], nodes[j], sem, errChan, &wg, &mutex, max)
			wg.Add(1)
			go worker(nodes[j], nodes[i], sem, errChan, &wg, &mutex, max)
		}
	}
	wg.Wait()
	close(errChan)

	return *max, <-errChan
}

func worker(f string, s string, sem chan int, errChan chan error, wg *sync.WaitGroup, mutex *sync.Mutex, max *int) {
	defer wg.Done()
	sem <- 1
	first, err := newNode(f)
	if err != nil {
		select {
		case errChan <- err:
			// we're the first worker to fail
		default:
			// some other failure has already happened
		}
	}
	second, err := newNode(s)
	if err != nil {
		select {
		case errChan <- err:
			// we're the first worker to fail
		default:
			// some other failure has already happened
		}
	} else {
		HEAD := first.join(second)
		HEAD.explore()
		val := HEAD.value()
		mutex.Lock()
		defer mutex.Unlock()
		if val > *max {
			*max = val
		}
		HEAD = nil
		first = nil
		second = nil
	}
	<-sem
}
