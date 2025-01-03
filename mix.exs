defmodule AdventOfCode.MixProject do
  use Mix.Project

  def project do
    [
      app: :advent_of_code,
      version: "0.1.0",
      elixir: "~> 1.17",
      start_permanent: Mix.env() == :prod,
      deps: deps()
    ]
  end

  # Run "mix help compile.app" to learn about applications.
  def application do
    [
      extra_applications: [:logger, :inets, :eex]
    ]
  end

  # Run "mix help deps" to learn about dependencies.
  defp deps do
    [
      {:httpoison, "~> 2.0"},
      {:benchee, "~> 1.3"},
      {:tz, "~> 0.26.5"},
      {:libgraph, "~> 0.16.0"},
      {:pex, "~> 0.3.0"},
      {:arrays, "~> 2.1.1"},
      {:nx, "~> 0.9.2"},
      {:math, "~> 0.7.0"}
    ]
  end
end
