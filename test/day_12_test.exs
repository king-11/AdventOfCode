defmodule AdventOfCode.Day12Test do
  use ExUnit.Case, async: true

  import AdventOfCode.Day12

  setup do
    [
      input: """
      RRRRIICCFF
      RRRRIICCCF
      VVRRRCCFFF
      VVRCCCJFFF
      VVVVCJJCFE
      VVIVCCJJEE
      VVIIICJJEE
      MIIIIIJJEE
      MIIISIJEEE
      MMMISSJEEE
      """
    ]
  end

  test "part1", %{input: input} do
    result = part1(input)

    assert result == 1930
  end

  test "part2", %{input: input} do
    result = part2(input)

    assert result == 1206
  end
end
