import kotlin.math.abs
import kotlin.math.max

fun main() {
    val cases = listOf(0,-1,1)
    val check8Sides = { idx: Pair<Int, Int>, matrix: List<List<Char>>, condition: (value: Char) -> Boolean ->
        (idx.first-1..idx.first+1).any { i ->
            var outerReturnValue = false
            if (i >= 0 && i < matrix.size) {
                outerReturnValue = (idx.second-1..idx.second+1).any { j ->
                    var returnValue = false
                    if (j >= 0 && j < matrix[i].size) {
                        val nearbyValue = matrix[i][j]
                        if (condition(nearbyValue))
                            returnValue = true
                    }
                    returnValue
                }
            }
            outerReturnValue
        }
    }
    fun part1(input: List<String>): Int {
        val parsedArray = input.map { it.toCharArray().toList() }
        val validNumbers = mutableListOf<Int>()
        parsedArray.mapIndexed { i, line ->
            var j = 0
            line.mapIndexed { id, ch ->
                if (id >= j && ch.isDigit()) {
                    j = id
                    while (j < line.size && line[j].isDigit()) {
                        j += 1
                    }
                    var isValid = false
                    for (idx in id..<j) {
                        if (check8Sides(
                                Pair(i, idx),
                                parsedArray
                            ) { value -> !value.isDigit() && value != '.' }
                        ) {
                            val number = line.subList(id, j).joinToString("").toInt()
                            validNumbers.add(number)
                            isValid = true
                            break
                        }
                    }
                }
            }
        }
        return validNumbers.sum()
    }

    operator fun Pair<Int, Int>.minus(adder: Pair<Int, Int>) = Pair(first - adder.first, second - adder.second)
    fun Pair<Int, Int>.sum() = abs(first) + abs(second)

    fun part2(input: List<String>): Int {
        val parsedArray = input.map { it.toCharArray().toList() }
        val mapOfNumbers = mutableMapOf<Pair<Int,Int>, Int>()
        parsedArray.mapIndexed { i, line ->
            var j = 0
            line.mapIndexed { id, ch ->
                if (id >= j && ch.isDigit()) {
                    j = id
                    while (j < line.size && line[j].isDigit()) {
                        j += 1
                    }
                    val number = line.subList(id, j).joinToString("").toInt()
                    (id..<j).forEach { mapOfNumbers[Pair(i, it)] = number }
                }
            }
        }
        return parsedArray.flatMapIndexed { i, line ->
            line.mapIndexed { j, ch ->
                var returnValue = 0
                if (ch == '*') {
                    var firstNumber = -1
                    var firstNumberLocation = Pair(-1, -1)
                    var secondNumber = -1
                    for (deltaI in cases) {
                        for (deltaJ in cases) {
                            val nearbyI = i + deltaI
                            val nearbyJ = j + deltaJ
                            if (nearbyI >= 0
                                && nearbyI < parsedArray.size
                                && nearbyJ >= 0
                                && nearbyJ < parsedArray[nearbyI].size) {
                                mapOfNumbers[Pair(nearbyI, nearbyJ)]?.let {
                                    if (firstNumber == -1) {
                                        firstNumber = it
                                        firstNumberLocation = Pair(nearbyI, nearbyJ)
                                    } else if (secondNumber == -1 && (firstNumberLocation - Pair(nearbyI, nearbyJ)).sum() > 1) {
                                        secondNumber = it
                                    }
                                }
                            }
                        }
                    }
                    if (firstNumber != -1 && secondNumber != -1) {
                        returnValue = firstNumber * secondNumber
                    }
                }
                returnValue
            }
        }.sum()
    }

    // test if implementation meets criteria from the description, like:
    val testInput = readInput("Day03_test")
    check(part1(testInput) == 4361)
    check(part2(testInput) == 467835)

    val input = readInput("Day03")
    part1(input).println()
    part2(input).println()
}