defmodule AdventOfCode.Day10 do
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
      Map.put(map, {row_no, col_no}, String.to_integer(col))
    end)
  end

  defp possible_moves(grid, {i, j}) do
    [{1, 0}, {0, 1}, {-1, 0}, {0, -1}]
    |> Enum.filter(fn {di, dj} -> Map.has_key?(grid, {i + di, j + dj}) end)
    |> Enum.map(fn {di, dj} -> {i + di, j + dj} end)
  end

  defp find_next(grid, coord, value) do
    if value == 9,
      do: [coord],
      else:
        possible_moves(grid, coord)
        |> Enum.filter(fn {i, j} -> Map.fetch!(grid, {i, j}) == value + 1 end)
        |> Enum.flat_map(fn {i, j} -> find_next(grid, {i, j}, value + 1) end)
  end

  def part1(input) do
    grid = parse_grid(input)

    Enum.filter(grid, fn {_, value} -> value == 0 end)
    |> Enum.map(fn {coord, _} -> length(Enum.uniq(find_next(grid, coord, 0))) end)
    |> Enum.sum()
  end

  def part2(input) do
    grid = parse_grid(input)

    Enum.filter(grid, fn {_, value} -> value == 0 end)
    |> Enum.map(fn {coord, _} -> length(find_next(grid, coord, 0)) end)
    |> Enum.sum()
  end
end
