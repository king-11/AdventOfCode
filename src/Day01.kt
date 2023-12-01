import java.lang.StringBuilder

fun main() {
    fun part1(input: List<String>): Int {
       return input.sumOf { s ->
           val firstDigit = s.find { it.isDigit() }
           val lastDigit = s.findLast { it.isDigit() }
           "${firstDigit}${lastDigit}".toInt()
       }
    }

    fun part2(input: List<String>): Int {
        val numbers = mapOf(Pair("one", 1), Pair("two", 2), Pair("three", 3), Pair("four", 4), Pair("five", 5), Pair("six", 6), Pair("seven", 7), Pair("eight", 8), Pair("nine", 9))
        return input.sumOf { line ->
            val numbersFound = arrayListOf<Int>()
            val window = StringBuilder()
            line.forEach { ch ->
                window.append(ch)
                if (ch.isDigit()) {
                    numbersFound.add(ch.code - '0'.code)
                    window.clear()
                }
                while (window.isNotEmpty() && !numbers.any { it.key.commonPrefixWith(window.toString()).length == window.length }) {
                    window.delete(0, 1)
                }
                numbers[window.toString()]?.let {
                    numbersFound.add(it)
                    window.delete(0, 1)
                }
            }
            val value = "${numbersFound.first()}${numbersFound.last()}".toInt()
            value
        }
    }

    // test if implementation meets criteria from the description, like:
    val testInput = readInput("Day01_test")
    check(part1(testInput) == 142)
    val testInput1 = readInput("Day01_test_1")
    check(part2(testInput1) == 281)

    val sampleTest = listOf("eighthree", "oneight")
    check(part2(sampleTest) == (83 + 18))

    val input = readInput("Day01")
    part1(input).println()
    part2(input).println()
}