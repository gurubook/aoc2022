package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func itemDistinct(c1 string) map[string]int {
	r := make(map[string]int)
	for _, c := range c1 {
		k := string(c)
		if r[k] == 0 {
			r[k] = 1
		}
	}
	return r
}

func mergeMap(r1 map[string]int, r2 map[string]int) map[string]int {
	for key, element := range r2 {
		r1[key] = r1[key] + element
	}
	return r1
}

func deleteUnder(r map[string]int, limit int) map[string]int {
	for key, element := range r {
		if element <= limit {
			delete(r, key)
		}
	}
	return r
}

func findDuplicates(c1 string, c2 string) map[string]int {
	r := make(map[string]int)
	for _, c := range c1 {
		k := string(c)
		if r[k] == 0 {
			r[k] = 1
		}
	}
	for _, c := range c2 {
		k := string(c)
		if r[k] == 1 {
			r[k] = 2
		}
	}
	r = deleteUnder(r, 1)
	return r
}

func calcPri(d map[string]int) int {
	pri := 0
	for key, _ := range d {
		as := int(key[0])
		if as > 96 {
			pri += as - 96
		} else {
			pri += as - 64 + 26
		}
	}
	return pri
}

func main() {
	// open file
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	priTotal := 0
	pri2Total := 0

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	var rmerged = make(map[string]int)
	var r [3]map[string]int

	var ci = 1
	for scanner.Scan() {
		var sack = scanner.Text()
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s\n", sack)

		// part one
		slen := len(sack)
		s1 := sack[0 : slen/2]
		s2 := sack[slen/2 : slen]

		dup := findDuplicates(s1, s2)
		fmt.Printf("%s - %s = %s\n", s1, s2, dup)
		priTotal += calcPri(dup)

		// part two
		r[ci-1] = itemDistinct(sack)

		if ci%3 == 0 {
			rmerged = mergeMap(rmerged, r[0])
			rmerged = mergeMap(rmerged, r[1])
			rmerged = mergeMap(rmerged, r[2])
			rmerged = deleteUnder(rmerged, 2)
			fmt.Printf("\t %s\n", rmerged)
			pri2Total += calcPri(rmerged)
			rmerged = make(map[string]int)
			ci = 0
		}
		ci++

	}
	fmt.Printf("part one: %d\n", priTotal)
	fmt.Printf("part two: %d\n", pri2Total)
}
