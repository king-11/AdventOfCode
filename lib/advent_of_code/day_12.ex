defmodule AdventOfCode.Day12 do
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
    |> Enum.reduce(map, fn {col, col_no}, map -> Map.put(map, {row_no, col_no}, col) end)
  end

  defp moves() do
    [{0, -1}, {-1, 0}, {0, 1}, {1, 0}]
  end

  defp add_delta({i, j}, {di, dj}) do
    {i + di, j + dj}
  end

  defp points_between({i1, j1}, {i2, j2}) do
    {{i1, j2}, {i2, j1}}
  end

  defp possible_moves({i, j}) do
    moves()
    |> Enum.map(fn move -> add_delta({i, j}, move) end)
  end

  defp neighbours(grid, point) do
    value = Map.fetch!(grid, point)

    possible_moves(point)
    |> Enum.filter(fn move -> Map.has_key?(grid, move) && Map.fetch!(grid, move) == value end)
  end

  defp not_neighbours(grid, point) do
    possible_moves(point) -- neighbours(grid, point)
  end

  defp exterior_corners(points, node) do
    moves()
    |> Enum.chunk_every(2, 1, moves())
    |> Enum.map(fn [a, b] -> {add_delta(node, a), add_delta(node, b)} end)
    |> Enum.filter(fn {a, b} -> !MapSet.member?(points, a) && !MapSet.member?(points, b) end)
    |> length()
  end

  defp interior_corners(points, node) do
    moves()
    |> Enum.chunk_every(2, 1, moves())
    |> Enum.map(fn [a, b] -> {add_delta(node, a), add_delta(node, b)} end)
    |> Enum.filter(fn {a, b} ->
      {ea, eb} = points_between(a, b)

      MapSet.member?(points, a) && MapSet.member?(points, b) &&
        (!MapSet.member?(points, ea) || !MapSet.member?(points, eb))
    end)
    |> length()
  end

  defp corners(points) do
    Enum.reduce(points, 0, fn point, acc ->
      acc + exterior_corners(points, point) + interior_corners(points, point)
    end)
  end

  defp group(grid, node, visited) do
    adjacent =
      neighbours(grid, node)
      |> Enum.filter(fn neighbour -> !Map.has_key?(visited, neighbour) end)

    visited = Map.update(visited, node, true, fn _ -> true end)
    points = MapSet.new([node])

    Enum.reduce(adjacent, {points, visited}, fn child, {points, visited} ->
      if Map.has_key?(visited, child) do
        {points, visited}
      else
        {internal_points, new_visited} = group(grid, child, visited)
        {MapSet.union(points, internal_points), new_visited}
      end
    end)
  end

  defp perimeter(grid, points) do
    Enum.reduce(points, 0, fn point, acc -> acc + (not_neighbours(grid, point) |> length) end)
  end

  def part1(input) do
    grid = parse_grid(input)

    {final_sum, _} =
      Enum.reduce(Map.keys(grid), {0, %{}}, fn node, {sum, visited} ->
        if Map.has_key?(visited, node) do
          {sum, visited}
        else
          {points, visited} = group(grid, node, visited)
          perimeter = perimeter(grid, points)
          {sum + perimeter * MapSet.size(points), visited}
        end
      end)

    final_sum
  end

  def part2(input) do
    grid = parse_grid(input)

    {final_sum, _} =
      Enum.reduce(Map.keys(grid), {0, %{}}, fn node, {sum, visited} ->
        if Map.has_key?(visited, node) do
          {sum, visited}
        else
          {points, visited} = group(grid, node, visited)
          sides = corners(points)
          {sum + sides * MapSet.size(points), visited}
        end
      end)

    final_sum
  end
end
