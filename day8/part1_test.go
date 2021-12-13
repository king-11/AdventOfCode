package day8

import "testing"

func TestPart1(t *testing.T) {
	val, err := Part1("test.txt")
	if err != nil {
		t.Error(err)
	}

	if val != 26 {
		t.Errorf("Expected %d, got %d", 26, val)
	}
}
