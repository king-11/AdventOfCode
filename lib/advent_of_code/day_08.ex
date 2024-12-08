defmodule AdventOfCode.Day08 do
  defp parse_input(input) do
    input
    |> String.split("\n", trim: true)
    |> Enum.with_index()
    |> Enum.reduce(Map.new(), fn {line, i}, map ->
      String.split(line, "", trim: true)
      |> Enum.with_index()
      |> Enum.reduce(map, fn {char, j}, map -> Map.put(map, {i, j}, char) end)
    end)
  end

  defp antinode_pair({i1, j1}, {i2, j2}, grid, distance) do
    up = {i1 + distance * (i1 - i2), j1 + distance * (j1 - j2)}
    down = {i2 + distance * (i2 - i1), j2 + distance * (j2 - j1)}
    total = if Map.has_key?(grid, up), do: [up], else: []
    (if Map.has_key?(grid, down), do: [down], else: []) ++ total
  end

  defp antinode_list({i1, j1}, {i2, j2}, grid, distance) do
    total = antinode_pair({i1, j1}, {i2, j2}, grid, distance)
    if length(total) == 0,
      do: total,
      else: total ++ antinode_list({i1, j1}, {i2, j2}, grid, distance + 1)
  end

  defp antinode_location(grid, values, antinode_finder, options) do
    Enum.with_index(values)
    |> Enum.flat_map(fn {{i, j}, index} ->
      Enum.flat_map(Enum.slice(values, (index + 1)..-1//1), fn {i2, j2} -> antinode_finder.({i, j}, {i2, j2}, grid, options[:start]) end)
    end)
  end

  defp antinode_locations(grid, antinode_finder, options \\ [start: 1]) do
    Enum.filter(grid, fn {_, value} -> value != "." end)
    |> Enum.group_by(fn {_, value} -> value end)
    |> Enum.flat_map(fn {_, values} ->
      antinode_location(grid, Enum.map(values, fn {{i, j}, _} -> {i, j} end), antinode_finder, options)
    end)
    |> Enum.uniq()
    |> length()
  end

  def part1(input) do
    parse_input(input)
    |> then(fn grid -> antinode_locations(grid, &antinode_pair/4) end)
  end

  def part2(input) do
    parse_input(input)
    |> then(fn grid -> antinode_locations(grid, &antinode_list/4, start: 0) end)
  end
end
