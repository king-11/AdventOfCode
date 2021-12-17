package day14

import (
	"math"
)

func Part(filename string, steps int) (int, error) {
	scanner, err := getScanner(filename)
	if err != nil {
		return 0, err
	}

	scanner.Scan()
	template := scanner.Text()
	scanner.Scan()

	rules, err := getInsertionPair(scanner)
	if err != nil {
		return 0, err
	}

	trackPairs := make(map[string]int)
	for i := 0; i < len(template)-1; i++ {
		trackPairs[string(template[i])+string(template[i+1])]++
	}

	for i := 0; i < steps; i++ {
		update := make(map[string]int)
		for k, v := range trackPairs {
			if val, ok := rules[k]; ok {
				update[string(k[0])+val] += v
				update[val+string(k[1])] += v
			}
		}
		trackPairs = update
	}

	counts := make(map[string]int)
	for k, v := range trackPairs {
		counts[string(k[0])] += v
	}
	counts[string(template[len(template)-1])]++

	max, min := 0, math.MaxInt

	for _, v := range counts {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	return max - min, nil
}
