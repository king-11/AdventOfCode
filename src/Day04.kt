import kotlin.math.abs
import kotlin.math.max

fun main() {
    fun part1(input: List<String>): Int {
        return input.sumOf { line ->
            val winningNumbers = line
                .substringAfter(":")
                .substringBefore("|")
                .split(" ")
                .filter { it.isNotEmpty() }
                .map { it.toInt() }
                .toSet()

            val winningMatches = line
                .substringAfter(":")
                .substringAfter("|")
                .split(" ")
                .filter { it.isNotEmpty() }
                .map { it.toInt() }
                .toSet()

            val matchCount = winningNumbers.intersect(winningMatches).size
            if (matchCount > 0) 1.shl(matchCount - 1) else 0
        }
    }

    fun part2(input: List<String>): Int {
        val scratchCards = mutableMapOf<Int, Int>()
        input.mapIndexed { idx, line ->
            if (!scratchCards.containsKey(idx)) {
                scratchCards[idx] = 1
            } else {
                scratchCards[idx] = scratchCards[idx]!! + 1
            }
            val winningNumbers = line
                .substringAfter(":")
                .substringBefore("|")
                .split(" ")
                .filter { it.isNotEmpty() }
                .map { it.toInt() }
                .toSet()

            val winningMatches = line
                .substringAfter(":")
                .substringAfter("|")
                .split(" ")
                .filter { it.isNotEmpty() }
                .map { it.toInt() }
                .toSet()

            val matchCount = winningNumbers.intersect(winningMatches).size
            repeat(matchCount) {
                scratchCards[idx + it + 1] = (scratchCards[idx + it + 1] ?: 0) + scratchCards[idx]!!
            }
        }
        return scratchCards.map { it.value }.sum()
    }

    // test if implementation meets criteria from the description, like:
    val testInput = readInput("Day04_test")
    check(part1(testInput) == 13)
    check(part2(testInput) == 30)

    val input = readInput("Day04")
    part1(input).println()
    part2(input).println()
}