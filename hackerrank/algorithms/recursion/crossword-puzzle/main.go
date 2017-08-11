package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

const (
	boardSize = 10
	plus      = byte('+')
	minus     = byte('-')
)

type cell struct {
	x, y    int
	isVert  bool
	wordLen int
}

func main() {
	rd := bufio.NewReader(os.Stdin)
	// read board
	board := make([][]byte, boardSize)
	for i := range board {
		board[i], _ = rd.ReadBytes('\n')
		board[i] = bytes.TrimSpace(board[i])
	}
	// read words
	s, _ := rd.ReadBytes('\n')
	s = bytes.TrimSpace(s)
	words := bytes.Split(s, []byte(";"))
	// evaluate
	cells := findStartCells(board)
	//fmt.Println(cells)
	result := fillBoard(board, words, cells)
	// print
	printBoard(result)
}

func printBoard(b [][]byte) {
	for i := range b {
		fmt.Println(string(b[i]))
	}
}

func fillBoard(b, w [][]byte, c []cell) [][]byte {
	if isBoardFilled(b) {
		return b
	}
	for i := range w {
		for j := range c {
			if isPossibleToPlace(b, w[i], c[j]) {
				//fmt.Println(string(w[i])) // DEBUG
				nb := placeWord(b, w[i], c[j])
				nw := filterWords(w, i)
				nc := filterCells(c, j)
				bb := fillBoard(nb, nw, nc)
				if bb == nil {
					continue
				}
				return bb
			}
		}
	}
	return nil
}

func placeWord(b [][]byte, w []byte, c cell) [][]byte {
	nb := make([][]byte, boardSize)
	for i := range nb {
		nb[i] = make([]byte, boardSize)
		for j := range nb[i] {
			nb[i][j] = b[i][j]
		}
	}
	if c.isVert {
		for i := range w {
			nb[c.x+i][c.y] = w[i]
		}
	} else {
		for i := range w {
			nb[c.x][c.y+i] = w[i]
		}

	}
	return nb
}

func isPossibleToPlace(b [][]byte, w []byte, c cell) bool {
	if len(w) != c.wordLen {
		return false
	}
	if c.isVert {
		for i := range w {
			if b[c.x+i][c.y] == minus {
				continue
			}
			if b[c.x+i][c.y] != w[i] {
				return false
			}
		}
	} else {
		for i := range w {
			if b[c.x][c.y+i] == minus {
				continue
			}
			if b[c.x][c.y+i] != w[i] {
				return false
			}
		}
	}
	return true
}

func filterWords(w [][]byte, idx int) [][]byte {
	nw := make([][]byte, 0)
	for i := range w {
		if i != idx {
			nw = append(nw, w[i])
		}
	}
	return nw
}

func filterCells(c []cell, idx int) []cell {
	nc := make([]cell, 0)
	for i := range c {
		if i != idx {
			nc = append(nc, c[i])
		}
	}
	return nc
}

func findStartCells(b [][]byte) []cell {
	var l int
	cells := make([]cell, 0)
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b); j++ {
			if b[i][j] == plus {
				continue
			} else {
				//vert
				if i == 0 ||
					(i+1 < boardSize && b[i-1][j] == plus && b[i+1][j] != plus) {
					l = findSpaceLen(b, cell{i, j, true, 0})
					cells = append(cells, cell{i, j, true, l})
				}
				// hor
				if j == 0 ||
					(j+1 < boardSize && b[i][j-1] == plus && b[i][j+1] != plus) {
					l = findSpaceLen(b, cell{i, j, false, 0})
					cells = append(cells, cell{i, j, false, l})
				}
			}
		}
	}
	return cells
}

func findSpaceLen(b [][]byte, c cell) int {
	var l int
	if c.isVert {
		for i := c.x; i < boardSize && b[i][c.y] != plus; i++ {
			l++
		}
	} else {
		for i := c.y; i < boardSize && b[c.x][i] != plus; i++ {
			l++
		}
	}
	return l
}

func isBoardFilled(b [][]byte) bool {
	for i := range b {
		for j := range b[i] {
			if b[i][j] == minus {
				return false
			}
		}
	}
	return true
}
