package day10

import "testing"

func TestPart1(t *testing.T) {
	val, err := Part1("test.txt")
	if err != nil {
		t.Error(err)
	}

	if val != 26397 {
		t.Errorf("Expected 26397, got %d", val)
	}
}
