package day3

import (
	"bufio"
	"os"
	"strings"
)

func readFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		val := strings.Trim(scanner.Text(), " ")
		lines = append(lines, val)
	}
	if err = scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func binaryToInt(b []byte) int {
	var result int
	for _, v := range b {
		result = result<<1 + int(v-'0')
	}
	return result
}

func Part1(filename string) (int, error) {
	lines, err := readFile(filename)
	if err != nil {
		return 0, err
	}

	n, size := len(lines), len(lines[0])
	gamma, epilson := make([]byte, size), make([]byte, size)
	for i := 0; i < size; i++ {
		ones, zeroes := 0, 0
		for j := 0; j < n; j++ {
			if lines[j][i] == '1' {
				ones++
			} else {
				zeroes++
			}
		}
		if ones >= zeroes {
			gamma[i] = '1'
			epilson[i] = '0'
		} else {
			gamma[i] = '0'
			epilson[i] = '1'
		}
	}

	return binaryToInt(gamma)*binaryToInt(epilson), nil
}
