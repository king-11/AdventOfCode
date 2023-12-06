import kotlin.math.sqrt

fun main() {
    val numberRegex = Regex("\\d+")
    val quadraticValue = { x: Long, t: Long, d: Long -> x * x - x * t + d }
    val solutionOfQuadratic = { t: Long, d: Long ->
        Pair(
            ((t - sqrt((t * t - 4 * d).toDouble())) / 2).toLong(),
            ((t + sqrt((t * t - 4 * d).toDouble())) / 2).toLong()
        )
    }
    fun part1(input: List<String>): Long {
        val raceTimes = numberRegex.findAll(input[0]).map { it.value.toLong() }.toList()
        val requiredDistance = numberRegex.findAll(input[1]).map { it.value.toLong() }.toList()
        return raceTimes
            .mapIndexed { index, time -> time to requiredDistance[index] }
            .map { (time, distance) ->
                var (first, second) = solutionOfQuadratic(time, distance)
                while (first < 0 || quadraticValue(first, time, distance) >= 0) first++
                while (quadraticValue(second, time, distance) >= 0) second--
                second - first + 1
            }
            .fold(1) { acc, i -> acc * i }
    }

    fun part2(input: List<String>): Long {
        val time = numberRegex.findAll(input[0]).map { it.value }.joinToString("").toLong()
        val distance = numberRegex.findAll(input[1]).map { it.value }.joinToString("").toLong()
        var (first, second) = solutionOfQuadratic(time.toLong(), distance.toLong())
        while (first < 0 || quadraticValue(first, time, distance) >= 0) first++
        while (quadraticValue(second, time, distance) >= 0) second--
        return second - first + 1
    }

    // test if implementation meets criteria from the description, like:
    val testInput = readInput("Day06_test")
    val input = readInput("Day06")
    check(part1(testInput) == 288L)
    part1(input).println()

    check(part2(testInput) == 71503L)
    part2(input).println()
}