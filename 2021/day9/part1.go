package day9

func Part1(filename string) (int, error) {
	scanner, err := getScaner(filename)
	if err != nil {
		return 0, err
	}

	vals, err := processLines(scanner)
	if err != nil {
		return 0, err
	}

	count := 0
	for i, arr := range vals {
		for j, val := range arr {
			check := true
			if i > 0 && vals[i-1][j] <= val {
				check = false
			}
			if i < len(vals)-1 && vals[i+1][j] <= val {
				check = false
			}
			if j > 0 && vals[i][j-1] <= val {
				check = false
			}
			if j < len(vals[i])-1 && vals[i][j+1] <= val {
				check = false
			}
			if check {
				count += (val + 1)
			}
		}
	}

	return count, nil
}
