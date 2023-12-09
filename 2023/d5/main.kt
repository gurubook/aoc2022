import java.io.File
import java.io.FileReader
import java.io.BufferedReader
import java.io.IOException

fun main(args: Array<String>) {
    println(args.contentToString())
    readFileByLine(args[0]);
}

data class Key(val ss:Long, val se:Long)
data class Range(val ds:Long, val de:Long)

fun readFileByLine(filePath: String) {
    try {
        val file = File(filePath)
        val fileReader = FileReader(file)
        val bufferedReader = BufferedReader(fileReader)

        val seeds = arrayListOf<Long>()
        val maps = mutableMapOf<String, MutableMap<Key, Range>>();
        var mapName = ""

        val sequence = listOf("seed-to-soil", "soil-to-fertilizer", "fertilizer-to-water", "water-to-light", 
            "light-to-temperature", "temperature-to-humidity", "humidity-to-location")

        var line: String?
        while (true) {
            line = bufferedReader.readLine() ?: break
            if (line.isNotBlank()) {
                if (line.startsWith("seeds:")) {
                    seeds.addAll(line.substringAfter("seeds:")
                        .split(' ')
                        .filter { x -> x.isNotBlank() }
                        .map{x -> x.trim().toLong()}.toList())
                    println("seeds: $seeds")
                } else if (line.contains("map:")) {
                    mapName = line.trim().split(' ')[0]
                    println("map name: '$mapName'")
                } else {
                    println(line)
                    val mv = line.split(' ')
                    val ed = mv[2].toInt()-1
                    val k = Key(mv[1].toLong(), mv[1].toLong()+ed)
                    val r = Range(mv[0].toLong(), mv[0].toLong()+ed)
                    if (!maps.contains(mapName)) {
                        maps.put(mapName, mutableMapOf<Key, Range>())
                    }
                    maps[mapName]?.put(k, r)
                }
            }
        }

        bufferedReader.close()
        fileReader.close()

        // Q1
        println("------------------------------- Q1 -------------------------------")
        // find location for seeds
        var min1 = Long.MAX_VALUE
        for (s in seeds) {
            var sd = s
            for (m in sequence) {
                print("in : $sd -> ")
                sd = findDest(maps[m]!!, sd)
                println("$m -> $sd")
            }
            if (sd<min1) {
                min1 = sd
            }
        }    

        val q1 = min1
        println("q1 : $q1")
        
        // Q2
        println("------------------------------- Q2 -------------------------------")
        var min2 = Long.MAX_VALUE
        val iter = seeds.iterator()
        while  (iter.hasNext()) {
            val start = iter.next()
            val end = start + iter.next()

            var c = start

            var gap = 1L
            var lastc = -1L
            var lastsd = -1L

            while(c<end) {
                var sd = c
                print("in : $sd -> ")
                for (m in sequence) {
                    // print("in : $sd -> ")
                    sd = findDest(maps[m]!!, sd)
                    // println("$m -> $sd")
                }
                println("$sd")
                if (sd<min2) {
                    min2 = sd
                }
                
                // println("c=$c sd=$sd gap=$gap")
                // println("lastc=$lastc lastsd=$lastsd")
                
                if (lastsd > 0) {
                    if ( sd - lastsd == gap) {
                        //if (end < c + gap + 1) {
                            gap = gap + 1
                        //}
                        lastc = c
                        lastsd = sd
                    } else {
                        println("reset gap")
                        // reset gap
                        c = lastc
                        lastsd = -1
                        gap = 1
                    }
                } else {
                    lastc = c
                    lastsd = sd
                }
                
                println("c=$c sd=$sd gap=$gap")
                println("lastc=$lastc lastsd=$lastsd")
                
                println("-------------------------------")
                c = c + gap
                
            }
        }

        val q2 = min2
        println("q2 : $q2")

    } catch (e: IOException) {
        e.printStackTrace()
    }
}

fun findDest(map: MutableMap<Key, Range>, src: Long) :Long {
    val keys = map.keys.
        filter { e -> src >= e.ss && src<=e.se }
        .toList()

    if (keys.isEmpty()) {
        return src
    } 
    return map[keys[0]]!!.de + (src - keys[0].se)  
}