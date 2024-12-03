defmodule AdventOfCode.Day03 do
  def part1(input) do
    Regex.scan(~r"mul\((\d+),(\d+)\)", input, capture: :all)
    |> Enum.map(fn [_, a, b] -> elem(Integer.parse(a), 0) * elem(Integer.parse(b), 0) end)
    |> Enum.sum()
  end

  def part2(input) do
    Regex.scan(~r"((do\(\)|don't\(\)).*?)?mul\((\d+),(\d+)\)", input, capture: :all)
    |> Enum.map(fn [_, _, op, a, b] ->
      {op, elem(Integer.parse(a), 0) * elem(Integer.parse(b), 0)}
    end)
    |> Enum.reduce({"do()", 0}, fn {op, val}, {ongoing_op, acc} ->
      if (op == "" && ongoing_op == "do()") || op == "do()",
        do: {"do()", acc + val},
        else: {"don't()", acc}
    end)
    |> then(fn {_, val} -> val end)
  end
end
