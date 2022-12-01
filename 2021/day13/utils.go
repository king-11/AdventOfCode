package day13

import (
	"bufio"
	"fmt"
	"os"
)

type Grid2D [][]rune
type Point struct {
	x int
	y int
}

type Points []*Point

func(ps Points) getMax() Point {
	p := Point{}
	for _, point := range ps {
		if point.x > p.x {
			p.x = point.x
		}
		if point.y > p.y {
			p.y = point.y
		}
	}
	return p
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

func newGrid(x, y int) Grid2D {
	grid := make(Grid2D, y + 1)
	for i := range grid {
		grid[i] = make([]rune, x + 1)
		for j := range grid[i] {
			grid[i][j] = rune('.')
		}
	}
	return grid
}

func(grid Grid2D) fillGrid(points Points) {
	for _, point := range points {
		grid[point.y][point.x] = rune('#')
	}
}

func(grid Grid2D) getPoints() []Point {
	points := make([]Point, 0)
	for y, row := range grid {
		for x, val := range row {
			if val == rune('#') {
				points = append(points, Point{x, y})
			}
		}
	}
	return points
}

func(p *Point) foldPoint(fold Point) {
	if fold.x != - 1 && p.x >= fold.x {
		p.x = 2*fold.x - p.x
	}
	if fold.y != - 1 && p.y >= fold.y {
		p.y = 2*fold.y - p.y
	}
}

func getPoints(scanner *bufio.Scanner) (Points, error) {
	points := make([]*Point, 0)
	for scanner.Scan() {
		row := scanner.Text()
		if row == "" {
			break
		}
		x, y := 0, 0
		fmt.Sscanf(row, "%d,%d", &x, &y)
		points = append(points, &Point{x, y})
	}
	return points, scanner.Err()
}
