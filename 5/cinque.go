package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type stack []string

func (s stack) push(v string) stack {
	return append(s, v)
}

func (s stack) pop() (stack, string) {
	l := len(s)
	if l == 0 {
		panic("stack underflow")
	}
	return s[:l-1], s[l-1]
}

func (s stack) insert(v string) stack {
	if len(s) == 0 {
		s = append(s, v)
	} else {
		var a stack
		a = append(a, v)
		s = append(a, s...)
	}
	return s
}

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

	const skip = 4
	const stacksNo = 9

	var stacks [9]stack

	for scanner.Scan() {
		var line = scanner.Text()
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		if len(line) == 0 {
			continue
		}

		//log.Printf("line %s: ", line)

		if line[0] == 'm' {
			movePack2(line, &stacks)
		} else {
			for i := 0; i < stacksNo; i++ {
				pack := string(line[i*skip+1])
				_, err = strconv.Atoi(pack)
				if err == nil {
					continue
				}
				//log.Printf("pack %s\n", pack)
				if pack != " " {
					stacks[i] = stacks[i].insert(pack)
				}
			}
			log.Printf("stacks %s\n", stacks)
		}
	}
	fmt.Printf("result 1:")
	for i := 0; i < stacksNo; i++ {
		_, top := stacks[i].pop()
		fmt.Printf("%s", top)
	}
	fmt.Printf("\n")
}

func movePack(line string, stacks *[9]stack) {
	var n, from, to int
	_, err := fmt.Sscanf(line, "move %d from %d to %d", &n, &from, &to)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("move %d from %d to %d:\n", n, from, to)
	for i := 0; i < n; i++ {
		s1, pack := stacks[from-1].pop()
		fmt.Printf("popped %s\n", pack)
		s2 := stacks[to-1].push(pack)

		// update stacks
		stacks[from-1] = s1
		stacks[to-1] = s2
	}
	log.Printf("stacks %s", stacks)
}

func movePack2(line string, stacks *[9]stack) {
	var n, from, to int
	_, err := fmt.Sscanf(line, "move %d from %d to %d", &n, &from, &to)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("move %d from %d to %d:\n", n, from, to)

	fs := stacks[from-1]

	picked := fs[len(fs)-n:]
	fmt.Printf("picked %s\n", picked)
	for i := (n - 1); i >= 0; i-- {
		stacks[from-1], _ = stacks[from-1].pop()
	}
	stacks[to-1] = append(stacks[to-1], picked...)

	log.Printf("stacks %s", stacks)
}
