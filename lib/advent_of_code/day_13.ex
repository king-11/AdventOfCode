defmodule AdventOfCode.Day13 do
  require Nx

  defmodule Machine do
    defstruct [:a_move, :b_move, :prize_location]
  end

  defp parse_input(input) do
    String.split(input, "\n\n", trim: true)
    |> Enum.map(fn block ->
      [line_a, line_b, line_prize] = String.split(block, "\n", trim: true, parts: 3)

      [ax, ay] =
        Regex.run(~r/Button A: X\+(\d+), Y\+(\d+)/, line_a, capture: :all_but_first)
        |> Enum.map(&String.to_integer/1)

      [bx, by] =
        Regex.run(~r/Button B: X\+(\d+), Y\+(\d+)/, line_b, capture: :all_but_first)
        |> Enum.map(&String.to_integer/1)

      [px, py] =
        Regex.run(~r/Prize: X=(\d+), Y=(\d+)/, line_prize, capture: :all_but_first)
        |> Enum.map(&String.to_integer/1)

      %Machine{
        a_move: {ax, ay},
        b_move: {bx, by},
        prize_location: {px, py}
      }
    end)
  end

  defp solve_machine(machine) do
    {ax, ay} = machine.a_move
    {bx, by} = machine.b_move
    {px, py} = machine.prize_location

    a = Nx.tensor([[ax, bx], [ay, by]], type: :f64)
    b = Nx.tensor([px, py], type: :f64)
    res = Nx.LinAlg.solve(a, b) |> Nx.round()

    if check_solution(res, a, b),
      do: {:ok, Nx.to_number(Nx.sum(Nx.dot(res, Nx.tensor([3.0, 1.0], type: :f64))))},
      else: {:error, nil}
  end

  defp check_solution(res, a, b) do
    Nx.all(Nx.equal(Nx.dot(a, res), b)) == Nx.tensor(1, type: :u8)
  end

  def part1(input) do
    parse_input(input)
    |> Enum.map(&solve_machine/1)
    |> Enum.filter(fn {res, _} -> res == :ok end)
    |> Enum.reduce(0.0, fn {_, val}, acc -> acc + val end)
  end

  def part2(input) do
    parse_input(input)
    |> Enum.map(fn machine ->
      %Machine{
        a_move: machine.a_move,
        b_move: machine.b_move,
        prize_location:
          {elem(machine.prize_location, 0) + 10_000_000_000_000,
           elem(machine.prize_location, 1) + 10_000_000_000_000}
      }
    end)
    |> Enum.map(&solve_machine/1)
    |> Enum.filter(fn {res, _} -> res == :ok end)
    |> Enum.reduce(0.0, fn {_, val}, acc -> acc + val end)
  end
end
