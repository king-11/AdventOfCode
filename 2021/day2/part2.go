package day2

import (
	"strconv"
	"strings"
)

func Part2(filename string) (int, error) {
  data, err := readFile(filename)
  if err != nil {
    return 0, err
  }

  hor, ver, aim := 0, 0, 0
  for _, direction :=  range data {
    split := strings.Split(direction, " ")
    val, err := strconv.Atoi(split[1])
    if err != nil {
      return 0, err
    }

    switch {
      case split[0] == "forward":
        hor += val
        ver += aim*val
      case split[0] == "down":
        aim += val
      case split[0] == "up":
        aim -= val
    }
  }

  return hor*ver, nil
}
