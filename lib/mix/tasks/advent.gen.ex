defmodule Mix.Tasks.Advent.Gen do
  use Mix.Task
  require Mix.Generator

  @shortdoc "Generates source files for a new year of Advent of Code puzzles"

  @moduledoc """
  # USAGE
  ```
  mix advent.gen <--year <year>>
  ```

  # DESCRIPTION
  Generates source files for a new year of Advent of Code puzzles and populates them with boilerplate code.

  ```
  /
  |- lib/
  | |- advent_of_code/
  | | |- day_01.ex
  | | |- day_02.ex
  | | |- ...
  | | |- day_25.ex
  |- test/
  | |- advent_of_code/
  | | |- day_01_test.ex
  | | |- day_02_test.ex
  | | |- ...
  | | |- day_25_test.ex
  ```
  """

  @days 1..25

  @impl Mix.Task
  def run(_args) do
    generate()
  end

  defp generate() do
    solution_dir = Path.join(lib_root_dir(), "advent_of_code")
    test_dir = test_root_dir()

    Enum.each([solution_dir, test_dir], &Mix.Generator.create_directory/1)

    Enum.each(
      @days,
      &Mix.Generator.create_file(
        Path.join(
          solution_dir,
          :io_lib.format("day_~2..0B.ex", [&1])
        ),
        solution_template(day: &1)
      )
    )

    Enum.each(
      @days,
      &Mix.Generator.create_file(
        Path.join(
          test_dir,
          :io_lib.format("day_~2..0B_test.exs", [&1])
        ),
        test_template(day: &1)
      )
    )
  end

  defp lib_root_dir, do: Path.join(File.cwd!(), "lib")
  defp test_root_dir, do: Path.join(File.cwd!(), "test")

  Mix.Generator.embed_template(:solution, """
  defmodule AdventOfCode.Day<%= :io_lib.format("~2..0B", [@day]) %> do
    def part1(_input) do
    end

    def part2(_input) do
    end
  end
  """)

  Mix.Generator.embed_template(:test, """
  defmodule AdventOfCode.Day<%= :io_lib.format("~2..0B", [@day]) %>Test do
    use ExUnit.Case, async: true

    import AdventOfCode.Day<%= :io_lib.format("~2..0B", [@day]) %>

    setup do
      [
        input: \"""
        \"""
      ]
    end

    @tag :skip
    test "part1", %{input: input} do
      result = part1(input)

      assert result
    end

    @tag :skip
    test "part2", %{input: input} do
      result = part2(input)

      assert result
    end
  end
  """)
end
