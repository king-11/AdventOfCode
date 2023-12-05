fun main() {
    data class Mappings(val src: Long, val dest: Long, val range: Long)
    val parseThree = { line:String ->
        val values = line.split(" ")
        Mappings(values[1].toLong(), values[0].toLong(), values[2].toLong())
    }
    fun part1(input: List<String>): Long {
        var seeds = input
            .first()
            .substringAfter("seeds: ")
            .split(" ")
            .filter { it.isNotEmpty() }
            .map { it.toLong() }
            .associateWith { it }

        val inputGroups = input
            .subList(1, input.size)
            .joinToString("\n")
            .trim('\n')
            .split("\n\n")

        inputGroups.map {
            val mutableSeeds = mutableMapOf<Long, Long>()
            it
                .lines()
                .asSequence()
                .drop(1)
                .map(parseThree)
                .forEach { mapping ->
                    seeds.forEach { (key, value) ->
                        if (value in mapping.src..<(mapping.src + mapping.range)) {
                            mutableSeeds[key] = (value - mapping.src) + mapping.dest
                        }
                    }
                }
            seeds = seeds
                .map { (key, value) ->
                    key to (mutableSeeds[key] ?: value)
                }
                .toMap()
        }
        return seeds
            .map { it.value }
            .minOrNull() ?: error("No solution found")
    }

    operator fun LongRange.contains(value: ClosedRange<Long>): Boolean {
        return value.start in this && value.endInclusive in this
    }
    operator fun LongRange.plus(value: Long): LongRange {
        return (this.first + value)..(this.last + value)
    }
    fun part2(input: List<String>): Long {
        var seeds = input
            .first()
            .substringAfter("seeds: ")
            .split(" ")
            .filter { it.isNotEmpty() }
            .windowed(2)
            .map { (first, second) -> first.toLong()..<first.toLong() + second.toLong() }
            .associateWith { it }
            .toSortedMap(compareBy { it.first })

        val inputGroups = input
            .subList(1, input.size)
            .joinToString("\n")
            .trim('\n')
            .split("\n\n")

        inputGroups.map {
            val copyOfSeeds = seeds.toMutableMap()
            it
                .lines()
                .asSequence()
                .drop(1)
                .map(parseThree)
                .forEach { mapping ->
                    seeds.forEach { (key, value) ->
                        val mapRange = mapping.src..<(mapping.src + mapping.range)
                        // TODO(king-11): use map range properly
                        if (value.last in mapRange && value.first in mapRange) {
                            copyOfSeeds[key] = (value.first - mapRange.first) + mapping.dest..(value.first - mapRange.first) + mapping.dest
                        } else if (value.last in mapRange) {
                            copyOfSeeds.remove(key)
                        } else if (value.first in mapRange) {

                        }
                    }
                }
            seeds = seeds
                .map { (key, value) ->
                    key to (copyOfSeeds[key] ?: value)
                }
                .toMap()
                .toSortedMap(compareBy { it -> it.first })
        }
//        return seeds
//            .map { it.value }
//            .minOrNull() ?: error("No solution found")
        return input.size.toLong()
    }

    // test if implementation meets criteria from the description, like:
    val testInput = readInput("Day05_test")
    val input = readInput("Day05")
    check(part1(testInput) == 35L)
    part1(input).println()

    check(part2(testInput) == 46L)
    part2(input).println()
}