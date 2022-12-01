package day12

import "testing"

func TestPart1(t *testing.T) {
	tests := []struct{
		f string
		res int
	}{
		{
			"test1.txt",
			10,
		},
		{
			"test2.txt",
			19,
		},
		{
			"test3.txt",
			226,
		},
	}

	for _, val := range tests {
		t.Run(val.f , func(t *testing.T) {
			res, err := Part(val.f, -1)
			if err != nil {
				t.Error(err)
			}
			if res != val.res {
				t.Errorf("expected %d, got %d", val.res, res)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct{
		f string
		res int
	}{
		{
			"test1.txt",
			36,
		},
	}

	for _, val := range tests {
		t.Run(val.f , func(t *testing.T) {
			res, err := Part(val.f, 1)
			if err != nil {
				t.Error(err)
			}
			if res != val.res {
				t.Errorf("expected %d, got %d", val.res, res)
			}
		})
	}
}
