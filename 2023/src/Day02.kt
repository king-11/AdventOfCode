import kotlin.math.max

fun main() {
    val gameRegex = Regex("Game (?<gameNumber>\\d+)")
    val ballRegex = Regex("(?<count>\\d+) (?<color>red|green|blue)[,;]?")
    fun part1(input: List<String>): Int {
        val counts = mapOf(Pair("red", 12), Pair("green", 13), Pair("blue", 14))
        return input.sumOf { line ->
            val gameMatch = gameRegex.find(line)
            var returnValue = gameMatch?.groups?.get("gameNumber")!!.value.toInt()
            val matches = ballRegex.findAll(line)
            matches.forEach { match ->
                val count = match.groups["count"]!!.value.toInt()
                val color = match.groups["color"]!!.value
                if (count > counts[color]!!) {
                    returnValue = 0
                }
            }
            returnValue
        }
    }

    fun part2(input: List<String>): Int {
        return input.map { line ->
            val maxCount = mutableMapOf(Pair("red", 0), Pair("green", 0), Pair("blue", 0))
            val matches = ballRegex.findAll(line)
            matches.forEach { match ->
                val count = match.groups["count"]!!.value.toInt()
                val color = match.groups["color"]!!.value
                maxCount[color] = max(maxCount[color]!!, count)
            }
            maxCount.keys.fold(1) { acc, s -> acc * maxCount[s]!!  }
        }.sum()
    }

    // test if implementation meets criteria from the description, like:
    val testInput = readInput("Day02_test")
    check(part1(testInput) == 8)
    check(part2(testInput) == 2286)

    val input = readInput("Day02")
    part1(input).println()
    part2(input).println()
}