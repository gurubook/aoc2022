package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
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

	var cmax []int
	var totalcal int = 0

	for scanner.Scan() {
		var line = scanner.Text()
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		if line == "" {
			cmax = append(cmax, totalcal)
			totalcal = 0
		} else {
			fmt.Printf("%s ", line)
			cal, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			totalcal = totalcal + cal
		}
	}

	sort.Slice(cmax, func(i, j int) bool {
		return cmax[i] > cmax[j]
	})

	fmt.Printf("\n")
	var total int = 0
	for i := 0; i < 3; i++ {
		fmt.Printf("cmax %d - %d\n", i, cmax[i])
		total = total + cmax[i]
	}
	fmt.Printf("total %d\n", total)

}
