package day8

import "testing"

func TestPart2(t *testing.T) {
	val, err := Part2("test.txt")
	if err != nil {
		t.Error(err)
	}

	if val != 61229 {
		t.Errorf("Expected %d, got %d", 61229, val)
	}
}
