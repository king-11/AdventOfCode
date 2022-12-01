package day11

func Part1(filename string) (int, error) {
	scanner, err := getScanner(filename)
	if err != nil {
		return 0, err
	}
	data, err := processLines(scanner)
	if err != nil {
		return 0, err
	}

	n, m := len(data), len(data[0])

	count := 0
	for k := 0; k < 100; k++ {
		queue := make([]point, 0)
		s := make(set)
		for i, val := range data {
			for j := range val {
				val[j] += 1
				if val[j] > 9 {
					queue = append(queue, point{i, j})
					s.insertIntoSet(i, j)
				}
			}
		}

		for len(queue) > 0 {
			p := queue[0]
			queue = queue[1:]
			count++
			data[p.x][p.y] = 0
			for _, i := range []int{-1, 0, 1} {
				for _, j := range []int{-1, 0, 1} {
					if i == 0 && j == 0 {
						continue
					}

					x, y := p.x + i, p.y + j
					if x >= 0 && x < n && y >= 0 && y < m {
						if data[x][y] != 0 {
							data[x][y] += 1
						}
						if data[x][y] > 9 && !s.checkInSet(x, y) {
							queue = append(queue, point{x, y})
							s.insertIntoSet(x, y)
						}
					}
				}
			}
		}
	}


	return count, nil
}
