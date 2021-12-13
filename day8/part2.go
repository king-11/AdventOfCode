package day8

import (
	"sort"
	"strings"
)

func findCommon(s, t string) int {
	count := 0
	for _, v := range s {
		if strings.ContainsRune(t, v) {
			count++
		}
	}
	return count
}

func sortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })
	return string(r)
}

func Part2(filename string) (int, error) {
	scanner, err := getScanner(filename)
	if err != nil {
		return 0, err
	}

	entries, err := processEntries(scanner)
	if err != nil {
		return 0, err
	}

	count := 0
	for _, val := range entries {
		mapper := make(map[int]string)
		for _, v := range val.signals {
			if i := len(v); i == 2 {
				mapper[1] = v
			} else if i == 3 {
				mapper[7] = v
			} else if i == 4 {
				mapper[4] = v
			} else if i == 7 {
				mapper[8] = v
			}
		}
		// fmt.Print(mapper, " | ")

		for _, v := range val.signals {
			if i := len(v); i == 6 {
				if findCommon(mapper[7], v) == 2 {
					mapper[6] = v
				} else {
					if findCommon(mapper[4], v) == 3 {
						mapper[0] = v
					} else {
						mapper[9] = v
					}
				}
			} else if i == 5 {
				if findCommon(mapper[1], v) == 2 {
					mapper[3] = v
				}
			}
		}
		// fmt.Println(mapper, " | ")

		for _, v := range val.signals {
			if i := len(v); i == 5 {
				if findCommon(mapper[6], v) == 5 {
					mapper[5] = v
				} else if v != mapper[3] {
					mapper[2] = v
				}
			}
		}

		mapped := make([]int, 0)
		for _, v := range val.digits {
			for key, val := range mapper {
				if sortString(v) == sortString(val) {
					mapped = append(mapped, key)
				}
			}
		}

		val := 0
		for _, v := range mapped {
			val += v
			val *= 10
		}
		val /= 10
		count += val
	}

	return count, nil
}
