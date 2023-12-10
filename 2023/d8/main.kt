import java.io.File
import java.io.FileReader
import java.io.BufferedReader
import java.io.IOException

fun main(args: Array<String>) {
    println(args.contentToString())
    readFileByLine(args[0]);
}
 
class Node(val key:String ) {
    var left:Node? = null
    var right:Node? = null
}

val lines = arrayListOf<String>();

val nodes = arrayListOf<Node>();

fun readFileByLine(filePath: String) {
    try {
        val file = File(filePath)
        val fileReader = FileReader(file)
        val bufferedReader = BufferedReader(fileReader)

        // read path
        var path = bufferedReader.readLine()
        println(path)

        bufferedReader.readLine()

        var line: String?
        while (true) {
            line = bufferedReader.readLine() ?: break
            if (line.isNotBlank()) {
                // println(line)
                lines.add(line)
            }
        }

        bufferedReader.close()
        fileReader.close()

        // create nodes
        for (l in lines) {
            val s = l.split(' ')
            nodes.add(Node(s[0]))
        }

        // bind nodes
        for (l in lines) {
            val s = l.split(' ')
            val n = findNode(nodes, s[0])

            n!!.left = findNode(nodes, s[2].substring(1..3))!!
            n.right = findNode(nodes, s[3].substring(0..2))!!
            
            println("${n.key} = (${n.left?.key} ${n.right?.key})")
        }

        // Q1
        println("------------------------------- Q1 -------------------------------")
        var idx = 0
        var tot1 = 0
        var node = findNode(nodes, "AAA")

        if (node != null) {
            while(true) {
                val p = path[idx]
                print("$p ")
                print("${node!!.key} = (${node.left?.key} ${node.right?.key}) -> ")
                // node = when (p) {
                //     'L' -> node!!.left
                //     'R' -> node!!.right
                //     else -> null
                // }

                node = selectNode(node, p)
                println("${node!!.key} = (${node.left?.key} ${node.right?.key})")

                idx++
                tot1++

                if (idx >= path.length) {
                    idx = 0
                }
                if (node.key == "ZZZ") {
                    break;
                }
            }

            println("tot1 = $tot1")
        }
        // Q2
        println("------------------------------- Q2 -------------------------------")
        
        var starts = findNodeEndingWith(nodes, 'A')
        var currents = arrayListOf<Node>()
        val startsCount = starts.size

        idx = 0

        var tot2 = 0L
 
        for (n in starts) {
            println("${n.key} = (${n.left?.key} ${n.right?.key})")
        }

        var endings = arrayListOf<Long>()

        while(true) {
            val p = path[idx]
            //print("$p")

            for (n in starts) {
                val cn = selectNode(n, p)!!
                if (cn.key.endsWith('Z')) {
                    endings.add((tot2+1).toLong());
                } else {
                    currents.add(cn)
                }
                // println("${n.key} = (${n.left?.key} ${n.right?.key})")
            }

            idx++
            tot2++

            if (idx >= path.length) {
                idx = 0
            }

            //val endsCount = currents.filter { it.key.endsWith('Z') }.count()
            // if (endsCount == currents.size) {
            //     break;
            // }

            if (startsCount == endings.size) {
                println("Endings $endings")
                break
            }

            starts = currents.toList()
            currents.clear()
        }       

        tot2 = findLCMOfListOfNumbers(endings)
        println("tot2 = $tot2")

    } catch (e: IOException) {
        e.printStackTrace()
    }
}

fun selectNode(node: Node, p: Char): Node? {
   return when (p) {
        'L' -> node.left
        'R' -> node.right
        else -> null
    }
}

fun findNode(nodes : List<Node>, key: String) : Node? {
    return nodes.find { 
        it.key == key
    }
}

fun findNodeEndingWith(nodes : List<Node>, key: Char) : List<Node> {
    return nodes.filter { 
        it.key.endsWith(key)
    }
}

fun findLCM(a: Long, b: Long): Long {
    val larger = if (a > b) a else b
    val maxLcm = a * b
    var lcm = larger
    while (lcm <= maxLcm) {
        if (lcm % a == 0L && lcm % b == 0L) {
            return lcm
        }
        lcm += larger
    }
    return maxLcm
}

fun findLCMOfListOfNumbers(numbers: List<Long>): Long {
    var result = numbers[0]
    for (i in 1 until numbers.size) {
        result = findLCM(result, numbers[i])
    }
    return result
}

