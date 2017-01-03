package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	sides := make([]int, n)
	for i := range sides {
		fmt.Scan(&sides[i])
	}
	sort.Ints(sides)
	for i := len(sides) - 1; i > 1; i-- {
		if sides[i] < sides[i-1]+sides[i-2] {
			fmt.Println(sides[i-2], sides[i-1], sides[i])
			return
		}
	}
	fmt.Println(-1)
}
