package day3

import "testing"

func TestPart2(t *testing.T) {
	data := []string {
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	val, err := workOnStrings(data)
	if err != nil {
		t.Error(err)
	}

	if val != 230 {
		t.Errorf("Expected 230, got %d", val)
	}
}
