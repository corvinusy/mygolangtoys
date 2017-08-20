package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}
	b := makeGreedyBuying(a, k)
	// print
	fmt.Println(b)
}

func makeGreedyBuying(a []int, k int) int {
	var count int
	// make sort reversed
	sort.Ints(a)
	// push to knapsack
	for i := range a {
		if k > a[i] {
			k = k - a[i]
			count++
		}
	}
	return count
}
