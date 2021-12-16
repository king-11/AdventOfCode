package day12

import (
	"bufio"
	"os"
	"strings"
)

type cave struct {
	name    string
	isLarge bool
}

type tunnel struct {
	start cave
	end   cave
}

func getScanner(filename string) (*bufio.Scanner, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	return scanner, nil
}

func isUpper(s string) bool {
	return strings.ToUpper(s) == s
}

func processLines(scanner *bufio.Scanner) ([]*tunnel, error) {
	lines := make([]*tunnel, 0)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "-")
		t := &tunnel{cave{line[0], isUpper(line[0])}, cave{line[1], isUpper(line[1])}}
		lines = append(lines, t)
	}
	return lines, scanner.Err()
}

const START_CAVE_ID int = 1
const END_CAVE_ID int = 2
const MIN_SMALL_CAVE_ID int = 79 // prime at index 21 is 79

func getCaveId(caveName string, bigCavePrimeIndex int, smallCavePrimeIndex int) (int, int, int) {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199, 211, 223, 227, 229, 233, 239, 241, 251, 257, 263, 269, 271, 277, 281, 283, 293, 307, 311, 313, 317, 331, 337, 347, 349, 353, 359, 367, 373, 379, 383, 389, 397, 401, 409, 419, 421, 431, 433, 439, 443, 449, 457, 461, 463, 467, 479, 487, 491, 499, 503, 509, 521, 523, 541}
	caveId := 0
	if caveName == "start" {
		caveId = START_CAVE_ID
	} else if caveName == "end" {
		caveId = END_CAVE_ID
	} else if caveName[0] > 90 { // Lowercase is > 90 (A = 65, Z = 90, a = 97, z = 122)
		caveId = primes[smallCavePrimeIndex]
		smallCavePrimeIndex += 1
	} else {
		caveId = primes[bigCavePrimeIndex]
		bigCavePrimeIndex += 1
	}
	return caveId, bigCavePrimeIndex, smallCavePrimeIndex
}

func convertLinesToNeighbours(t []*tunnel) map[int][]int {
	bigCavePrimeIndex := 1
	smallCavePrimeIndex := 21
	neighbours := map[int][]int{}
	idLookup := map[string]int{}

	for _, tun := range t {
		for _, caveName := range []string{tun.start.name, tun.end.name} {
			_, ok := idLookup[caveName]
			if !ok {
				caveId := 0
				caveId, bigCavePrimeIndex, smallCavePrimeIndex = getCaveId(caveName, bigCavePrimeIndex, smallCavePrimeIndex)
				idLookup[caveName] = caveId
				neighbours[caveId] = []int{}
			}
		}

		fromId := idLookup[tun.start.name]
		toId := idLookup[tun.end.name]

		// Add 'to' location to 'from' cave destinations. Can't go back to start or exit from the end
		if toId != START_CAVE_ID && fromId != END_CAVE_ID {
			neighbours[fromId] = append(neighbours[fromId], toId)
		}

		// Reverse: add 'from' location to 'to' cave destinations. Can't reverse from the end or go back to start
		if toId != END_CAVE_ID && fromId != START_CAVE_ID {
			neighbours[toId] = append(neighbours[toId], fromId)
		}
	}
	return neighbours
}
