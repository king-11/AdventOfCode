package day13

import "fmt"

func Part(filename string) (int, error) {
	scanner, err := getScanner(filename)
	if err != nil {
		return 0, err
	}
	points, err := getPoints(scanner)
	if err != nil {
		return 0, err
	}
	for scanner.Scan() {
		line := scanner.Text()
		dir := rune('.')
		val := 0
		point := Point{-1, -1}
		fmt.Sscanf(line, "fold along %c=%d", &dir, &val)
		if dir == 'x' {
			point.x = val
		} else {
			point.y = val
		}
		for _, p := range points {
			p.foldPoint(point)
		}
	}

	point := points.getMax()
	grid := newGrid(point.x, point.y)
	grid.fillGrid(points)
	for _, val := range grid {
		fmt.Printf("%s\n", string(val))
	}
	return len(grid.getPoints()), nil
}
