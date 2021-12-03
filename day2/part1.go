package day2

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func readFile(filename string) ([]string, error) {
  data := make([]string, 0)
  f, err := os.Open(filename)
  if err != nil {
    return nil, err
  }
  scanner := bufio.NewScanner(f)
  scanner.Split(bufio.ScanLines)
  for scanner.Scan() {
    data = append(data, scanner.Text())
  }

  return data, nil
}

func Part1(filename string) (int, error) {
  data, err := readFile(filename)
  if err != nil {
    return 0, err
  }

  hor, ver := 0, 0
  for _, direction :=  range data {
    split := strings.Split(direction, " ")
    val, err := strconv.Atoi(split[1])
    if err != nil {
      return 0, err
    }

    switch {
      case split[0] == "forward":
        hor += val
      case split[0] == "down":
        ver += val
      case split[0] == "up":
        ver -= val
    }
  }

  return hor*ver, nil
}
