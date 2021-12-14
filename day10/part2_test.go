package day10

import "testing"

func TestPart2(t *testing.T) {
	val, err := Part2("test.txt")
	if err != nil {
		t.Error(err)
	}

	if val != 288957 {
		t.Errorf("Expected 288957, got %d", val)
	}
}
