defmodule AdventOfCode.Day05 do
  require Graph

  @spec parse_input(String.t()) :: %{rules: [{integer(), integer()}], updates: [[integer()]]}
  defp parse_input(input) do
    String.split(input, "\n\n", trim: true)
    |> then(fn [raw_rules, raw_updates] ->
      rules =
        raw_rules
        |> String.split("\n", trim: true)
        |> Enum.map(fn line -> String.split(line, "|", trim: true) end)
        |> Enum.map(fn [a, b] -> {String.to_integer(a), String.to_integer(b)} end)

      updates =
        raw_updates
        |> String.split("\n", trim: true)
        |> Enum.map(fn line -> String.split(line, ",", trim: true) end)
        |> Enum.map(fn values -> Enum.map(values, fn val -> String.to_integer(val) end) end)

      %{rules: rules, updates: updates}
    end)
  end

  @spec create_adjaceny_list([{integer(), integer()}]) :: %{integer() => [integer()]}
  defp create_adjaceny_list(rules) do
    rules
    |> Enum.reduce({nil, Map.new()}, fn {a, b}, {_, map} ->
      Map.get_and_update(map, a, fn current_list ->
        if current_list == nil, do: {nil, [b]}, else: {current_list, [b | current_list]}
      end)
    end)
    |> elem(1)
  end

  @spec create_topological(%{integer() => [integer()]}, [integer()]) :: %{integer() => integer()}
  defp create_topological(adjaceny_list, update) do
    existing_nodes = Enum.reduce(update, Map.new(), fn a, map -> Map.put(map, a, nil) end)

    Graph.new()
    |> Graph.add_edges(
      Enum.flat_map(update, fn a ->
        if Map.has_key?(adjaceny_list, a),
          do:
            Enum.map(
              Map.get(adjaceny_list, a)
              |> Enum.filter(fn b -> Map.has_key?(existing_nodes, b) end),
              fn b -> {a, b} end
            ),
          else: []
      end)
    )
    |> Graph.topsort()
    |> Enum.with_index()
    |> Enum.reduce(Map.new(), fn {a, idx}, map -> Map.put(map, a, idx) end)
  end

  @spec correct_order?([integer()], %{integer() => integer()}) :: boolean()
  defp correct_order?(order, topological_sort),
    do: order == order_by_topology(order, topological_sort)

  @spec order_by_topology([integer()], %{integer() => integer()}) :: [integer()]
  defp order_by_topology(order, topological_sort) do
    Enum.map(order, fn val -> {val, Map.get(topological_sort, val)} end)
    |> List.keysort(1)
    |> Enum.map(fn {val, _} -> val end)
  end

  @spec part1(String.t()) :: integer()
  def part1(input) do
    %{:rules => rules, :updates => updates} = parse_input(input)
    adjacency_list = create_adjaceny_list(rules)

    updates
    |> Task.async_stream(fn update ->
      topological_sort = create_topological(adjacency_list, update)

      if correct_order?(update, topological_sort) do
        Enum.at(update, div(length(update), 2))
      else
        0
      end
    end)
    |> Enum.reduce(0, fn
      {:ok, value}, acc -> acc + value
      _, acc -> acc
    end)
  end

  @spec part1(String.t()) :: integer()
  def part2(input) do
    %{:rules => rules, :updates => updates} = parse_input(input)
    adjacency_list = create_adjaceny_list(rules)

    updates
    |> Task.async_stream(fn update ->
      topological_sort = create_topological(adjacency_list, update)

      if !correct_order?(update, topological_sort) do
        sorted_update = order_by_topology(update, topological_sort)
        Enum.at(sorted_update, div(length(sorted_update), 2))
      else
        0
      end
    end)
    |> Enum.reduce(0, fn
      {:ok, value}, acc when is_integer(value) -> acc + value
      _, acc -> acc
    end)
  end
end
