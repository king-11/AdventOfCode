defmodule AdventOfCode.Day05 do
  require Graph

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

  defp create_adjaceny_list(rules) do
    rules
    |> Enum.reduce({nil, Map.new()}, fn {a, b}, {_, map} ->
      Map.get_and_update(map, a, fn current_list ->
        if current_list == nil, do: {nil, [b]}, else: {current_list, current_list ++ [b]}
      end)
    end)
    |> elem(1)
  end

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

  defp correct_order?(order, topological_sort),
    do: order == order_by_topology(order, topological_sort)

  defp order_by_topology(order, topological_sort) do
    Enum.map(order, fn val -> {val, Map.get(topological_sort, val)} end)
    |> List.keysort(1)
    |> Enum.map(fn {val, _} -> val end)
  end

  def part1(input) do
    %{:rules => rules, :updates => updates} = parse_input(input)
    adjaceny_list = create_adjaceny_list(rules)

    updates
    |> Enum.map(fn update -> {update, create_topological(adjaceny_list, update)} end)
    |> Enum.filter(fn {update, g} -> correct_order?(update, g) end)
    |> Enum.map(fn {update, _} -> Enum.at(update, div(length(update), 2)) end)
    |> Enum.sum()
  end

  def part2(input) do
    %{:rules => rules, :updates => updates} = parse_input(input)
    adjaceny_list = create_adjaceny_list(rules)

    updates
    |> Enum.map(fn update -> {update, create_topological(adjaceny_list, update)} end)
    |> Enum.filter(fn {update, g} -> !correct_order?(update, g) end)
    |> Enum.map(fn {update, g} -> order_by_topology(update, g) end)
    |> Enum.map(fn sorted_update -> Enum.at(sorted_update, div(length(sorted_update), 2)) end)
    |> Enum.sum()
  end
end
