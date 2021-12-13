package day7

import (
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getData(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(strings.Trim(string(data), "\n")), nil
}

func getLocations(filename string) ([]int, error) {
	var locations []int
	data, err := getData(filename)
	if err != nil {
		return nil, err
	}
	for _, value := range strings.Split(data, ",") {
		val, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}
		locations = append(locations, val)
	}
	sort.Ints(locations)
	return locations, nil
}

func getMedian(values []int) int {
	length := len(values)
	if length%2 == 0 {
		return values[length/2-1]
	}
	return values[length/2]
}

func getMean(values []int) int {
	var total int
	for _, value := range values {
		total += value
	}
	return total / len(values)
}
