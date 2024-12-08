defmodule AdventOfCode.Day06 do
  defp parse_character(char) do
    case char do
      "." -> :blank
      "#" -> :block
      "^" -> :up
      ">" -> :right
      "<" -> :left
      "v" -> :down
      "X" -> :visited
    end
  end

  defp parse_input(input) do
    input
    |> String.split("\n", trim: true)
    |> Enum.with_index()
    |> Enum.reduce(Map.new(), fn {line, i}, map ->
      String.split(line, "", trim: true)
      |> Enum.with_index()
      |> Enum.reduce(map, fn {char, j}, map -> Map.put(map, {i, j}, parse_character(char)) end)
    end)
  end

  defp rotate_list(list, n) do
    {left, right} = Enum.split(list, n)
    right ++ left
  end

  defp possible_moves_ordered(direction) do
    move = [{-1, 0}, {0, 1}, {1, 0}, {0, -1}]

    case direction do
      :up -> move
      :right -> rotate_list(move, 1)
      :down -> rotate_list(move, 2)
      :left -> rotate_list(move, 3)
    end
  end

  defp find_next(grid, {i, j}, direction) do
    possible_moves_ordered(direction)
    |> Enum.map(fn {di, dj} -> {i + di, j + dj} end)
    |> Enum.find(nil, fn move ->
      Map.fetch(grid, move) != {:ok, :block} || Map.fetch(grid, move) == :error
    end)
  end

  defp find_direction({i, j}, {ni, nj}) do
    case {ni - i, nj - j} do
      {0, 1} -> :right
      {1, 0} -> :down
      {0, -1} -> :left
      {-1, 0} -> :up
    end
  end

  defp modify_grid(grid, {i, j}, direction) do
    next = find_next(grid, {i, j}, direction)

    modified_grid = Map.update!(grid, {i, j}, fn _ -> :visited end)

    case Map.fetch(grid, next) do
      {:ok, _} -> modify_grid(modified_grid, next, find_direction({i, j}, next))
      _ -> modified_grid
    end
  end

  def part1(input) do
    grid = parse_input(input)

    Enum.find(grid, fn {_, value} -> value == :up end)
    |> then(fn {{i, j}, value} -> modify_grid(grid, {i, j}, value) end)
    |> Map.filter(fn {_, val} -> val == :visited end)
    |> Kernel.map_size()
  end

  def part2(_input) do
  end
end
