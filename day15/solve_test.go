package day15

import "testing"

func TestPart1(t *testing.T) {
	val, err := Solve("test.txt", false)
	if err != nil {
		t.Error(err)
	}

	if val != 40 {
		t.Errorf("Expected 40, got %d", val)
	}
}

func TestExpansion(t *testing.T) {
	scan, err := getScanner("test.txt")
	if err != nil {
		t.Error(err)
	}
	mat, err := getMatrix(scan)
	if err != nil {
		t.Error(err)
	}

	exp := expandMatrix(mat)
	scan, err = getScanner("exp.txt")
	if err != nil {
		t.Error(err)
	}
	mat, err = getMatrix(scan)
	if err != nil {
		t.Error(err)
	}
	for i := range exp {
		for j := range exp[i] {
			if exp[i][j] != mat[i][j] {
				t.Errorf("Expected %d, got %d at (%d, %d)", mat[i][j], exp[i][j], i, j)
			}
		}
	}
}

func TestPart2(t *testing.T) {
	val, err := Solve("test.txt", true)
	if err != nil {
		t.Error(err)
	}

	if val != 315 {
		t.Errorf("Expected 315, got %d", val)
	}
}
