defmodule AdventOfCode.Day11 do
  defp parse_input(input) do
    String.split(input, " ", trim: true)
    |> Enum.map(fn val -> String.to_integer(String.trim(val)) end)
    |> Enum.reduce(Map.new(), fn val, acc -> Map.update(acc, val, 1, &(&1 + 1)) end)
  end

  defp digits(number) do
    case number do
      0 -> 1
      n -> 1 + floor(:math.log10(n))
    end
  end

  defp split_number(number) do
    digit_count = div(digits(number), 2)
    first = div(number, Integer.pow(10, digit_count))
    second = rem(number, Integer.pow(10, digit_count))
    [first, second]
  end

  defp blink([], acc), do: acc
  defp blink([{0, n} | rest], acc),
    do: blink(rest, Map.update(acc, 1, n, &(&1 + n)))
  defp blink([{x, n} | rest], acc) do
    case rem(digits(x), 2) do
      0 ->
        [first, second] = split_number(x)
        acc =
          acc
          |> Map.update(first, n, &(&1 + n))
          |> Map.update(second, n, &(&1 + n))
          blink(rest, acc)
      1 ->
        blink(rest, Map.update(acc, x * 2024, n, &(&1 + n)))
    end
  end

  defp blink_times(stones, 0), do: stones
  defp blink_times(stones, times) do
    stones
    |> Map.to_list()
    |> blink(%{})
    |> blink_times(times - 1)
  end

  def part1(input) do
    parse_input(input)
    |> blink_times(25)
    |> Enum.reduce(0, fn {_, value}, acc -> acc + value end)
  end

  def part2(input) do
    parse_input(input)
    |> blink_times(75)
    |> Enum.reduce(0, fn {_, value}, acc -> acc + value end)
  end
end
