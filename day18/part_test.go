package day18

import "testing"

func TestPart1(t *testing.T) {
	tests := []struct {
		filename string
		res int
	}{
		{"test1.txt", 143},
		{"test2.txt", 3488},
		{"test3.txt", 445},
		{"test4.txt", 791},
		{"test5.txt", 1137},
		{"test6.txt", 1384},
	}

	for _, test := range tests {
		t.Run(test.filename, func(t *testing.T) {
			res, err := Part1(test.filename)
			if err != nil {
				t.Error(err)
			}
			if res != test.res {
				t.Errorf("got %v, want %v", res, test.res)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		filename string
		res int
	}{
		{"test7.txt", 3993},
	}

	for _, test := range tests {
		t.Run(test.filename, func(t *testing.T) {
			res, err := Part2(test.filename)
			if err != nil {
				t.Error(err)
			}
			if res != test.res {
				t.Errorf("got %v, want %v", res, test.res)
			}
		})
	}
}
