package day7

import "math"

func calcFuel(l, r int) int {
	val := math.Abs(float64(l) - float64(r))
	return int(val*(val+1)/2)
}

func Part2(filename string) (int, error) {
	locations, err := getLocations(filename)
	if err != nil {
		return 0, err
	}
	mean := getMean(locations)
	check_on := []int{
		mean - 5,
		mean - 4,
		mean - 3,
		mean - 2,
		mean - 1,
		mean,
		mean + 1,
		mean + 2,
		mean + 3,
		mean + 4,
		mean + 5,
	}
	sum := math.MaxInt64
	for _, val := range check_on {
		localSum := 0
		for _, l := range locations {
			localSum += calcFuel(l, val)
		}

		if localSum < sum {
			sum = localSum
		}
	}
	return sum, nil
}
