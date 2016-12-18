package main

import (
	"bufio"
	"fmt"
	"os"
)

type theBot struct {
	pos   coord
	board []string
}

type coord struct {
	x, y int
}

const boardSize = 5

func main() {
	var (
		s   string
		bot theBot
	)
	rd := bufio.NewReader(os.Stdin)
	s, _ = rd.ReadString('\n')
	fmt.Sscanf(s, "%d %d", &bot.pos.x, &bot.pos.y)
	bot.board = make([]string, boardSize)
	for i := range bot.board {
		s, _ = rd.ReadString('\n')
		bot.board[i] = s[:boardSize]
	}
	bot.nextMove()
}

func (bot theBot) nextMove() {
	// search dirty cell
	cell, found := bot.findDirty()
	if !found {
		cell = bot.moveToSearch()
	}
	direction := bot.moveToCell(cell)
	fmt.Println(direction)
}

func (bot theBot) findDirty() (coord, bool) {
	var (
		x, y int
		cell coord
	)
	// at first look current cell
	cell = coord{bot.pos.x, bot.pos.y}
	if bot.checkCell(cell, 'd') {
		return cell, true
	}
	// look cell to right and left
	for y = bot.pos.y + 1; y >= bot.pos.y-1; y -= 2 {
		cell = coord{bot.pos.x, y}
		if bot.checkCell(cell, 'd') {
			return cell, true
		}
	}
	// look cell down and up
	for x = bot.pos.x + 1; x >= bot.pos.x-1; x -= 2 {
		cell = coord{x, bot.pos.y}
		if bot.checkCell(cell, 'd') {
			return cell, true
		}
	}

	// look cells at diags
	for x = bot.pos.x + 1; x >= bot.pos.x-1; x -= 2 {
		for y = bot.pos.y + 1; y >= bot.pos.y-1; y -= 2 {
			cell = coord{x, y}
			if bot.checkCell(cell, 'd') {
				return cell, true
			}
		}
	}
	return coord{}, false
}

func (bot *theBot) moveToCell(cell coord) string {
	switch {
	case bot.pos.y < cell.y:
		return "RIGHT"
	case bot.pos.y > cell.y:
		return "LEFT"
	case bot.pos.x > cell.x:
		return "UP"
	case bot.pos.x < cell.x:
		return "DOWN"
	default:
		return "CLEAN"
	}
}

func (bot theBot) moveToSearch() coord {
	var (
		x, y, c, closed int
	)

	candidate := bot.pos
	tmp := bot.pos
	// look cell to right and left
	for y = bot.pos.y + 1; y >= bot.pos.y-1; y -= 2 {
		bot.pos = coord{bot.pos.x, y}
		c = bot.countClosed()
		if c > closed {
			closed = c
			candidate = bot.pos
		}
	}
	// look cell to up and down
	for x = bot.pos.x + 1; x >= bot.pos.x-1; x -= 2 {
		bot.pos = coord{x, bot.pos.y}
		c = bot.countClosed()
		if c > closed {
			closed = c
			candidate = bot.pos
		}
	}

	bot.pos = tmp
	if candidate == bot.pos {
		panic("stalled")
	}
	return candidate
}

func (bot theBot) countClosed() int {
	var (
		cell      coord
		x, y, res int
	)
	// look cell to right and left
	for y = bot.pos.y + 1; y >= bot.pos.y-1; y -= 2 {
		cell = coord{bot.pos.x, y}
		if bot.checkCell(cell, 'o') {
			res++
		}
	}
	// look cell down and up
	for x = bot.pos.x + 1; x >= bot.pos.x-1; x -= 2 {
		cell = coord{x, bot.pos.y}
		if bot.checkCell(cell, 'o') {
			res++
		}
	}

	// look cells at diags
	for x = bot.pos.x + 1; x >= bot.pos.x-1; x -= 2 {
		for y = bot.pos.y + 1; y >= bot.pos.y-1; y -= 2 {
			cell = coord{x, y}
			if bot.checkCell(cell, 'o') {
				res++
			}
		}
	}
	return res
}

func (bot theBot) checkCell(c coord, v byte) bool {
	if !(c.x >= 0 && c.x < 5 && c.y >= 0 && c.y < 5) {
		return false
	}
	return bot.board[c.x][c.y] == v
}
