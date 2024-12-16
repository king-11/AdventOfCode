defmodule AdventOfCode.Day10Test do
  use ExUnit.Case, async: true

  import AdventOfCode.Day10

  setup do
    [
      input: """
      89010123
      78121874
      87430965
      96549874
      45678903
      32019012
      01329801
      10456732
      """
    ]
  end

  test "part1", %{input: input} do
    result = part1(input)

    assert result == 36
  end

  test "part2", %{input: input} do
    result = part2(input)

    assert result == 81
  end
end
