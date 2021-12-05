package day5

import "testing"

func TestPart1(t *testing.T) {
	val, err := Part1("test.txt")
	if err != nil {
		t.Error(err)
	}

	if val != 5 {
		t.Errorf("Expected 5, got %d", val)
	}
}
