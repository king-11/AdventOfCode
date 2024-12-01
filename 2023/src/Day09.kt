fun main() {
  fun part1(input: List<String>): Int {
    return input
      .map { line -> line.split(" ").map { it.toInt() } }
      .sumOf {
        val stack = ArrayDeque<Int>()
        var currentList = it.toList()
        while (currentList.toSet().size != 1) {
          stack.add(currentList.last())
          val newList = currentList
            .mapIndexedNotNull { index, value -> if (index == 0) null else value - currentList[index - 1] }
            .toList()
          currentList = newList
        }
        stack.add(currentList.last())
        stack.sum()
      }
  }

  fun part2(input: List<String>): Int {
    return input
      .map { line -> line.split(" ").map { it.toInt() } }
      .sumOf {
        val dequeue = ArrayDeque<Int>()
        var currentList = it.toList()
        while (currentList.toSet().size != 1) {
          dequeue.add(currentList.first())
          val newList = currentList
            .mapIndexedNotNull { index, value -> if (index == 0) null else value - currentList[index - 1] }
            .toList()
          currentList = newList
        }
        dequeue.add(currentList.first())

        var top = dequeue.removeLast()
        while (dequeue.isNotEmpty()) {
          val next = dequeue.removeLast()
          top = next - top
        }
        top
      }
  }

  // test if implementation meets criteria from the description, like:
  val testInput = readInput("Day09_test")
  val input = readInput("Day09")
  check(part1(testInput) == 114)
  part1(input).println()

  check(part2(testInput) == 2)
  part2(input).println()
}