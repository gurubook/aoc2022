package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	id          int
	items       []*big.Int
	opLine      string
	divisor     *big.Int
	trueTarget  int
	falseTarget int
	active      int64
}

const ROUND = 10000
const WORRY_DIV = 1

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

	var monkeys = make([]monkey, 0)
	superModulo := big.NewInt(1)

	for scanner.Scan() {
		var line = readLine(scanner)
		if strings.HasPrefix(line, "Monkey ") {
			var mid int
			fmt.Sscanf(line, "Monkey %d:", &mid)
			fmt.Printf("Monkey %d\n", mid)

			m := monkey{id: mid}

			scanner.Scan()
			line = readLine(scanner)
			itemsLine := line[18:]
			fmt.Printf("items : ")
			m.items = make([]*big.Int, 0)
			for _, itemStr := range strings.Split(itemsLine, ", ") {
				item, _ := new(big.Int).SetString(itemStr, 10)
				m.items = append(m.items, item)

				fmt.Printf("%d ", item)
			}
			fmt.Printf("\n")

			scanner.Scan()
			line = readLine(scanner)
			opLine := line[19:]
			m.opLine = opLine
			fmt.Printf("operation : %s\n", m.opLine)

			scanner.Scan()
			line = readLine(scanner)
			testLine := line[21:]
			m.divisor, _ = new(big.Int).SetString(testLine, 10)
			fmt.Printf("test divisible by : %d\n", m.divisor)

			superModulo = superModulo.Mul(superModulo, m.divisor)

			scanner.Scan()
			line = readLine(scanner)
			trueLine := line[29:]
			m.trueTarget, _ = strconv.Atoi(trueLine)
			fmt.Printf("true throw to: %d\n", m.trueTarget)

			scanner.Scan()
			line = readLine(scanner)
			falseLine := line[30:]
			m.falseTarget, _ = strconv.Atoi(falseLine)
			fmt.Printf("false throw to: %d\n", m.falseTarget)

			monkeys = append(monkeys, m)
		}
	}

	for l := 0; l < ROUND; l++ {
		// process
		for mid := 0; mid < len(monkeys); mid++ {
			removed := make([]int, 0)
			for idx, item := range monkeys[mid].items {
				monkeys[mid].active++
				// op
				var left, operator, right string
				fmt.Sscanf(monkeys[mid].opLine, "%s %s %s", &left, &operator, &right)
				var lv, rv *big.Int
				var worry *big.Int

				// // if left == "old" && right == "old" && operator == "*" {
				// if true {
				// 	worry = item
				// } else {
				if left == "old" {
					lv = item
				} else {
					lv, _ = new(big.Int).SetString(left, 10)
				}
				if right == "old" {
					rv = item
				} else {
					rv, _ = new(big.Int).SetString(right, 10)
				}

				switch operator {
				case "+":
					worry = lv.Add(lv, rv)
				case "-":
					worry = lv.Sub(lv, rv)
				case "*":
					worry = lv.Mul(lv, rv)
				case "/":
					worry = lv.Quo(lv, rv)
				}
				if WORRY_DIV > 1 {
					worry = worry.Quo(worry, big.NewInt(WORRY_DIV))
					fmt.Printf("Worry %s\n", worry)
				} else {
					worry = worry.Rem(worry, superModulo)
				}
				// }

				// test
				var target int
				rem := big.NewInt(0)
				rem = rem.Rem(worry, monkeys[mid].divisor)
				if rem.Cmp(big.NewInt(0)) == 0 {
					target = monkeys[mid].trueTarget
				} else {
					target = monkeys[mid].falseTarget
				}
				removed = append(removed, idx)
				monkeys[target].items = append(monkeys[target].items, worry)
			}

			// remove throwed stuff
			sort.Slice(removed, func(i, j int) bool {
				return removed[i] > removed[j]
			})
			for i := 0; i < len(removed); i++ {
				monkeys[mid].items = remove(monkeys[mid].items, removed[i])
			}
		}
		if l%1000 == 0 {
			fmt.Printf("*")
		}
	}

	// sort for active
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].active > monkeys[j].active
	})
	fmt.Printf("Monkeys\n")
	fmt.Printf("Monkeys business %d\n", monkeys[0].active*monkeys[1].active)

}

func remove(slice []*big.Int, s int) []*big.Int {
	return append(slice[:s], slice[s+1:]...)
}

func readLine(scanner *bufio.Scanner) string {
	var line = scanner.Text()
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return line
}
