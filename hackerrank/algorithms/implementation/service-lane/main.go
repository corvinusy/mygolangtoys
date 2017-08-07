package main

import "fmt"

func main() {
	var n, q, start, end int
	fmt.Scan(&n, &q)
	widths := make([]int, n)
	for i := range widths {
		fmt.Scan(&widths[i])
	}
	for ; q > 0; q-- {
		fmt.Scan(&start, &end)
		fmt.Println(findMinWidth(widths, start, end))
	}
}

func findMinWidth(widths []int, start, end int) int {
	if end >= len(widths) {
		end = len(widths) - 1
	}
	min := 3
	for i := start; i <= end; i++ {
		if widths[i] < min {
			min = widths[i]
		}
	}
	return min
}
