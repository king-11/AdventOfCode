package day7

import (
	"math"
)

func Part1(filename string) (int, error) {
	locations, err := getLocations(filename)
	if err != nil {
		return 0, err
	}
	median := getMedian(locations)
	sum := 0.0
	for _, value := range locations {
		sum += math.Abs(float64(value - median))
	}
	return int(sum), nil
}
