package day7

import "testing"

func TestPart2(t *testing.T) {
	result, err := Part2("./test.txt")
	if err != nil {
		t.Error(err)
	}

	if result != 168 {
		t.Errorf("Expected %d, got %d", 168, result)
	}
}
