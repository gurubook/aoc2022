package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const SIZE int = 100
const START int = 0
const ROPLEN int = 10

func main() {
	app := tview.NewApplication()
	table := tview.NewTable().SetBorders(true)

	table.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEscape {
			app.Stop()
		}
		if key == tcell.KeyEnter {
			table.SetSelectable(true, true)
		}
	}).SetSelectedFunc(func(row int, column int) {
		table.GetCell(row, column).SetTextColor(tcell.ColorRed)
		table.SetSelectable(false, false)
	})

	defer app.Stop()

	go nove(app, table)

	app.SetRoot(table, true).SetFocus(table)

	if err := app.Run(); err != nil {
		panic(err)
	}

}

func nove(app *tview.Application, table *tview.Table) {
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
	hx := [ROPLEN]int{START, START, START, START, START, START, START, START, START}
	hy := [ROPLEN]int{START, START, START, START, START, START, START, START, START}

	visited[key(hx[ROPLEN-1], hy[ROPLEN-1])] = 1

	setCell(app, table, hx, hy, visited)

	for scanner.Scan() {
		var line = scanner.Text()
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		move := strings.Split(line, " ")
		dir := move[0]
		step, _ := strconv.Atoi(string(move[1]))

		for i := 0; i < step; i++ {
			// move head
			switch dir {
			case "U":
				hy[0]++
			case "D":
				hy[0]--
			case "L":
				hx[0]--
			case "R":
				hx[0]++
			}

			for c := 1; c < ROPLEN; c++ {
				dx := hx[c] - hx[c-1]
				dy := hy[c] - hy[c-1]

				ds, ss := dist(hx[c], hy[c], hx[c-1], hy[c-1])

				if ds+ss > 1 {
					if ds > 0 {
						hx[c] = hx[c] - (sign(dx))
						hy[c] = hy[c] - (sign(dy))
					}
					if ss > 1 {
						hx[c] = hx[c] - (sign(dx))
						hy[c] = hy[c] - (sign(dy))
					}
					//fmt.Printf("\tstep %s %d[%d] h[%d]=%d,%d \n", dir, step, i, c, hx[c], hy[c])

				}
				//logCell(app, table, fmt.Sprintf("dx=%d dy=%d ds=%d, ss=%d", dx, dy, ds, ss))

			}

			visited[key(hx[ROPLEN-1], hy[ROPLEN-1])]++
		}
		clearCell(app, table, hx, hy)
		setCell(app, table, hx, hy, visited)
		go func() {
			app.Draw()
		}()
		//time.Sleep(1000 * time.Millisecond)
	}
	logCell(app, table, fmt.Sprintf("visited %d\n", len(visited)))
}

func logCell(app *tview.Application, table *tview.Table, s string) {
	app.QueueUpdateDraw(func() {
		table.GetCell(0, SIZE).SetTextColor(tcell.ColorRed).SetText(s)
	})
}

func setCell(app *tview.Application, table *tview.Table, hx [ROPLEN]int, hy [ROPLEN]int, visited map[string]int) {
	app.QueueUpdateDraw(func() {
		table.GetCell(SIZE/2-hy[ROPLEN-1], hx[ROPLEN-1]+SIZE/2).SetText("T")
		for c := ROPLEN - 2; c > 0; c-- {
			table.GetCell(SIZE/2-hy[c], hx[c]+SIZE/2).SetText(strconv.Itoa(c))
		}
		table.GetCell(SIZE/2-hy[0], hx[0]+SIZE/2).SetText("H")
		for k, _ := range visited {
			vc := strings.Split(k, ",")
			vx, _ := strconv.Atoi(vc[0])
			vy, _ := strconv.Atoi(vc[1])
			table.GetCell(SIZE/2-vy, vx+SIZE/2).SetText("#")
		}
	})
}

func clearCell(app *tview.Application, table *tview.Table, hx [ROPLEN]int, hy [ROPLEN]int) {
	app.QueueUpdateDraw(func() {
		for r := 0; r < SIZE; r++ {
			for c := 0; c < SIZE; c++ {
				table.SetCell(r, c, tview.NewTableCell("."))
			}
		}
		table.SetCell(0, SIZE, tview.NewTableCell("."))
	})

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

func sign(n int) int {
	if n == 0 {
		return 0
	} else if n < 0 {
		return -1
	} else {
		return 1
	}

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
