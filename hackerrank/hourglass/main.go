package main

import "fmt"

const (
	size   = 6
	hgSize = 3
)

func main() {
	x := make([][]int, size)
	for i := range x {
		x[i] = make([]int, size)
		for j := range x[i] {
			fmt.Scan(&x[i][j])
		}
	}
	fmt.Println(getAllSum(x))
}

func getAllSum(x [][]int) int {
	max := 0
	sum := 0
	for i := 0; i < size-hgSize+1; i++ {
		for j := 0; j < size-hgSize+1; j++ {
			sum = getHgSum(x, i, j)
			if max < sum {
				max = sum
			}
		}
	}
	return max
}

func getHgSum(x [][]int, r, c int) int {
	sum := 0
	for i := r; i < r+hgSize; i++ {
		for j := c; j < c+hgSize; j++ {
			sum += x[i][j]
		}
	}
	sum = sum - x[r+1][c] - x[r+1][c+2]
	fmt.Println(sum)
	return sum
}
