package day13

import "testing"

func TestPart1(t *testing.T) {
	res, err := Part("test.txt")
	if err != nil {
		t.Error(err)
	}
	if res != 16 {
		t.Errorf("expected %d, got %d", 16, res)
	}
}
