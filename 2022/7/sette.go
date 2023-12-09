package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// open file
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	fs := make(map[string]int)
	cwd := "/"

	for scanner.Scan() {

		var line = scanner.Text()
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		//fmt.Printf("line %s\n", line)

		if string(line[0]) == "$" {
			cmd := line[2:]
			fmt.Printf("cmd %s\n", cmd)

			if strings.HasPrefix(cmd, "cd") {
				p := cmd[3:]
				cwd = path.Join(cwd, p)
				fs[cwd] = 0

			}
			// else if strings.HasPrefix(cmd, "ls") {

			// }
		} else if strings.HasPrefix(line, "dir") {
			// ls output dir
			dir := line[4:]
			fs[path.Join(cwd, dir)] = 0
		} else {
			// ls output file
			f := strings.Split(line, " ")
			size, _ := strconv.Atoi(f[0])
			fs[path.Join(cwd, f[1])] = size
		}
	}

	keys := sortKeys(fs)
	dirSizes := make(map[string]int)
	for _, k := range keys {
		// is dir ?
		if fs[k] == 0 {
			dirSizes[k] = 0
		} else {
			for rk, _ := range dirSizes {
				if strings.HasPrefix(k, rk) {
					fmt.Printf("dir %s add file %s size %d\n", rk, k, fs[k])
					dirSizes[rk] = dirSizes[rk] + fs[k]
				}
			}
		}
	}

	//printFs(fs)

	printSumFsAtMost(dirSizes, 100000)

	printSmallestDirToDeleteToFree(dirSizes, 70000000, 30000000)

}

func sortKeys(fs map[string]int) []string {
	keys := make([]string, 0, len(fs))
	for k := range fs {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func printFs(fs map[string]int) {
	keys := sortKeys(fs)

	for _, k := range keys {
		fmt.Println(k, fs[k])
	}
}

func printSumFsAtMost(fs map[string]int, atMost int) {
	keys := sortKeys(fs)
	sum := 0
	for _, k := range keys {
		if fs[k] < atMost {
			fmt.Println(k, fs[k])
			sum += fs[k]
		}
	}

	fmt.Printf("total %d\n", sum)
}

func printSmallestDirToDeleteToFree(fs map[string]int, total int, free int) {
	used := fs["/"]
	needed := free - (total - used)

	var target string = ""
	var targetSize int = total

	keys := sortKeys(fs)
	for _, k := range keys {
		if fs[k] > needed && fs[k] < targetSize {
			target = k
			targetSize = fs[k]
		}
	}

	fmt.Printf("delete dir %s sized %d\n", target, targetSize)

}
