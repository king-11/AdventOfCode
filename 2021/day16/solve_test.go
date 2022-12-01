package day16

import (
	"fmt"
	"os"
	"testing"
)

func createAppend(filename, val string) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.WriteString(val); err != nil {
		return err
	}
	return nil
}

func deleteFile(filename string) error {
	return os.Remove(filename)
}

func TestPart1(t *testing.T) {
	tests := []struct {
		text string
		result int
	}{
		{"D2FE28", 6},
		{"8A004A801A8002F478", 16},
		{"620080001611562C8802118E34", 12},
		{"C0015000016115A2E0802F182340", 23},
		{"A0016C880162017C3686B18A3D4780", 31},
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
		{"C200B40A82", 3},
		{"04005AC33890", 54},
		{"880086C3E88112", 7},
		{"CE00C43D881120", 9},
		{"D8005AC2A8F0", 1},
		{"F600BC2D8F", 0},
		{"9C005AC2F8F0", 0},
		{"9C0141080250320F1802104A08", 1},
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



func TestByteConvert(t *testing.T) {
	val, err := hexToBinaryString("38006F45291200")
	if err != nil {
		t.Error(err)
		return
	}
	if val != "00111000000000000110111101000101001010010001001000000000" {
		t.Errorf("Expected %s, got %s", "00111000000000000110111101000101001010010001001000000000", val)
	}
}
