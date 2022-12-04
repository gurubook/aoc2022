package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

type Range struct {
	Start int
	End   int
}

// A slice of ranges that implements sort.Interface
type Ranges []Range

func (r Ranges) Len() int {
	return len(r)
}

func (r Ranges) Less(i, j int) bool {
	return r[i].Start < r[j].Start
}

func (r Ranges) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func fullyContained(a, b Range) bool {
	return a.Start >= b.Start && a.End <= b.End
}

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

	var ranges Ranges
	for scanner.Scan() {
		var line = scanner.Text()
		var a, b Range
		fmt.Sscanf(line, "%d-%d,%d-%d", &a.Start, &a.End, &b.Start, &b.End)
		ranges = append(ranges, a, b)
	}

	// Sort the ranges by starting value
	sort.Sort(ranges)

	// Count the number of pairs that are fully contained
	var count int
	for i := 0; i < len(ranges)-1; i++ {
		if fullyContained(ranges[i], ranges[i+1]) {
			count++
		}
	}

	fmt.Println(count)
}
