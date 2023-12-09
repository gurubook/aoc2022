package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const CYCLESTART int = 20
const CYCLECOUNT int = 40
const CRTCOLS int = 40
const CRTTLINES int = 6

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

	cycle := 0
	ss := 0
	x := 1

	var crt [CRTTLINES][CRTCOLS]string

	for scanner.Scan() {
		var line = scanner.Text()
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		code := strings.Split(line, " ")
		instr := code[0]
		switch instr {
		case "noop":
			cycle++
			ss += checkCycle(cycle, x)
			updateCRT(cycle, x, &crt)
		case "addx":
			param, _ := strconv.Atoi(code[1])
			for i := 0; i < 2; i++ {
				cycle++
				ss += checkCycle(cycle, x)
				updateCRT(cycle, x, &crt)
			}
			x = x + param
		}

	}

	fmt.Printf("\nSS=%d\n", ss)

	for cy := 0; cy < CRTTLINES; cy++ {
		for cx := 0; cx < CRTCOLS; cx++ {
			fmt.Printf("%s", crt[cy][cx])
		}
		fmt.Printf("\n")
	}

}

func checkCycle(cycle int, x int) int {
	if cycle == CYCLESTART || ((cycle-CYCLESTART)%CYCLECOUNT == 0) {
		fmt.Printf("%d ", cycle)
		return cycle * x
	}
	return 0
}

func updateCRT(cycle int, x int, crt *[CRTTLINES][CRTCOLS]string) {
	by := (cycle - 1) / CRTCOLS
	bx := (cycle - 1) - (by * CRTCOLS)
	if bx >= x-1 && bx <= x+1 {
		crt[by][bx] = "#"
	} else {
		crt[by][bx] = "."
	}
}
