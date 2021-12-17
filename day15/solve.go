package day15

import (
	"container/heap"
)

type pair struct {
	x, y int
}

func dijksta(graph [][]int) int {
	n, m := len(graph), len(graph[0])
	parents := make([][]pair, n)
	distances := make([][]int, n)
	for i := 0; i < n; i++ {
		parents[i] = make([]pair, m)
		distances[i] = make([]int, m)
		for j := range distances[i] {
			distances[i][j] = INF
			parents[i][j] = pair{-1, -1}
		}
	}
	distances[0][0] = 0
	visited := make(map[pair]bool)
	pq := &priorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &node{coord:pair{0, 0}, val:0})
	for pq.Len() > 0 {
		cur := heap.Pop(pq).(*node)
		if distances[cur.coord.x][cur.coord.y] < cur.val {
			continue
		}

		for _, val := range []pair{{1, 0}, {0, 1}, {0, -1}, {-1, 0}} {
			x := cur.coord.x + val.x
			y := cur.coord.y + val.y
			if x < 0 || x >= n || y < 0 || y >= m {
				continue
			}
			if visited[pair{x, y}] {
				continue
			}
			if distances[x][y] > distances[cur.coord.x][cur.coord.y]+graph[x][y] {
				distances[x][y] = distances[cur.coord.x][cur.coord.y] + graph[x][y]
				parents[x][y] = pair{cur.coord.x, cur.coord.y}
				heap.Push(pq, &node{coord: pair{x, y}, val: distances[x][y]})
			}
		}

		visited[cur.coord] = true
	}

	return distances[n-1][m-1]
}

func Solve(filename string, part2 bool) (int, error) {
	scanner, err := getScanner(filename)
	if err != nil {
		return 0, err
	}

	data, err := getMatrix(scanner)
	if part2 {
		data = expandMatrix(data)
	}
	if err != nil {
		return 0, err
	}
	val := dijksta(data)
	return val, nil
}
