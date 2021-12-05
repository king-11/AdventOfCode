package day5

import "testing"

func TestPart2(t *testing.T) {
	val, err := Part2("test.txt")
	if err != nil {
		t.Error(err)
	}

	if val != 12 {
		t.Errorf("Expected 12, got %d", val)
	}
}
