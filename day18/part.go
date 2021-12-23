package day18

import (
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

	max := 0
	for i := range nodes {
		for  j := i+1; j < len(nodes); j++ {
			cur_node, err := newNode(nodes[i])
			if err != nil {
				return 0, err
			}
			temp_node, err := newNode(nodes[j])
			if err != nil {
				return 0, err
			}
			HEAD = cur_node.join(temp_node)
			HEAD.explore()
			if val := HEAD.value(); val > max {
				max = val
			}
			HEAD = nil
			cur_node = nil
			temp_node = nil
			cur_node, _ = newNode(nodes[i])
			temp_node, _ = newNode(nodes[j])
			HEAD = temp_node.join(cur_node)
			HEAD.explore()
			if val := HEAD.value(); val > max {
				max = val
			}
			HEAD = nil
			cur_node = nil
			temp_node = nil
		}
	}

	return max, nil
}
