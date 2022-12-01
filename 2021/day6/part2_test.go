package day6

import "testing"

func TestPart2(t *testing.T) {
	result, err := Part2("./test.txt")
	if err != nil {
		t.Error(err)
	}

	if result != 26984457539 {
		t.Errorf("Expected %d, got %d", 5934, result)
	}
}
