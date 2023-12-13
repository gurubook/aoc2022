import java.io.File
import java.io.FileReader
import java.io.BufferedReader
import java.io.IOException

fun main(args: Array<String>) {
    println(args.contentToString())
    readFileByLine(args[0]);
}
 
fun readFileByLine(filePath: String) {
    try {
        val file = File(filePath)
        val fileReader = FileReader(file)
        val bufferedReader = BufferedReader(fileReader)

        var line: String?
        var tot1 = 0L
        var tot2 = 0L

        while (true) {
            line = bufferedReader.readLine() ?: break
            if (line.isNotBlank()) {
                // println(line)
                val seq = arrayListOf<List<Long>>()
                seq.add( line.split(' ').map { x -> x.toLong() }.toList() )
                var i = 0
                while(true) {
                    val cur = seq[i]
                    val next = arrayListOf<Long>()
                    for (j in 0..cur.size-2) {
                        next.add(cur[j+1]-cur[j])
                    } 
                    if (next.filter { it != 0L }.toList().size == 0) {
                        break
                    }
                    seq.add(next)
                    i++
                }
                println("seq : $seq")
                
                var add = 0L
                var add2 = 0L
                
                for (r in seq.reversed()) {
                    add = add + r[r.size-1]
                    add2 = r[0] - add2
                }
                tot1 = tot1 + add
                tot2 = tot2 + add2
                // Q1
                println("------------------------------- Q1 -------------------------------")
                
                println("tot1 = $tot1")

                // Q2
                println("------------------------------- Q2 -------------------------------")
                
                println("tot2 = $tot2")
            }
        }

        bufferedReader.close()
        fileReader.close()


    } catch (e: IOException) {
        e.printStackTrace()
    }
}
