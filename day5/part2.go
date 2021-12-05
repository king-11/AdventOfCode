package day5

import "math"

func Part2(filename string) (int, error) {
	scanner, err := getScanner(filename)
	if err != nil {
		return 0, err
	}

	lines, err := processLines(scanner)
	if err != nil {
		return 0, err
	}

	limit := getGridLimits(lines)

	grid := make([][]int, limit.y+1)
	for i := range grid {
		grid[i] = make([]int, limit.x+1)
	}

	for _, line := range lines {
		stepx := 0
		if line.start.x != line.end.x {
			stepx = int(math.Abs(float64(line.end.x)-float64(line.start.x)) / (float64(line.end.x) - float64(line.start.x)))
		}
		stepy := 0
		if line.start.y != line.end.y {
			stepy = int(math.Abs(float64(line.end.y)-float64(line.start.y)) / (float64(line.end.y) - float64(line.start.y)))
		}

		if stepx == 0 {
			x := line.start.x
			for y := line.start.y; y != line.end.y; y += stepy {
				grid[y][x]++
			}
			grid[line.end.y][x]++
		} else if stepy == 0 {
			y := line.start.y
			for x := line.start.x; x != line.end.x; x += stepx {
				grid[y][x]++
			}
			grid[y][line.end.x]++
		} else {
			x := line.start.x
			y := line.start.y
			for {
				grid[y][x]++
				if x == line.end.x && y == line.end.y {
					break
				}
				x += stepx
				y += stepy
			}
		}

	}

	count := 0
	for _, val := range grid {
		for _, v := range val {
			if v > 1 {
				count += 1
			}
		}
	}

	return count, nil
}
