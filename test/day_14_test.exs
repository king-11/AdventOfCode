defmodule AdventOfCode.Day14Test do
  use ExUnit.Case, async: true

  import AdventOfCode.Day14

  setup do
    [
      input: """
      p=0,4 v=3,-3
      p=6,3 v=-1,-3
      p=10,3 v=-1,2
      p=2,0 v=2,-1
      p=0,0 v=1,3
      p=3,0 v=-2,-2
      p=7,6 v=-1,-3
      p=3,0 v=-1,-2
      p=9,3 v=2,3
      p=7,3 v=-1,2
      p=2,4 v=2,-3
      p=9,5 v=-3,-3
      """
    ]
  end

  test "part1", %{input: input} do
    result = part1(input, x_size: 11, y_size: 7)

    assert result == 12
  end

  test "part2", %{input: input} do
    result = part2(input, x_size: 11, y_size: 7)

    assert result == 24
  end
end
