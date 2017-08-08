package main

import (
	"fmt"
)

const boardSize = 16

func main() {
	var t, x, y int
	board := make([][]int, boardSize)
	for i := range board {
		board[i] = make([]int, boardSize)
	}
	setBoardOddity(board)
	for fmt.Scan(&t); t > 0; t-- {
		fmt.Scan(&x, &y)
		if isOddSquare(board, x, y) {
			fmt.Println("Second")
		} else {
			fmt.Println("First")
		}
	}
}

func isOddSquare(board [][]int, x, y int) bool {
	return board[x-1][y-1] == 2
}

func setBoardOddity(board [][]int) {
	for y := 0; y < boardSize; y++ {
		for x := 0; x < boardSize; x++ {
			// reverse move 1
			if x%4 < 2 && y%4 < 2 {
				board[x][y] = 2
			} else {
				board[x][y] = 1
			}
		}
	}
	// printBoard(board)
	// fmt.Println()
}

// func printBoard(board [][]int) {
// 	for i := range board {
// 		fmt.Println(board[i])
// 	}
// 	fmt.Println()
// }
