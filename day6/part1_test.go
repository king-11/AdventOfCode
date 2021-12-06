package day6

import "testing"

func TestPart1(t *testing.T) {
	result, err := Part1("./test.txt")
	if err != nil {
		t.Error(err)
	}

	if result != 5934 {
		t.Errorf("Expected %d, got %d", 5934, result)
	}
}
