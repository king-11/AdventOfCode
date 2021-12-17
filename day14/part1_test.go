package day14

import "testing"

func TestPart1(t *testing.T) {
	val, err := Part("test.txt", 10)
	if err != nil {
		t.Error(err)
	}

	if val != 1588 {
		t.Errorf("Expected 1588, got %d", val)
	}
}

func TestPart2(t *testing.T) {
	val, err := Part("test.txt", 40)
	if err != nil {
		t.Error(err)
	}

	if val != 2188189693529 {
		t.Errorf("Expected 2188189693529, got %d", val)
	}
}
