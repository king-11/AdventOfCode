package day8

func Part1(filename string) (int, error) {
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
		for _, v := range val.digits {
			if i := len(v); i == 2 || i == 3 || i == 7 || i == 4 {
				count += 1
			}
		}
	}

	return count, nil
}
