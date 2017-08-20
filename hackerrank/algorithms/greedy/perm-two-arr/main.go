package main

import (
	"fmt"
	"sort"
)

func main() {
	var t int
	for fmt.Scan(&t); t > 0; t-- {
		var n, k int
		fmt.Scan(&n, &k)
		a := make([]int, n)
		b := make([]int, n)
		for i := range a {
			fmt.Scan(&a[i])
		}
		for i := range b {
			fmt.Scan(&b[i])
		}
		result := solve(a, b, k)
		// print
		if result {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func solve(a, b []int, k int) bool {
	// make sorts
	sort.Ints(a)
	sort.Sort(sort.Reverse(sort.IntSlice(b)))
	// push to knapsack
	for i := range a {
		if a[i]+b[i] < k {
			return false
		}
	}
	return true
}
