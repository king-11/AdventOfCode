defmodule AdventOfCode.Day07 do
  defp parse_input(input) do
    String.split(input, "\n", trim: true)
    |> Enum.map(fn line -> String.split(line, ":", trim: true) end)
    |> Enum.map(fn [value, raw_list] ->
      {String.to_integer(value),
       String.split(raw_list, " ", trim: true)
       |> Enum.map(fn raw_val -> String.to_integer(raw_val) end)}
    end)
  end

  defp join_numbers(a, b), do: Integer.pow(10, length(Integer.digits(b, 10))) * a + b

  @default_options [concat: true]
  defp is_possible?(value, numbers, options \\ [])
  defp is_possible?(value, [head1, head2 | tail], options) do
    options = Keyword.merge(@default_options, options)
    is_possible?(value, [head1 + head2 | tail], options) ||
    is_possible?(value, [head1 * head2 | tail], options) ||
    (options[:concat] && is_possible?(value, [join_numbers(head1, head2) | tail]))
  end
  defp is_possible?(value, [head1], _options), do: value == head1

  def part1(input) do
    parse_input(input)
    |> Enum.filter(fn {value, numbers} -> is_possible?(value, numbers, concat: false) end)
    |> Enum.reduce(0, fn {value, _}, acc -> value + acc end)
  end

  def part2(input) do
    parse_input(input)
    |> Enum.filter(fn {value, numbers} -> is_possible?(value, numbers) end)
    |> Enum.reduce(0, fn {value, _}, acc -> value + acc end)
  end
end
