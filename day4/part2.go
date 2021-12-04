package day4

func Part2(filename string) (int, error) {
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

	cur := 0
	checked := make(map[int]bool)
	for i, idx := range order {
		for j, val := range boards {
			_, contains := checked[j]
			if contains {
				continue
			}
			val.markNumber(idx)
			if i >= 4 {
				index, _ := val.checkForColOrRow()
				if index != -1 {
					cur += 1
					if cur == len(boards) {
						return idx*val.unmarkedMultiplication(), nil
					}
					checked[j] = true
				}
			}
		}
	}

	return 0, nil
}
