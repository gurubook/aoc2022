import java.io.File
import java.io.FileReader
import java.io.BufferedReader
import java.io.IOException
import kotlin.math.abs

fun main(args: Array<String>) {
    println(args.contentToString())
    readFileByLine(args[0]);
}
 
fun readFileByLine(filePath: String) {
    try {
        val file = File(filePath)
        val fileReader = FileReader(file)
        val bufferedReader = BufferedReader(fileReader)

        val map = arrayListOf<List<Char>>()
        val loop = arrayListOf<ArrayList<Char>>()
        val loopPos = arrayListOf<Pos>()

        var line: String?
       
        var startX = 0
        var startY = -1 

        var ln = 0
        while (true) {
            line = bufferedReader.readLine() ?: break
            if (line.isNotBlank()) {
                // println(line)
                map.add(line.toCharArray().toList())
                loop.add(".".repeat(line.length).toCharArray().toCollection(ArrayList()))
                //loop.add(line.toCharArray().toCollection(ArrayList()))

                if (line.contains('S')) {
                    startX = line.indexOf('S')
                    startY = ln
                } 
                ln++
            }
        }

        bufferedReader.close()
        fileReader.close()

        println("start[$startY, $startX] ${map[startY][startX]}")
        
        val startPos = Pos(startY, startX)
        val startsDirs = startDirs(map, startPos)
        //loopPos.add(startPos)

        println("starts = $startsDirs")

        // Q1
        println("------------------------------- Q1 -------------------------------")
        var tot1 = 0
   
        val currents = arrayListOf<Move>()
        val nexts = arrayListOf<Move>()

        startsDirs.forEach { currents.add(Move(startPos, it)) }

        var cnt = 0
        while (true) {
            println("currents : $currents")
            nexts.clear()   
            for (c in currents) {
                val m = findNext(map, c.pos, c.dir)!!
                nexts.add(m)
                loop[c.pos.y][c.pos.x] = map[c.pos.y][c.pos.x]
                //loop[c.pos.y][c.pos.x] = (cnt % 10).toString().get(0)
                loopPos.add(c.pos)
                cnt++
            }
            currents.clear()
            currents.addAll(nexts)

            tot1++
            if (currents[0].pos.x == currents[1].pos.x && currents[0].pos.y == currents[1].pos.y) {
                loop[currents[0].pos.y][currents[0].pos.x] =  map[currents[0].pos.y][currents[0].pos.x]
                loopPos.add(currents[0].pos)
                break

            }
        }
        println("tot1 = $tot1")

        // Q2
        println("------------------------------- Q2 -------------------------------")
        //

        displayMap(map)
        displayLoop(loop)
        
       
        // for (y in 0..loop.size-1) {
        //     floodFill(loop, y, 0)
        //     floodFill(loop, y, loop[0].size -1)
        // }
        // for (x in 0..loop[0].size-1) {
        //     floodFill(loop, 0, x)
        //     floodFill(loop, loop.size -1, x)
        // }
        // displayLoop(loop)
        // var tot2 = countEnclosed(loop)

        // var tot2 = countEnclosedWithPos(map.size , map[0].size , loopPos, loop, map)

        val area = abs(calculateArea(loopPos))
        val boundaryPoints = loopPos.distinct().count()
        println("area=$area\tboundaries=$boundaryPoints")

        var tot2 = area - boundaryPoints / 2f + 1f
        // displayLoop(loop)

        println("tot2 = $tot2")

    } catch (e: IOException) {
        e.printStackTrace()
    }
}

fun calculateArea(loopPos : List<Pos>): Float {
    var area = 0f
    for (i in 0..loopPos.size-2 step 2) {
        val nextIdx = i+2
        val cp = loopPos[i]
        val np = loopPos[nextIdx]
        println("$i(${cp.x},${cp.y}) : $nextIdx(${np.x},${np.y})")
        area += (cp.x) * (np.y) - (cp.y) * (np.x)
    }

    println("--------")
    val rcp = loopPos[loopPos.size-1] 
    val rnp = loopPos[loopPos.size-2] 
    println("(${rcp.x},${rcp.y}) : (${rnp.x},${rnp.y})")
    area += (rcp.x) * (rnp.y) - (rcp.y) * (rnp.x)
    println("--------")

    for (i in loopPos.size-2 downTo 3 step 2 ) {
        val nextIdx = i-2
        val cp = loopPos[i]
        val np = loopPos[nextIdx]
        println("$i(${cp.x},${cp.y}) : $nextIdx(${np.x},${np.y})")
        area += (cp.x) * (np.y) - (cp.y) * (np.x)
    }
    return area / 2f
 }


data class Pos(val y: Int, val x: Int)
data class Move(val pos: Pos, val dir: Char)

val dirs = listOf('N','W', 'S', 'E')
val coord = mapOf(
    'N' to Pos(-1,  0),
    'S' to Pos( 1,  0),
    'E' to Pos( 0,  1),
    'W' to Pos( 0, -1)
)

val moves = mapOf(
    'N' to mapOf('|' to Move(Pos(-1,  0), 'N'), '7' to Move(Pos(-1, -1), 'W'), 'F' to Move(Pos(-1,  1), 'E')),
    'S' to mapOf('|' to Move(Pos( 1,  0), 'S'), 'L' to Move(Pos( 1,  1), 'E'), 'J' to Move(Pos(-1, -1), 'W')),
    'E' to mapOf('-' to Move(Pos( 0,  1), 'E'), 'J' to Move(Pos(-1,  1), 'N'), '7' to Move(Pos(1,  1), 'S')),
    'W' to mapOf('-' to Move(Pos( 0, -1), 'W'), 'L' to Move(Pos(-1,  -1), 'N'), 'F' to Move(Pos(-1,  1), 'S'))
)


fun startDirs(map: ArrayList<List<Char>>, pos: Pos):List<Char> {
    var starts = arrayListOf<Char>()
    for (d in dirs) {
        val pd = coord[d]
        val x = pos.x+pd!!.x
        val y = pos.y+pd.y
        // check boundary 
        if (y >= 0 && y <= map.size && x >= 0 && x <= map[0].size) {
            val ns = map[y][x]
            val next = moves[d]!![ns]
            if (next != null) {
                starts.add(d)
            }
        } 
    }
    return starts
}

fun findNext(map: ArrayList<List<Char>>, pos: Pos, dir: Char): Move? {
    val dpos = coord[dir]!!
    val nextPos = Pos( pos.y+dpos.y, pos.x+dpos.x)
    val ns = map[nextPos.y][nextPos.x]

    // check boundary 
    // if (y >= 0 && y < map.size && x > 0 && x <= map[0].size) {

    println("dir=$dir $nextPos [$ns]")
    val next = moves[dir]!![ns]
    if (next != null) {
        return Move(nextPos, next.dir)
    }
    // }
    return null 
}


fun countEnclosedWithPos(height : Int, width : Int, loopPos : List<Pos>, loop: ArrayList<ArrayList<Char>>, map: ArrayList<List<Char>> ):Int {
    var count = 0
    for (y in 0..height-2) {
        for (x in 0..width-2) {
            //val edge = loopPos.filter {it.y== y && it.x == x}.any()
            if (map[y][x] == '.') {
                var intersect = false
                for (ix in ((x+1)..map[0].size-1)) {
                    if (map[y][ix] != '.') { 
                        intersect = !intersect
                    }
                }
                if (intersect) {
                    loop[y][x] = 'I'
                    count++
                }
            }
        }
    }
    return count 
}

fun countEnclosed(loop: ArrayList<ArrayList<Char>>): Int{
    var count = 0
    for (r in loop) {
        for (c in r) {
            if (c == '.') {
                count++
            }
        }
    }
    return count
}

fun displayLoop(loop: ArrayList<ArrayList<Char>>) {
    for (r in loop) {
        for (c in r) {
            print(c)
        }
        println()
    }
    println()
}

fun displayMap(map: ArrayList<List<Char>>,) {
    for (r in map) {
        for (c in r) {
            print(c)
        }
        println()
    }
    println()
}


fun floodFill(loop: ArrayList<ArrayList<Char>>, y: Int, x:Int) {
    if (loop[y][x] != '.')  {
        return
    }
    
    loop[y][x] = ' '

    if (x < loop[0].size - 1) {
        floodFill(loop, y, x+1)
    }
    if (x > 0) {
        floodFill(loop, y, x-1)
    }

    if (y < loop.size - 1) {
        floodFill(loop, y+1, x)
    }
    if (y > 0) {
        floodFill(loop, y-1, x)
    }

}


