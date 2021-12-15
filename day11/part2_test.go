package day11

import "testing"

func TestPart2(t *testing.T) {
	val, err := Part2("test.txt")
	if err != nil {
		t.Error(err)
	}

	if val != 195 {
		t.Errorf("Expected 195, got %d", val)
	}
}
