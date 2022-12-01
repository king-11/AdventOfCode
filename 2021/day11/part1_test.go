package day11

import "testing"

func TestPart1(t *testing.T) {
	val, err := Part1("test.txt")
	if err != nil {
		t.Error(err)
	}

	if val != 1656 {
		t.Errorf("Expected 1656, got %d", val)
	}
}
