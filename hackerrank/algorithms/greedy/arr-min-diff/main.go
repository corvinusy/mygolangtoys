package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n) //nolint
	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i]) //nolint
	}
	sort.Ints(a)
	fmt.Println(findMinDiff(a))
}

func findMinDiff(a []int) int {
	minDiff := abs(a[0] - a[1])
	var diff int
	for i := 1; i < len(a)-1; i++ {
		diff = abs(a[i] - a[i+1])
		if diff < minDiff {
			minDiff = diff
		}
	}
	return minDiff
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
