import java.io.File
import java.io.FileReader
import java.io.BufferedReader
import java.io.IOException


fun main(args: Array<String>) {
    println(args.contentToString())
    readFileByLine(args[0]);
}
 
val cardValueOrder = arrayListOf<Char>('A','K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2')
val cardValueOrderWithJoker = arrayListOf<Char>('A','K', 'Q', 'T', '9', '8', '7', '6', '5', '4', '3', '2', 'J')

val cardCountComparator = CardCountComparator()
val evaluationComparator = EvaluationComparator()
val cardToCardComparator = CardToCardComparator()
val evaluationComparatorWithJoker = EvaluationComparatorWithJoker()
val cardToCardComparatorWithJoker = CardToCardComparatorWithJoker()


data class Count(val card: Char , val qty: Int)
data class Hand(val cards: String, val point: Long, val counts:List<Count>, val countsWithJoker:List<Count>)
data class EvaluatedHand(val hand: Hand, val evaluation: Int, val evaluationWithJoker: Int )

fun readFileByLine(filePath: String) {
    try {
        val file = File(filePath)
        val fileReader = FileReader(file)
        val bufferedReader = BufferedReader(fileReader)

        val hands = arrayListOf<EvaluatedHand>()

        var line: String?
        while (true) {
            line = bufferedReader.readLine() ?: break
            if (line.isNotBlank()) {
                println(line)
                val s = line.split(" ")
            
                val cards = s[0]
                val point = s[1].toLong()

                var counts = arrayListOf<Count>()
                val distinct = cards.toCharArray().distinct()
                for (c  in distinct) {
                    counts.add(Count(c,  cards.count { x -> x == c  }))
                }
                val sortedCount = counts.sortedWith(cardCountComparator)

                // Joker subst
                var countsWithJoker = arrayListOf<Count>()
                if (cards.contains('J')) {
                    val sc = if (sortedCount[0].card != 'J' || sortedCount.size == 1) {sortedCount[0].card} else {sortedCount[1].card}

                    val cardsWithJoker = cards.replace('J', sc)
                    val distinctWithJoker = cardsWithJoker.toCharArray().distinct()
                    for (c  in distinctWithJoker) {
                        countsWithJoker.add(Count(c,  cardsWithJoker.count { x -> x == c  }))
                    }    
                } else {
                    countsWithJoker = counts
                }
                val sortedCountWithJoker = countsWithJoker.sortedWith(cardCountComparator)

                val hand = Hand(cards, point, sortedCount, sortedCountWithJoker)

                val res = evaluateHand(hand.counts)
                var resWithJoker = evaluateHand(hand.countsWithJoker)

                println(" $cards -> $res -> $resWithJoker")

                val evaluatedHand = EvaluatedHand(hand, res, resWithJoker)
                hands.add(evaluatedHand)
            }
        }

        bufferedReader.close()
        fileReader.close()

        // Q1
        println("------------------------------- Q1 -------------------------------")
        val rankedHands =  hands.sortedWith(evaluationComparator)
        //println("Hands : $rankedHands")
        val tot1 = rankedHands
            .mapIndexed { i, r -> 
                val p = r.hand.point * (i + 1)
                println("$i(${r.hand.cards}) -> $p : ${r.hand.counts}")
                p
             }
            .sum()
        println("tot1 = $tot1")

        // Q2
        println("------------------------------- Q2 -------------------------------")
        val rankedHandsWithJoker =  hands.sortedWith(evaluationComparatorWithJoker)
        //println("HandsWithJoker : $rankedHandsWithJoker")
        val tot2 = rankedHandsWithJoker
            .mapIndexed { i, r -> 
                val p = r.hand.point * (i + 1)
                println("$i(${r.hand.cards}) -> $p : ${r.hand.countsWithJoker}")
                p
             }
            .sum()
        println("tot2 = $tot2")

    } catch (e: IOException) {
        e.printStackTrace()
    }
}


// 0 Five of a kind .. 6 High card
fun evaluateHand(counts: List<Count>): Int {
    println("counts="  + counts.size)
    when (counts.size) {
        1 -> {
            // Five of a kind
            return 0
        }
        2 -> {
            // Four of a kind or Full house
            if (counts[0].qty == 4) {
                return 1
            } else {
                return 2
            }
        }
        3 -> {
            // Three of a kind or Two pair
            if (counts[0].qty == 3) {
                return 3
            } else {
                return 4
            }
        }
        4 ->  {
            return 5
        }
        else -> {
            return 6
        }
    }
}

class CardCountComparator : Comparator <Count> { 
    override fun compare (c0: Count, c1: Count) : Int {  
        val p0 = c0.qty
        val p1 = c1.qty
        if (p0 > p1) { 
            return -1
        } 
        if (p0 == p1) { 
            return 0
        }  
        return 1
    }    
}

class EvaluationComparator : Comparator <EvaluatedHand> { 
    override fun compare (c0: EvaluatedHand, c1: EvaluatedHand) : Int {  
        val p0 = c0.evaluation
        val p1 = c1.evaluation
        if (p0 > p1) { 
            return -1
        } 
        if (p0 == p1) { 
            // card to card compare 
            return cardToCardComparator.compare(c0.hand.cards, c1.hand.cards)
        }  
        return 1
    }    
}


class EvaluationComparatorWithJoker : Comparator <EvaluatedHand> { 
    override fun compare (c0: EvaluatedHand, c1: EvaluatedHand) : Int {  
        val p0 = c0.evaluationWithJoker
        val p1 = c1.evaluationWithJoker
        if (p0 > p1) { 
            return -1
        } 
        if (p0 == p1) { 
            // card to card compare 
            return cardToCardComparatorWithJoker.compare(c0.hand.cards, c1.hand.cards)
        }  
        return 1
    }    
}

class CardToCardComparator : Comparator <String> {
  override fun compare (c0: String, c1: String) : Int { 
    val a0 = c0.toCharArray()
    val a1 = c1.toCharArray()
    for (i in 0..5) {
        val v0 = cardValueOrder.indexOf(a0[i])
        val v1 = cardValueOrder.indexOf(a1[i])
        if (v0 != v1) {
            return v1.compareTo(v0)
        }
    }
    return 0
  }
}

class CardToCardComparatorWithJoker : Comparator <String> {
  override fun compare (c0: String, c1: String) : Int { 
    val a0 = c0.toCharArray()
    val a1 = c1.toCharArray()
    for (i in 0..5) {
        val v0 = cardValueOrderWithJoker.indexOf(a0[i])
        val v1 = cardValueOrderWithJoker.indexOf(a1[i])
        if (v0 != v1) {
            return v1.compareTo(v0)
        }
    }
    return 0
  }
}

