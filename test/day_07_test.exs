defmodule AdventOfCode.Day07Test do
  use ExUnit.Case, async: true

  import AdventOfCode.Day07

  setup do
    [
      input: """
      190: 10 19
      3267: 81 40 27
      83: 17 5
      156: 15 6
      7290: 6 8 6 15
      161011: 16 10 13
      192: 17 8 14
      21037: 9 7 18 13
      292: 11 6 16 20
      """
    ]
  end

  test "part1", %{input: input} do
    result = part1(input)

    assert result == 3749
  end

  test "part2", %{input: input} do
    result = part2(input)

    assert result == 11387
  end
end
