import java.io.File
import java.io.FileReader
import java.io.BufferedReader
import java.io.IOException
import kotlin.math.abs

var e2 = 0L

fun main(args: Array<String>) {
    println(args.contentToString())
    e2 = args[1].toLong()
    println("expansion factor 2 : $e2")

    readFileByLine(args[0]);
}
 
class Galaxy(var x: Long, var y:Long) {
    override fun toString(): String {
        return "($x,$y)"
    }
}

fun readFileByLine(filePath: String) {
    try {
        val file = File(filePath)
        val fileReader = FileReader(file)
        val bufferedReader = BufferedReader(fileReader)

        val universe = arrayListOf<ArrayList<Char>>()

        val galaxies = arrayListOf<Galaxy>()
        val galaxies1 = arrayListOf<Galaxy>()
        val galaxies2 = arrayListOf<Galaxy>()

        val doubleRow = arrayListOf<Int>()
        val doubleCol = arrayListOf<Int>()
        var i = 0

        var line: String?
        while (true) {
            line = bufferedReader.readLine() ?: break
            if (line.isNotBlank()) {
                // println(line)
                universe.add(line.toCharArray().toCollection(ArrayList()))
                // check empty rows
                if (line == ".".repeat(line.length)) {
                    doubleRow.add(i)
                }
                i++
            }
        }

        bufferedReader.close()
        fileReader.close()

        // check empty cols
        for (x in 0..universe[0].size -1) {
            var emptyCol = true
            for (y in 0..universe.size -1) {
                if (universe[y][x] != '.') {
                    emptyCol = false
                    break
                }
            }
            if (emptyCol) {
                doubleCol.add(x)
            }
        }

        println("doubleRow = $doubleRow")
        println("doubleCol = $doubleCol")

        printUniverse(universe)

        // locate galaxies
        for (y in 0..universe.size -1) {
            for (x in 0..universe[y].size -1) {
                if (universe[y][x] == '#') {
                    galaxies.add(Galaxy(x.toLong(), y.toLong()))
                }
            }
        }

        // expand universes
        for (gi in 0..galaxies.size -1) {
            var g = galaxies[gi]

            galaxies1.add(Galaxy(g.x, g.y))
            galaxies2.add(Galaxy(g.x, g.y))

            var dx1 = doubleCol.filter{ it < g.x }.count() * 1
            var dx2 = doubleCol.filter{ it < g.x }.count() * (e2 - 1)
            
            var dy1 = doubleRow.filter{ it < g.y }.count() * 1
            var dy2 = doubleRow.filter{ it < g.y }.count() * (e2 - 1)
            
            galaxies1[gi].x = g.x + dx1 
            galaxies2[gi].x = g.x + dx2 

            galaxies1[gi].y = g.y + dy1 
            galaxies2[gi].y = g.y + dy2
        }

        println("galaxies: $galaxies")
        println("galaxies1: $galaxies1")
        println("galaxies2: $galaxies2")
        

        // Q1
        println("------------------------------- Q1 -------------------------------")
        var tot1 = calculateDistances(galaxies1)

        println("tot1 = $tot1")

        // Q2
        println("------------------------------- Q2 -------------------------------")
        var tot2 = calculateDistances(galaxies2)
        println("tot2 = $tot2")

    } catch (e: IOException) {
        e.printStackTrace()
    }
}

fun calculateDistances(galaxies : List<Galaxy>) : Long {
    var tot1 = 0L
    for (i in 0..galaxies.size-1) {
        val f = galaxies[i]
        for (ti in (i+1)..galaxies.size-1) {
            val t = galaxies[ti]
            val md = abs(t.x - f.x) + abs(t.y - f.y)
            println("$i-$ti = $md")
            tot1 = tot1 + md
        }
    }
    return tot1
}

fun printUniverse(universe : ArrayList<ArrayList<Char>>) {
    for (l in universe) {
        for (c in l) {
            print(c)
        }
        println()
    }
}