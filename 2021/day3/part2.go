package day3

func workOnStrings(data []string) (int, error) {
	n, size := len(data), len(data[0])
	oxy := make([]byte, size)

	consider := make(map[int]bool)
	for i := 0; i < n; i++ {
		consider[i] = true
	}
	for i := 0; i < size; i++ {
		ones, zeroes := 0, 0
		for j := 0; j < n; j++ {
			_, contains := consider[j]
			if i == 0 || (contains && data[j][i-1] == oxy[i-1]) {
				if data[j][i] == '1' {
					ones += 1
				} else {
					zeroes += 1
				}
			} else if contains {
				delete(consider, j)
			}
		}
		if ones >= zeroes {
			oxy[i] = '1'
		} else {
			oxy[i] = '0'
		}
	}

	co2 := make([]byte, size)
	consider = nil
	consider = make(map[int]bool)
	for i := 0; i < n; i++ {
		consider[i] = true
	}
	for i := 0; i < size; i++ {
		ones, zeroes := 0, 0
		if len(consider) == 1 {
			for key := range consider {
				co2 = []byte(data[key])
			}
			break
		}
		for j := 0; j < n; j++ {
			_, contains := consider[j]
			if i == 0 || (contains && data[j][i-1] == co2[i-1]) {
				if data[j][i] == '1' {
					ones += 1
				} else {
					zeroes += 1
				}
			} else if contains {
				delete(consider, j)
			}
		}
		if zeroes <= ones {
			co2[i] = '0'
		} else {
			co2[i] = '1'
		}
	}
	return binaryToInt(oxy) * binaryToInt(co2), nil
}

func Part2(filename string) (int, error) {
	data, err := readFile(filename)
	if err != nil {
		return 0, err
	}

	return workOnStrings(data)
}
