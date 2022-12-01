package day10

func Part1(filename string) (int, error) {
	scanner, err := getScanner(filename)
	if err != nil {
		return 0, err
	}
	data, err := processLines(scanner)
	if err != nil {
		return 0, err
	}

	cost := 0
	for _, val := range data {
		stack := NewStack()
		for _, v := range val {
			if v == '[' || v == '(' || v == '{' || v == '<' {
				stack.Push(v)
			} else {
				if stack.isMatch(v) {
					stack.Pop()
				} else {
					cost += costOfRunePart1(v)
					break
				}
			}
		}
	}
	return cost, nil
}
