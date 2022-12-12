package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

	visited := make(map[string]int)
	var hx, hy int = 0, 0
	var lhx, lhy int = hx, hy
	var tx, ty int = hx, hy

	visited[key(tx, ty)] = 1

	for scanner.Scan() {
		var line = scanner.Text()
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		move := strings.Split(line, " ")
		dir := move[0]
		step, _ := strconv.Atoi(string(move[1]))
		for i := 0; i < step; i++ {
			fmt.Printf("h=%d,%d t=%d,%d ", hx, hy, tx, ty)
			lhx = hx
			lhy = hy
			// move head
			switch dir {
			case "U":
				hy++
			case "D":
				hy--
			case "L":
				hx--
			case "R":
				hx++
			}

			ds, ss := dist(hx, hy, tx, ty)
			fmt.Printf("dist ds=%d ss=%d ", ds, ss)

			// move tail
			if ds+ss > 1 {
				tx = lhx
				ty = lhy
			}

			fmt.Printf("step %s %d[%d] h=%d,%d t=%d,%d \n", dir, step, i, hx, hy, tx, ty)
			visited[key(tx, ty)]++
		}
	}

	fmt.Printf("visited %d\n", len(visited))

}

func printMap(m map[string]int, hx, hy, tx, ty int) {

	for y := 0; y < 40; y++ {
		for x := 0; x < 40; x++ {
			if tx == x && ty == y {
				fmt.Printf("T")
			} else if hx == x && hy == y {
				fmt.Printf("H")
			}
			fmt.Printf(".")

		}
		fmt.Printf("\n")
	}

}

func key(x int, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func dist(hx int, hy int, tx int, ty int) (int, int) {
	dx := abs(tx - hx)
	dy := abs(ty - hy)

	min := min(dx, dy)
	max := max(dx, dy)

	ds := min
	ss := max - min

	return ds, ss
}

func abs(n int) int {
	if n < 0 {
		return -n
	} else {
		return n
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func stepDir(d int) int {
	if d > 0 {
		return -1
	} else {
		return 1
	}
}
