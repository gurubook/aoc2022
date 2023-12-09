package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/maja42/goval"
)

func main() {
	// open file
	f, err := os.Open("case.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	rightOrdered := 0
	item := 1

	for scanner.Scan() {
		var line = scanner.Text()
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		leftLine := line

		scanner.Scan()
		line = scanner.Text()
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		rightLine := line

		// skip blank line
		scanner.Scan()

		left := eval(leftLine)
		right := eval(rightLine)

		fmt.Printf("\n == Item %d == \n", item)
		fmt.Printf("\tleft\t %s\n", leftLine)
		fmt.Printf("\tright\t %s\n", rightLine)
		fmt.Printf("\n")

		if compare(left.([]any), right.([]any), false) {
			fmt.Printf("ORDERED\n")
			rightOrdered += item
		} else {
			fmt.Printf("UNORDERED\n")
		}

		item++
	}

	fmt.Printf("Right ordered %d\n", rightOrdered)
}

// func isArray(e any) ([]any, bool) {
// 	a, ok := e.([]any)
// 	return a, ok
// }

func isInt(e any) (int, bool) {
	if a, ok := e.(int); ok {
		return a, ok
	}
	return -1, false
}

func compare(lany, rany any, checkOnlyFirst bool) bool {
	// is lany a int ?
	fmt.Printf("Compare %v == %v\n", lany, rany)
	if ln, okl := isInt(lany); okl {
		// how about rany ?
		if rn, okr := isInt(rany); okr {
			// r is int
			if ln > rn {
				fmt.Printf("%d > %d\n", ln, rn)
				return false
			} else {
				fmt.Printf("%d <= %d\n", ln, rn)
			}

		} else {
			// r is arr so convert l  and compare
			// if len(rany.([]any)) == 0) {
			// 	return false
			// }
			return compare([]any{lany}, rany, true)
		}
	} else {
		// lany is an array
		// how about rany ?
		if rn, okr := isInt(rany); okr {
			// convert right so check only first element
			return compare(lany, []any{rn}, true)
		} else {
			// both are array, check elements
			for i, ln := range lany.([]any) {
				// if ()
				fmt.Printf("\tCompare %v == %v\n", ln, rany)

				if i > len(rany.([]any))-1 {
					fmt.Printf("right run out of items\n")
					return false
				}
				rn := rany.([]any)[i]
				if !compare(ln, rn, false) {
					return false
				}
				if checkOnlyFirst {
					break
				}
			}
		}
	}
	return true
}

func eval(s string) any {
	eval := goval.NewEvaluator()
	result, err := eval.Evaluate(s, nil, nil) // Returns <true, nil>
	if err != nil {
		log.Fatal(err)
	}
	return result
}
