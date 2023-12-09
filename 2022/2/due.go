package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var game map[string]int
var cardPoint map[string]int
var strat map[string]string

func load() {
	game = map[string]int{
		"A X": 3,
		"A Y": 6,
		"A Z": 0,

		"B X": 0,
		"B Y": 3,
		"B Z": 6,

		"C X": 6,
		"C Y": 0,
		"C Z": 3,
	}

	strat = map[string]string{
		"A X": "Z",
		"A Y": "X",
		"A Z": "Y",

		"B X": "X",
		"B Y": "Y",
		"B Z": "Z",

		"C X": "Y",
		"C Y": "Z",
		"C Z": "X",
	}

	cardPoint = map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

}

func main() {
	load()
	// open file
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	var result = 0
	var result2 = 0
	for scanner.Scan() {
		var hand = scanner.Text()
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s", hand)
		// part 1
		result = result + game[hand] + cardPoint[strings.Split(hand, " ")[1]]

		// part 2
		var myplay = strings.Split(hand, " ")[0] + " " + strat[hand]
		fmt.Printf(" - %s\n", myplay)

		result2 = result2 + game[myplay] + cardPoint[strings.Split(myplay, " ")[1]]
	}
	fmt.Printf("\n")
	fmt.Printf("result %d\n", result)
	fmt.Printf("result2 %d\n", result2)

}
