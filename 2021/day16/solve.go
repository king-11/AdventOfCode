package day16

import (
	"errors"
)

func Part1(filename string) (int, error) {
	scanner, err := getScanner(filename)
	if err != nil {
		return 0, err
	}

	if !scanner.Scan() {
		return 0, errors.New("no input")
	}

	rawHex := scanner.Text()
	binary, err := hexToBinaryString(rawHex)
	if err != nil {
		return 0, err
	}
	b := Binary(binary)
	packets, _, err := newPacket(b)
	if err != nil {
		return 0, err
	}

	return packets.versionSum(), nil
}

func Part2(filename string) (int, error) {
	scanner, err := getScanner(filename)
	if err != nil {
		return 0, err
	}

	if !scanner.Scan() {
		return 0, errors.New("no input")
	}

	rawHex := scanner.Text()
	binary, err := hexToBinaryString(rawHex)
	if err != nil {
		return 0, err
	}
	b := Binary(binary)
	packets, _, err := newPacket(b)
	if err != nil {
		return 0, err
	}

	return packets.evaluateExpression(), nil
}
