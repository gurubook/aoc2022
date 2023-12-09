package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/yourbasic/graph"
	"golang.org/x/exp/slices"
)

type loc struct {
	x     int
	y     int
	h     int
	edges []*loc
	start bool
	end   bool
}

func (l loc) String() string {
	se := "-"
	if l.start {
		se = "S"
	} else if l.end {
		se = "E"
	}
	return fmt.Sprintf("(%d,%d)%d %s", l.x, l.y, l.h, se)
}

type graphNodes []loc

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

	land := make([][]string, 0)
	nodes := make(graphNodes, 0)

	var startNodeIdx int
	var endNodeIdx int

	y := 0
	var cols int
	var rows int

	for scanner.Scan() {
		var line = scanner.Text()
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		land = append(land, make([]string, len(line)))
		for x := 0; x < len(line); x++ {
			s := false
			e := false
			c := string(line[x])
			switch c {
			case "S":
				s = true
			case "E":
				e = true
			}

			// fill map and nodes
			land[y][x] = c
			cl := loc{x: x, y: y, h: height(x, y, land), start: s, end: e}
			nodes = append(nodes, cl)

		}
		y++
	}

	cols = len(land[0])
	rows = len(land)

	g := graph.New(len(nodes))

	for i, n := range nodes {
		if n.start {
			startNodeIdx = i
		}
		if n.end {
			endNodeIdx = i
		}

		// up
		if n.y > 0 && checkHeight(n, nodes[i-cols]) {
			nodes[i].edges = append(nodes[i].edges, &nodes[i-cols])
			g.AddCost(i, i-cols, heightCost(n, nodes[i-cols]))
		}
		// down
		if n.y < rows-1 && checkHeight(n, nodes[i+cols]) {
			nodes[i].edges = append(nodes[i].edges, &nodes[i+cols])
			g.AddCost(i, i+cols, heightCost(n, nodes[i+cols]))
		}

		// left
		if n.x > 0 && checkHeight(n, nodes[i-1]) {
			n.edges = append(n.edges, &nodes[i-1])
			g.AddCost(i, i-1, heightCost(n, nodes[i-1]))
		}
		// right
		if n.x < cols-1 && checkHeight(n, nodes[i+1]) {
			nodes[i].edges = append(nodes[i].edges, &nodes[i+1])
			g.AddCost(i, i+1, heightCost(n, nodes[i+1]))
		}

		// fmt.Printf("%d=%s ", i, nodes[i])
	}

	// printMap(land)
	fmt.Printf("\n\nStart %d,%d end %d,%d\n", nodes[startNodeIdx].x, nodes[startNodeIdx].y, nodes[endNodeIdx].x, nodes[endNodeIdx].y)
	fmt.Printf("StartIdx %d endIdx %d\n\n", startNodeIdx, endNodeIdx)
	// fmt.Printf("Graph %v\n\n", g)

	path, dist := graph.ShortestPath(g, startNodeIdx, endNodeIdx)
	if dist != -1 {
		// printPath(land, path)
	}
	fmt.Println()
	// fmt.Println("path:", path, "length:", dist)
	fmt.Printf("total step %d\n", len(path)-1)

	sop := 99999999999999
	for i, n := range nodes {
		if n.h == 1 {
			path, dist := graph.ShortestPath(g, i, endNodeIdx)
			if dist != -1 && len(path)-1 < sop {
				sop = len(path) - 1
			}
		}
	}
	fmt.Printf("shorted of all paths %d\n", sop)
}

func height(x, y int, land [][]string) int {
	if land[y][x] == "S" {
		return 1 // ground a
	} else if land[y][x] == "E" {
		return 26 // z
	} else {
		return int(land[y][x][0]) - 96 // ascii - 96 (a=1)
	}
}

func checkHeight(n1, n2 loc) bool {
	return abs(n1.h-n2.h) < 2 || n2.h < n1.h
}

func heightCost(n1, n2 loc) int64 {
	return int64(abs(n1.h-n2.h)) + 1
}

func abs(n int) int {
	if n < 0 {
		return -n
	} else {
		return n
	}
}

func printMap(land [][]string) {
	for _, row := range land {
		for _, c := range row {
			fmt.Printf("%s", c)
		}
		fmt.Printf("\n")
	}
}

func printPath(land [][]string, path []int) {
	for y, row := range land {
		for x, c := range row {
			ni := (len(land[0]) * y) + x
			pi := slices.Index(path, ni)
			if pi != -1 {
				fmt.Printf("%2d ", pi)
			} else {
				fmt.Printf(" %s ", c)
			}
		}
		fmt.Printf("\n")
	}
}
