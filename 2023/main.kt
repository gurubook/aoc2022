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

        println(path)

        var line: String?
        while (true) {
            line = bufferedReader.readLine() ?: break
            if (line.isNotBlank()) {
                // println(line)
            }
        }

        bufferedReader.close()
        fileReader.close()

        // Q1
        println("------------------------------- Q1 -------------------------------")
        var tot1 = 0
        println("tot1 = $tot1")

        // Q2
        println("------------------------------- Q2 -------------------------------")
        var tot2 = 0
        println("tot2 = $tot2")

    } catch (e: IOException) {
        e.printStackTrace()
    }
}
