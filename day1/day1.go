package day1

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

func readFile(name string) ([]int, error) {
  f, err := os.Open(name)
  if err != nil {
    panic(err)
  }

  scanner := bufio.NewScanner(f)
  scanner.Split(bufio.ScanWords)
  data := make([]int, 0)
  for scanner.Scan() {
    val, err := strconv.Atoi(scanner.Text())
    if err != nil {
      return nil, err
    }
    data = append(data, val)
  }

  if err := scanner.Err(); err != nil {
    return nil, err
  }
  return data, nil
}

func Part1(filename string) (int, error) {
  data, err := readFile(filename)
  if err != nil {
    return 0, err
  }
  prev := math.MaxInt64
  count := 0
  for _, val := range data {
    if prev < val {
      count += 1
    }
    prev = val
  }
  return count, nil
}

func Part2(filename string) (count int, err error) {
  data, err := readFile(filename)
  if err != nil {
    return
  }
  sum := data[0] + data[1] + data[2]
  for i, val := range data {
    if i == 0 || i == 1 || i == 2 {
      continue
    }

    newSum := sum - data[i-3] + val
    if newSum > sum {
      count += 1
    }
    sum = newSum
  }
  return
}
