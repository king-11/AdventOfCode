package day6

func Part2(filename string) (int, error) {
	scanner, err := getScanner(filename)
	if err != nil {
		return 0, err
	}

	init, err := initialState(scanner)
	if err != nil {
		return 0, err
	}

	for i := 0; i < 256; i++ {
		init = nextState(init)
	}
	count := 0
	for _, val := range init {
		count += val
	}

	return count, nil
}
