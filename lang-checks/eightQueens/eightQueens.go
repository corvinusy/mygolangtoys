package main

import 	"fmt"

const SIZE=12

type Solution [SIZE]int
/*Solution = list of integers, where index is horizontal, and value is vertical*/

type Result []Solution;

func printSolution(s Solution) {
	for i := 0; i < SIZE; i++ {
		fmt.Print("--")
	}
	fmt.Println();
	for i := 0; i < len(s); i++ {
		for j:=0; j<s[i]; j++ {
			fmt.Print("  ") //padding spaces up to Q in solution
		}
		fmt.Println("FF")
	}
	for i := 0; i < SIZE; i++ {
		fmt.Print("--")
	}
	fmt.Println();
}

func printResult(pr *Result) {
	fmt.Printf("Found %d solutions \n", len(*pr))
	for i := 0; i < len(*pr); i++ {
		fmt.Printf("Solution %d\n",i)
		printSolution((*pr)[i])
	}
	fmt.Printf("Found %d solutions \n", len(*pr))
	fmt.Println("done in Go")
}

func isEqual(s1 *Solution, s2 *Solution) bool {
	for i:=0; i<SIZE; i++ {
		if (*s1)[i] != (*s2)[i] {
			return false
		}
	}
	return true
}

func getVMirroredSolution (ps *Solution) Solution {
	var s Solution
	for i := 0; i < SIZE; i++ {
		if (*ps)[i] >= SIZE/2 {
			s[i] = (*ps)[i] - SIZE/2
		} else {
			s[i] = (*ps)[i] + SIZE/2
		}
	}
	return s
}

func getHMirroredSolution (ps *Solution) Solution {
	var s Solution
	for i := 0; i < SIZE; i++ {
		s[i] = (*ps)[SIZE-i-1]
	}
	return s
}

func getRotated90Solution (ps *Solution) Solution {
	var s Solution
	for i := 0; i<SIZE; i++ {
		s[SIZE-(*ps)[i]-1] = i
	}
	return s
}

func getRotated180Solution (ps *Solution) Solution {
	var s Solution
	s = getRotated90Solution(ps)
	s = getRotated90Solution(&s)
	return s
}

func getRotated270Solution (ps *Solution) Solution {
	var s Solution
	s = getRotated90Solution(ps)
	s = getRotated90Solution(&s)
	s = getRotated90Solution(&s)
	return s
}

func alreadyHave (pr *Result, ps *Solution) bool {
	var sH, sV, s90, s180, s270 Solution
	sH = getHMirroredSolution(ps)
	sV = getVMirroredSolution(ps)
	s90 = getRotated90Solution(ps)
	s180 = getRotated180Solution(ps)
	s270 = getRotated270Solution(ps)

	for i:=0; i<len(*pr); i++ {
		if isEqual(&((*pr)[i]), ps) {
			return true
		} else if isEqual(&((*pr)[i]), &sH) {
			return true
		} else if isEqual(&((*pr)[i]), &sV) {
			return true
		} else if isEqual(&((*pr)[i]), &s90) {
			return true
		} else if isEqual(&((*pr)[i]), &s180) {
			return true
		} else if isEqual(&((*pr)[i]), &s270) {
			return true
		}
	}
	return false
}

func canPlaceQueen (ps *Solution, x, y int) bool {
	for i:=0; i<x; i++ {
		if isQueensAct( i, (*ps)[i], x, y) {
			return false
		}
	}
	return true
}

func isQueensAct(x1, y1, x2, y2 int) bool {
	if y1 == y2 { //vertical comparizon
		return true
	} else if x1 == x2 { //horizontal comparizon
		return true
	} else if intAbs(y1-y2) == intAbs(x1-x2) { // diagonal comparizon
		return true
	}
	return false
}

func intAbs (x int) int {
	if x<0 {
		return -x
	}
	return x
}

func findSolution(s Solution, pr *Result, x int) {
	if x == SIZE {
		if !(alreadyHave(pr, &s)) {
			*pr = append(*pr,s)
		}
	} else {
		for y := 0; y < SIZE; y++ {
			if canPlaceQueen(&s, x, y) {
				s[x] = y
				findSolution(s, pr, x+1) //this branch goes for next horizontal (index)
				continue // this branch goes for next vertical (value)
			}
		}
	}
}

func main() {
	fmt.Print ("Looking for solutions...")
//	var r Result //global var - not good
	pr := new(Result)
	var sol Solution //global var - not good
	findSolution(sol, pr, 0)
	printResult(pr)
}
