package main

import (
	"bufio"
	"container/ring"
	"fmt"
	"log"
	"os"
)

func main() {
	// open file
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	markSize := 14 // 4 for SOP

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)
	var v int8 = 0
	for scanner.Scan() {
		r := ring.New(markSize)

		var line = scanner.Text()
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		for count := 0; count < len(line); count++ {
			r.Value = string(line[count])
			r = r.Next()

			mark := make(map[string]int8)

			r.Do(func(s any) {
				if s != nil {
					//fmt.Printf("%s", s)
					mark[s.(string)] = v
				}
			})
			printMark(mark)
			if len(mark) == markSize {
				fmt.Printf(" - SOP at %d\n", count+1)
				break
			} else {
				fmt.Printf(" - mark size %d at %d\n", len(mark), count-3)
			}

		}
	}
}

func printMark(m map[string]int8) {
	for k, _ := range m {
		fmt.Printf("%s", k)
	}
}
