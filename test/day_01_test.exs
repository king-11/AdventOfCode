defmodule AdventOfCode.Day01Test do
  use ExUnit.Case, async: true

  import AdventOfCode.Day01

  setup do
    [
      input: """
      3   4
      4   3
      2   5
      1   3
      3   9
      3   3
      """
    ]
  end

  test "part1", %{input: input} do
    result = part1(input)

    assert result == 11
  end

  test "part2", %{input: input} do
    result = part2(input)

    assert result == 31
  end
end
