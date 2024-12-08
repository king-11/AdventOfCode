defmodule AdventOfCode.Day06Test do
  use ExUnit.Case, async: true

  import AdventOfCode.Day06

  setup do
    [
      input: """
      ....#.....
      .........#
      ..........
      ..#.......
      .......#..
      ..........
      .#..^.....
      ........#.
      #.........
      ......#...
      """
    ]
  end

  test "part1", %{input: input} do
    result = part1(input)

    assert result == 41
  end

  @tag :skip
  test "part2", %{input: input} do
    result = part2(input)

    assert result == 6
  end
end
