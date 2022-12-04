package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Open the input file
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read the input and parse the Calories for each Elf
	scanner := bufio.NewScanner(file)
	var elves []int
	var total int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			// If we encounter a blank line, start a new Elf
			elves = append(elves, total)
			total = 0
		} else {
			// Parse the Calories for each food item
			calories, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			total += calories
		}
	}

	// Add the last Elf
	elves = append(elves, total)

	// Sort the Elves by the number of Calories
