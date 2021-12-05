package day5

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x int
	y int
}

type Line struct {
	start Point
	end   Point
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

func processLines(scanner *bufio.Scanner) ([]*Line, error) {
	lines := make([]*Line, 0)
	for scanner.Scan() {
		line := new(Line)
		fmt.Sscanf(scanner.Text(), "%d,%d -> %d,%d", &line.start.x, &line.start.y, &line.end.x, &line.end.y)
		lines = append(lines, line)
	}

	return lines, nil
}

func getGridLimits(lines []*Line) Point {
	maxx, maxy := 0, 0
	for _, line := range lines {
		if line.start.x > maxx {
			maxx = line.start.x
		}
		if line.end.x > maxx {
			maxx = line.end.x
		}
		if line.start.y > maxy {
			maxy = line.start.y
		}
		if line.end.y > maxy {
			maxy = line.end.y
		}
	}

	return Point{maxx, maxy}
}
