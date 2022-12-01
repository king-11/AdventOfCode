package day4

func Part1(filename string) (int, error) {
	scanner, err := getScanner(filename)
	if err != nil {
		return 0, err
	}

	order, err := readOrder(scanner)
	if err != nil {
		return 0, err
	}


	boards := make([]*Board, 0)
	for scanner.Scan() {
		board, err := readBoard(scanner)
		if err != nil {
			return 0, err
		}
		boards = append(boards, board)
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	for i, idx := range order {
		for _, val := range boards {
			val.markNumber(idx)
			if i >= 4 {
				index, _ := val.checkForColOrRow()
				if index != -1 {
					return idx*val.unmarkedMultiplication(), nil
				}
			}
		}
	}

	return 0, nil
}
