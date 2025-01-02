defmodule AdventOfCode.Day14 do
  require Math

  defmodule Robot do
    defstruct [:x, :y, :vx, :vy]
  end

  @regex ~r/p=(\d+),(\d+) v=(-?\d+),(-?\d+)/
  defp parse_line(line) do
    [x, y, vx, vy] =
      Regex.run(@regex, line, capture: :all_but_first) |> Enum.map(&String.to_integer/1)

    %Robot{x: x, y: y, vx: vx, vy: vy}
  end

  defp parse_input(input) do
    String.split(input, "\n", trim: true)
    |> Enum.map(&parse_line/1)
  end

  defp move_robot(robot, x_size, y_size, times),
    do: %Robot{
      x: Integer.mod(robot.x + times * robot.vx, x_size),
      y: Integer.mod(robot.y + times * robot.vy, y_size),
      vx: robot.vx,
      vy: robot.vy
    }

  defp move_robots(robots, x_size, y_size, times),
    do: Enum.map(robots, &move_robot(&1, x_size, y_size, times))

  defp in_quadrant?(robot, {{x1, x2}, {y1, y2}}),
    do: robot.x >= x1 && robot.x <= x2 && robot.y >= y1 && robot.y <= y2

  defp robots_in_quadrant(robots, quadrant),
    do: robots |> Enum.filter(&in_quadrant?(&1, quadrant)) |> length()

  defp quadrant_count(robots, x_size, y_size) do
    quad_1 = {{0, div(x_size, 2) - 1}, {0, div(y_size, 2) - 1}}
    quad_2 = {{div(x_size, 2) + 1, x_size - 1}, {0, div(y_size, 2) - 1}}
    quad_3 = {{0, div(x_size, 2) - 1}, {div(y_size, 2) + 1, y_size - 1}}
    quad_4 = {{div(x_size, 2) + 1, x_size - 1}, {div(y_size, 2) + 1, y_size - 1}}

    [quad_1, quad_2, quad_3, quad_4]
    |> Enum.reduce(1, fn quad, product -> robots_in_quadrant(robots, quad) * product end)
  end

  def part1(input, options \\ [x_size: 101, y_size: 103]) do
    parse_input(input)
    |> move_robots(options[:x_size], options[:y_size], 100)
    |> quadrant_count(options[:x_size], options[:y_size])
  end

  defp variance_x(robots), do: Math.Enum.variance(Enum.map(robots, fn robot -> robot.x end))
  defp variance_y(robots), do: Math.Enum.variance(Enum.map(robots, fn robot -> robot.y end))

  defp chinese_remainder(dividends, divisors) do
    n = Enum.reduce(divisors, 1, fn divisor, acc -> acc * divisor end)
    remaining = Enum.map(divisors, fn divisor -> div(n, divisor) end)
    mod_inv = Enum.zip(divisors, remaining) |> Enum.map(fn {ni, yi} -> Math.mod_inv!(yi, ni) end)

    Integer.mod(
      Enum.zip([dividends, remaining, mod_inv])
      |> Enum.reduce(0, fn {ai, yi, zi}, acc -> acc + ai * yi * zi end),
      n
    )
  end

  def part2(input, options \\ [x_size: 101, y_size: 103]) do
    robots = parse_input(input)

    min_x_variance =
      Enum.min_by(0..options[:x_size], fn times ->
        variance_x(move_robots(robots, options[:x_size], options[:y_size], times))
      end)

    min_y_variance =
      Enum.min_by(0..options[:y_size], fn times ->
        variance_y(move_robots(robots, options[:x_size], options[:y_size], times))
      end)

    chinese_remainder([min_x_variance, min_y_variance], [options[:x_size], options[:y_size]])
  end
end
