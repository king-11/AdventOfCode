package day12

func Part(filename string, allowDouble int) (int, error) {
	scanner, err := getScanner(filename)
	if err != nil {
		return 0, err
	}
	data, err := processLines(scanner)
	if err != nil {
		return 0, err
	}

	neighbours := convertLinesToNeighbours(data)
	cache := make(map[int]int)
	var countPaths func(int, int, int) int
	countPaths = func(cave int, seen int, allowDoubleVisit int) int {
		if cave == END_CAVE_ID {
			return 1
		}

		if cave >= MIN_SMALL_CAVE_ID {
			if seen%cave == 0 {
				if allowDoubleVisit == -1 {
					return 0
				}
				allowDoubleVisit = -1
			} else {
				seen *= cave
			}
		}

		total := 0
		for _, neighbour := range neighbours[cave] {
			cacheKey := (neighbour + 1) * seen * allowDoubleVisit
			count, ok := cache[cacheKey]
			if !ok {
				count = countPaths(neighbour, seen, allowDoubleVisit)
				cache[cacheKey] = count
			}
			total += count
		}
		return total
	}

	cave := START_CAVE_ID
	seen := cave
	count := countPaths(cave, seen, allowDouble)

	return count, nil
}
