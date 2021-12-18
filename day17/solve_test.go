package day17

import (
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		text string
		result int
	}{
		{"target area: x=20..30, y=-10..-5", 45},
	}

	for i, val := range tests {
		testName :=  "test"+fmt.Sprintf("%d",i)+".txt"
		t.Run(testName , func(t *testing.T) {
			err := createAppend(testName, val.text)
			defer deleteFile(testName)
			if err != nil {
				t.Error(err)
			}
			res, err := Part1(testName)
			if err != nil {
				t.Error(err)
			}
			if res != val.result {
				t.Errorf("expected %d, got %d", val.result, res)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		text string
		result int
	}{
		{"target area: x=20..30, y=-10..-5", 112},
	}

	for i, val := range tests {
		testName :=  "test"+fmt.Sprintf("%d",i)+".txt"
		t.Run(testName , func(t *testing.T) {
			err := createAppend(testName, val.text)
			defer deleteFile(testName)
			if err != nil {
				t.Error(err)
			}
			res, err := Part2(testName)
			if err != nil {
				t.Error(err)
			}
			if res != val.result {
				t.Errorf("expected %d, got %d", val.result, res)
			}
		})
	}
}
