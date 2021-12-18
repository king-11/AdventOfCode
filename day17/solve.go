package day17

import (
	"errors"
	"fmt"
)

func Part1(filename string) (int, error) {
	scanner, err := getScanner(filename)
	if err != nil {
		return 0, err
	}

	if !scanner.Scan() {
		return 0, errors.New("no input")
	}

	xmin, xmax, ymin, ymax := 0, 0, 0, 0
	fmt.Sscanf(scanner.Text(), "target area: x=%d..%d, y=%d..%d", &xmin, &xmax, &ymin, &ymax)

	// poss_height := math.Abs(float64(ymin))

	// overshoot ymin if we try for for height hence ymin is max possible at y = 0
	return triangularNumber(ymin+1), nil
}

func Part2(filename string) (int, error) {
	scanner, err := getScanner(filename)
	if err != nil {
		return 0, err
	}

	if !scanner.Scan() {
		return 0, errors.New("no input")
	}

	xmin, xmax, ymin, ymax := 0, 0, 0, 0
	fmt.Sscanf(scanner.Text(), "target area: x=%d..%d, y=%d..%d", &xmin, &xmax, &ymin, &ymax)

	total := 0
	vxmin := inverseTriangular(xmin)
	for v0x := vxmin; v0x <= xmax; v0x++ {
		for v0y := ymin; v0y <= -ymin; v0y++ {
			x, y := 0, 0
			vx, vy := v0x, v0y
			for x <= xmax && y >= ymin {
				if x >= xmin && y <= ymax {
					total++
					break
				}

				x, y = (x + vx), (y + vy)
				vy -= 1
				if vx > 0 {
					vx -= 1
				}
			}
		}
	}

	return total, nil
}
