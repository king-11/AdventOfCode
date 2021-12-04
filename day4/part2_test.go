package day4

import "testing"

func TestPart2(t *testing.T) {
	val, err := Part2("./test.txt")
	if err != nil {
		t.Error(err)
	}

	if val != 1924 {
		t.Errorf("Expected 1924, got %d", val)
	}
}
