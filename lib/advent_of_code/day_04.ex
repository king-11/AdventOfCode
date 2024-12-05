defmodule AdventOfCode.Day04 do
  defp parse_grid(input) do
    input
    |> String.split("\n", trim: true)
    |> Enum.with_index()
    |> Enum.reduce(Map.new(), &parse_row/2)
  end

  defp parse_row({row, row_no}, map) do
    row
    |> String.graphemes()
    |> Enum.with_index()
    |> Enum.reduce(map, fn {col, col_no}, map ->
      Map.put(map, {row_no, col_no}, col)
    end)
  end

  defp find_coords(map, character) do
    Map.filter(map, fn {_, value} -> value == character end)
    |> Map.keys()
  end

  defp matches?(grid, {i, j}, {di, dj}, offset, character) do
    case Map.fetch(grid, {i + di * offset, j + dj * offset}) do
      {:ok, value} -> value == character
      _ -> false
    end
  end

  defp find_xmas(grid, start) do
    [{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}]
    |> Enum.filter(fn delta ->
      matches?(grid, start, delta, 1, "M") && matches?(grid, start, delta, 2, "A") &&
        matches?(grid, start, delta, 3, "S")
    end)
  end

  defp find_max(grid, start) do
    [[{-1, -1}, {1, 1}], [{-1, 1}, {1, -1}]]
    |> Enum.all?(fn [m, s] ->
      (matches?(grid, start, m, 1, "M") && matches?(grid, start, s, 1, "S")) ||
        (matches?(grid, start, s, 1, "M") && matches?(grid, start, m, 1, "S"))
    end)
  end

  def part1(input) do
    grid = parse_grid(input)

    find_coords(grid, "X")
    |> Enum.flat_map(fn start -> find_xmas(grid, start) end)
    |> length()
  end

  def part2(input) do
    grid = parse_grid(input)

    find_coords(grid, "A")
    |> Enum.filter(fn start -> find_max(grid, start) end)
    |> length()
  end
end
