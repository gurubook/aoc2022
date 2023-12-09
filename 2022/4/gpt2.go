package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
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

func overlap(a, b Range) bool {
	return a.End >= b.Start && a.Start <= b.End
}

func main() {
	// Open the input file
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Parse the input into a slice of ranges
	scanner := bufio.NewScanner(file)
	var ranges Ranges
	for scanner.Scan() {
		s := scanner.Text()
		parts := strings.Split(s, ",")
		var a, b Range
		fmt.Sscanf(parts[0], "%d-%d", &a.Start, &a.End)
		fmt.Sscanf(parts[1], "%d-%d", &b.Start, &b.End)
		ranges = append(ranges, a, b)
	}

	// Sort the ranges by starting value
	sort.Sort(ranges)

	// Count the number of pairs that overlap
	var count int
	for i := 0; i < len(ranges)-1; i++ {
		if overlap(ranges[i], ranges[i+1]) {
			count++
		}
	}

	fmt.Println(count)
}
