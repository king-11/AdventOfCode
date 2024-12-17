defmodule AdventOfCode.Day11Test do
  use ExUnit.Case, async: true

  import AdventOfCode.Day11

  setup do
    [
      input: """
      125 17
      """
    ]
  end

  test "part1", %{input: input} do
    result = part1(input)

    assert result == 55312
  end

  test "part2", %{input: input} do
    result = part2(input)

    assert result == 65601038650482
  end
end
