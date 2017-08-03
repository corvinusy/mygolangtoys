package main

import "fmt"

func main() {
	a := make([][]int, 3)
	for i := range a {
		a[i] = make([]int, 3)
		for j := range a[i] {
			fmt.Scan(&a[i][j])
		}
	}
	msquares := [][][]int{
		{{8, 1, 6}, {3, 5, 7}, {4, 9, 2}},
		{{6, 1, 8}, {7, 5, 3}, {2, 9, 4}},
		{{4, 9, 2}, {3, 5, 7}, {8, 1, 6}},
		{{2, 9, 4}, {7, 5, 3}, {6, 1, 8}},
		{{8, 3, 4}, {1, 5, 9}, {6, 7, 2}},
		{{4, 3, 8}, {9, 5, 1}, {2, 7, 6}},
		{{6, 7, 2}, {1, 5, 9}, {8, 3, 4}},
		{{2, 7, 6}, {9, 5, 1}, {4, 3, 8}},
	}
	min := 15 * 9
	var diff int
	for i := range msquares {
		diff = calculateDiff(a, msquares[i])
		if diff < min {
			min = diff
		}
	}
	fmt.Println(min)
}

func calculateDiff(a, b [][]int) int {
	diff := 0
	for i := range a {
		for j := range a[i] {
			diff += abs(a[i][j] - b[i][j])
		}
	}
	return diff
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
