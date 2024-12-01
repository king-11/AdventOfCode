defmodule AdventOfCode.Day01 do
  def parse_input(input) do
    String.split(input, "\n", trim: true)
    |> Enum.map(fn line ->
      line
      |> String.split(" ", trim: true)
      |> Enum.map(&String.to_integer/1)
      |> List.to_tuple()
    end)
    |> Enum.reduce({[], []}, fn {a, b}, {acc_a, acc_b} ->
      {[a | acc_a], [b | acc_b]}
    end)
  end

  def part1(input) do
    parse_input(input)
    |> then(fn {a, b} -> Enum.zip(Enum.sort(a), Enum.sort(b)) end)
    |> Enum.reduce(0, fn {a, b}, acc -> acc + abs(a - b) end)
  end

  def part2(input) do
    parse_input(input)
    |> then(fn {a, b} ->
      map_b = Enum.frequencies(b)

      Enum.filter(a, fn value -> map_b[value] !== nil end)
      |> Enum.reduce(0, fn value, acc -> acc + value * map_b[value] end)
    end)
  end
end
