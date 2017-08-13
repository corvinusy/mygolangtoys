package main

import "fmt"

func main() {
	var t int
	for fmt.Scan(&t); t > 0; t-- {
		doRepl()
	}
}

func doRepl() {
	// read
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}
	// evaluate
	result := doEval(a)
	// print
	fmt.Println(result)
}

func doEval(a []int) int {
	var diffHighLow, diffLowHigh int
	var low, high, tmpLow, tmpHigh int
	for i := 1; i < len(a); i++ {
		diffHighLow = a[i-1] - 1
		diffLowHigh = a[i] - 1
		tmpLow = max(low, high+diffHighLow)
		tmpHigh = max(high, low+diffLowHigh)
		low = tmpLow
		high = tmpHigh
	}
	return max(low, high)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
