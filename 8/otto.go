package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const SIZE int = 99

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

	var trees [SIZE][SIZE]int

	y := 0
	for scanner.Scan() {
		var line = scanner.Text()
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		//fmt.Printf("line %s\n", line)

		for x := 0; x < SIZE; x++ {
			trees[x][y], _ = strconv.Atoi(string(line[x]))
		}
		y++
	}

	for y := 0; y < SIZE; y++ {
		for x := 0; x < SIZE; x++ {
			fmt.Printf("%d", trees[x][y])
		}
		fmt.Printf("\n")
	}

	visible := SIZE*2 + ((SIZE - 2) * 2)
	fmt.Printf("Border visible %d\n", visible)

	maxScenic := 0
	for y := 1; y < SIZE-1; y++ {
		for x := 1; x < SIZE-1; x++ {
			isVisible, scenic := scan(trees, x, y)
			if isVisible {
				visible++
			}
			if scenic > maxScenic {
				maxScenic = scenic
			}
		}
	}

	fmt.Printf("Visible %d - max scenic score %d\n", visible, maxScenic)
}

func scan(trees [SIZE][SIZE]int, x int, y int) (bool, int) {
	h := trees[x][y]
	visible := false
	var sl, sr, su, sd int = 0, 0, 0, 0

	// left
	max := 0
	sdone := false
	for ix := x - 1; ix >= 0; ix-- {
		if trees[ix][y] > max {
			max = trees[ix][y]
		}
		if !sdone {
			sl++
			if h <= trees[ix][y] {
				sdone = true
			}
		}
	}
	if h > max {
		visible = true
	}

	// right
	max = 0
	sdone = false
	for ix := x + 1; ix < SIZE; ix++ {
		if trees[ix][y] > max {
			max = trees[ix][y]
		}
		if !sdone {
			sr++
			if h <= trees[ix][y] {
				sdone = true
			}
		}
	}
	if h > max {
		visible = true
	}

	// up
	max = 0
	sdone = false
	for iy := y - 1; iy >= 0; iy-- {
		if trees[x][iy] > max {
			max = trees[x][iy]
		}
		if !sdone {
			su++
			if h <= trees[x][iy] {
				sdone = true
			}
		}
	}
	if h > max {
		visible = true
	}

	// down
	max = 0
	sdone = false
	for iy := y + 1; iy < SIZE; iy++ {
		if trees[x][iy] > max {
			max = trees[x][iy]
		}
		if !sdone {
			sd++
			if h <= trees[x][iy] {
				sdone = true
			}
		}
	}
	if h > max {
		visible = true
	}

	scenicIdx := sl * sr * su * sd
	fmt.Printf("[%d,%d]scenic %d %d %d %d = %d\n", x, y, sl, sr, su, sd, scenicIdx)
	return visible, scenicIdx
}
