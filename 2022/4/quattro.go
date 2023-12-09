package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	included := 0
	overlapped := 0
	for scanner.Scan() {
		var line = scanner.Text()
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		sections := strings.Split(line, ",")

		sec1 := strings.Split(sections[0], "-")
		sec2 := strings.Split(sections[1], "-")

		s1, _ := strconv.Atoi(sec1[0])
		e1, _ := strconv.Atoi(sec1[1])

		s2, _ := strconv.Atoi(sec2[0])
		e2, _ := strconv.Atoi(sec2[1])

		if s1 <= e2 && e1 >= s2 {
			fmt.Printf("%s overlapped ", line)
			overlapped++
			if s1 == s2 || e1 == e2 {
				fmt.Printf("\t %s included", line)
				included++
			} else {
				if s1 <= s2 {
					if e1 >= e2 {
						fmt.Printf("\t %s included", line)
						included++
					}
				} else {
					if e1 <= e2 {
						fmt.Printf("\t %s included", line)
						included++
					}
				}
			}
			fmt.Printf("\n")
		}
	}

	fmt.Printf("Overlapped count %d included %d\n", overlapped, included)
}
