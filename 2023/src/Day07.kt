import javax.naming.OperationNotSupportedException
import kotlin.math.max

data class Card(private val value: Char) {
    private var wildCard: Boolean = false
    fun setWildCard(wildCard: Boolean): Card {
        this.wildCard = wildCard
        return this
    }
    fun toInt(): Int {
        return when (value) {
            '2' -> 2
            '3' -> 3
            '4' -> 4
            '5' -> 5
            '6' -> 6
            '7' -> 7
            '8' -> 8
            '9' -> 9
            'T' -> 10
            'J' -> if (wildCard) 1 else 11
            'Q' -> 12
            'K' -> 13
            'A' -> 14
            else -> throw OperationNotSupportedException("Unknown card value: $value")
        }
    }
    override fun toString(): String {
        return value.toString()
    }
    operator fun compareTo(other: Card): Int {
        return this.toInt().compareTo(other.toInt())
    }
}
fun Char.toCard() = Card(this)

val hashMap = mutableMapOf<GameHand, Int>()
open class GameHand(private val cards: List<Card>, private val wildCard: Boolean = false): Comparable<GameHand>, Cloneable {
    private val sortedCardSets = cards.toSortedSet(compareBy { it.toInt() })
    fun isFiveOfAKind(): Boolean = sortedCardSets.size == 1
    fun isFourOfAKind(): Boolean = sortedCardSets.size <= 2 && sortedCardSets.any { card -> cards.count { it == card } == 4 }
    fun isFullHouse(): Boolean = sortedCardSets.size == 2 && sortedCardSets.any { card -> cards.count { it == card } == 3 } && sortedCardSets.any { card -> cards.count { it == card } == 2 }
    fun isThreeOfAKind(): Boolean = sortedCardSets.size <= 3 && sortedCardSets.any { card -> cards.count { it == card } == 3 }
    fun isTwoPair(): Boolean {
        val firstMatch = sortedCardSets.find { card -> cards.count { it == card } == 2 }
        val secondMatch = sortedCardSets.find { card -> cards.count { it == card } == 2 && card != firstMatch }
        return firstMatch != null && secondMatch != null
    }
    fun isOnePair(): Boolean = sortedCardSets.size <= 4 && sortedCardSets.any { sortedCard -> cards.count { card -> card == sortedCard } == 2 }
    fun isHighCard(): Boolean = sortedCardSets.size == 5
    fun toInt(): Int {
        return when {
            isFiveOfAKind() -> 7
            isFourOfAKind() -> 6
            isFullHouse() -> 5
            isThreeOfAKind() -> 4
            isTwoPair() -> 3
            isOnePair() -> 2
            isHighCard() -> 1
            else -> 0
        }
    }
    override fun toString(): String = cards.joinToString("")
    private fun replaceAndFindValue(): Int {
        val hand = this.clone()
        if (hashMap.containsKey(hand))
            return hashMap[hand]!!

        if (!hand.cards.contains('J'.toCard())) {
            hashMap[hand] = hand.toInt()
        }
        else {
            val topCards = hand
                .cards
                .filter { it != 'J'.toCard() }
                .groupingBy { it }
                .eachCount()
                .toList()
                .sortedByDescending { (_, value) -> value }

            if (topCards.size <= 1) {
                hashMap[hand] = 7
            }  else {
                val topCardValue = hand.cards.map { if (it == Card('J')) topCards[0].first else it }.toGameHand().toInt()
                val secondTopCardValue = hand.cards.map { if (it == Card('J')) topCards[1].first else it }.toGameHand().toInt()
                hashMap[hand] = max(topCardValue, secondTopCardValue)
            }
        }

        return hashMap[hand]!!
    }
    override operator fun compareTo(other: GameHand): Int {
        if (wildCard) {
            val thisValue = this.replaceAndFindValue()
            val otherValue = other.replaceAndFindValue()
            return if (thisValue != otherValue)
                thisValue.compareTo(otherValue)
            else {
                this.cards.mapIndexed { idx, card -> card.compareTo(other.cards[idx])  }.find { it != 0 } ?: 0
            }
        }

        val thisHand = this.toInt()
        val otherHand = other.toInt()
        if (thisHand != otherHand)
            return thisHand.compareTo(otherHand)

        return this.cards.mapIndexed { idx, card -> card.compareTo(other.cards[idx])  }.find { it != 0 } ?: 0
    }
    public override fun clone(): GameHand {
        return GameHand(cards.map { it })
    }
}
fun List<Card>.toGameHand(wildCard: Boolean = false) = GameHand(this, wildCard)
fun parseLine(line: String, wildCard: Boolean = false): Pair<GameHand, Int> {
    return Pair(
        line.substringBefore(' ').toCharArray().map { Card(it).setWildCard(wildCard) }.toGameHand(wildCard),
        line.substringAfter(' ').toInt()
    )
}
fun main() {
    fun part1(input: List<String>): Int {
        return input
            .map { parseLine(it) }
            .sortedBy { it.first }
            .mapIndexed { index, (_, second) -> (index + 1) * second  }
            .sum()
    }

    fun part2(input: List<String>): Int {
        return input
            .map { parseLine(it, true) }
            .sortedBy { it.first }
            .mapIndexed { index, (_, second) -> (index + 1) * second  }
            .sum()
    }

    // test if implementation meets criteria from the description, like:
    val testInput = readInput("Day07_test")
    val input = readInput("Day07")
    check(part1(testInput) == 6440)
    part1(input).println()

    check(part2(testInput) == 5905)
    part2(input).println()
}