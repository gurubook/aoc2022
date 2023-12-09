package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Open the input file
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read the input and calculate the sum of the priorities
	scanner := bufio.NewScanner(file)
	var sum int
	for scanner.Scan() {
		line := scanner.Text()

		// Split the line into the two compartments
		n := len(line)
		first := line[:n/2]
		second := line[n/2:]

		// Find the common item types in the two compartments
		var common string
		for _, c := range first {
			if strings.ContainsRune(second, c) {
				common += string(c)
			}
		}

		// Calculate the priority of the common item type
		if len(common) > 0 {
			if common[0] >= 'a' && common[0] <= 'z' {
				sum += int(common[0] - 'a' + 1)
			} else {
				sum += int(common[0] - 'A' + 27)
			}
		}
	}

	// Print the result
	fmt.Println(sum)
}
