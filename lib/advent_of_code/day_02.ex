defmodule AdventOfCode.Day02 do
  defp parse_input(input) do
    String.split(input, "\n", trim: true)
    |> Enum.map(fn line ->
      String.split(line, " ", trim: true)
      |> Enum.map(fn value -> String.to_integer(value) end)
    end)
  end

  defp is_safe(values) do
    sorted_values = Enum.sort(values)

    Enum.chunk_every(values, 2, 1, :discard)
    |> Enum.all?(fn [a, b] -> abs(a - b) >= 1 && abs(a - b) <= 3 end) &&
      (sorted_values == values || sorted_values == Enum.reverse(values))
  end

  def part1(input) do
    parse_input(input)
    |> Enum.filter(&is_safe/1)
    |> length()
  end

  def part2(input) do
    parse_input(input)
    |> Enum.filter(fn values ->
      Enum.with_index(values)
      |> Enum.any?(fn {_, i} -> is_safe(List.delete_at(values, i)) end)
    end)
    |> length()
  end
end
