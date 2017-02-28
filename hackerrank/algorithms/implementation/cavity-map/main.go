package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	bbs := make([][]byte, n)
	for i := range bbs {
		bbs[i] = make([]byte, n)
		fmt.Scanln(&bbs[i])
	}

	findCavities(bbs)

	for i := range bbs {
		fmt.Println(string(bbs[i]))
	}
}

func findCavities(bbs [][]byte) {
	bl := len(bbs)
	for i := 1; i < bl-1; i++ {
		for j := 1; j < bl-1; j++ {
			if isCavity(bbs, i, j) {
				bbs[i][j] = 'X'
			}
		}
	}
}

func isCavity(bbs [][]byte, i, j int) bool {
	return bbs[i][j] > bbs[i-1][j] &&
		bbs[i][j] > bbs[i+1][j] &&
		bbs[i][j] > bbs[i][j-1] &&
		bbs[i][j] > bbs[i][j+1]
}
