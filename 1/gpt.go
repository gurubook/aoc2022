package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	var most int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			// If we encounter a blank line, start a new Elf
			elves = append(elves, total)
			if total > most {
				most = total
			}
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
	if total > most {
		most = total
	}

	// Find the index of the Elf carrying the most Calories
	var index int
	for i, elf := range elves {
		if elf == most {
			index = i
			break
		}
	}

	fmt.Println("Elf", index+1, "is carrying the most Calories with a total of", most)
}
