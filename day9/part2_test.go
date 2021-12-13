package day9

import "testing"

func TestPart2(t *testing.T) {
	val, err := Part2("test.txt")
	if err != nil {
		t.Error(err)
	}

	if val != 1134 {
		t.Errorf("Expected 1134, got %d", val)
	}
}
