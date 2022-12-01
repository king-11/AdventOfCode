package day9

import "testing"

func TestPart1(t *testing.T) {
	res, err := Part1("test.txt")

	if err != nil {
		t.Error(err)
	}

	if res != 15 {
		t.Errorf("Expected 15, got %d", res)
	}
}
