import java.io.File
import java.io.FileReader
import java.io.BufferedReader
import java.io.IOException

fun main(args: Array<String>) {
    println(args.contentToString())
    readFileByLine(args[0]);
}

data class Race(val time:Long, val dist:Long)

fun readFileByLine(filePath: String) {
    try {
        val file = File(filePath)
        val fileReader = FileReader(file)
        val bufferedReader = BufferedReader(fileReader)

        val races = arrayListOf<Race>()

        var line1: String?
        var line2: String?
        
        line1 = bufferedReader.readLine()
        line2 = bufferedReader.readLine()

        val times = line1.substringAfter("Time:")
                        .split(' ')
                        .filter { x -> x.isNotBlank() }
                        .map{x -> x.trim().toLong()}
                        .toList()
        
        val distances = line2.substringAfter("Distance:")
                        .split(' ')
                        .filter { x -> x.isNotBlank() }
                        .map{x -> x.trim().toLong()}
                        .toList()


        val times2 =  line1.substringAfter("Time:").replace(" ", "")
        val distances2 =  line2.substringAfter("Distance:").replace(" ", "")
        

        bufferedReader.close()
        fileReader.close()

        for (i in 0..times.size-1) {
            val race = Race(times[i], distances[i])
            races.add(race)
        }

        val race2 = Race(times2.toLong(), distances2.toLong())
        // q1
        println("------------------------------- Q1 -------------------------------")

        var tot = 1;

        for (r in races) {
            println("Race duration=$r.time distance=$r.dist") 
            var wins = 0
            for (d in 1..r.time) {
                val speed = d
                val distance = speed * (r.time - d)
                print("hold=$d, speed=$speed, distance=$distance ")
                if (distance > r.dist) {
                    println (" WON ")
                    wins++ 
                } else {
                    println (" loose ")
                }
            }
            println("wins=$wins")
            tot = tot * wins
            println("-------------------------------")
        }
        println("tot=$tot")
        
        // q2
        println("------------------------------- Q2 -------------------------------")
        println("Race: $race2")

            var wins = 0
            for (d in 1..race2.time) {
                val speed = d
                val distance = speed * (race2.time - d)
                // print("hold=$d, speed=$speed, distance=$distance ")
                if (distance > race2.dist) {
                    // println (" WON ")
                    wins++ 
                } else {
                    // println (" loose ")
                }
            }
            println("wins=$wins")

    } catch (e: IOException) {
        e.printStackTrace()
    }
}

