package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	weis := make([]int, n)
	for i := range weis {
		fmt.Scan(&weis[i])
	}
	sort.Ints(weis)
	result := solve(weis)
	fmt.Println(result)
}

func solve(w []int) int {
	const diff = 4
	groupStart := -1 - diff
	groups := 0
	for i := range w {
		if w[i] <= groupStart+diff {
			continue
		} else {
			groups++
			groupStart = w[i]
		}
	}
	return groups
}
