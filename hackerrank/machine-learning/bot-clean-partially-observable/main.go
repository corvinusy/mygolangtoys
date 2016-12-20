package main

import (
	"bufio"
	"fmt"
	"os"
)

type theBot struct {
	pos   cell
	board []string
}

type cell struct {
	x, y int
}

type quadrant struct {
	xmin, xmax, ymin, ymax int
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

func (bot *theBot) nextMove() {
	// find quadrant
	q := bot.findQuadrant()
	// search dirty cell
	c, found := bot.findDirtyCell(q)
	if !found {
		c = bot.moveToNextQuadrant()
	}
	direction := bot.moveToCell(c)
	fmt.Println(direction)
}

func (bot *theBot) findQuadrant() quadrant {
	switch {
	case bot.pos.x < 3 && bot.pos.y < 3:
		return quadrant{0, 2, 0, 2}
	case bot.pos.x < 3 && bot.pos.y >= 3:
		return quadrant{0, 2, 3, 4}
	case bot.pos.x >= 3 && bot.pos.y < 3:
		return quadrant{3, 4, 0, 2}
	case bot.pos.x >= 3 && bot.pos.y >= 3:
		return quadrant{3, 4, 3, 4}
	}
	panic("out of board")
}

func (bot *theBot) findDirtyCell(q quadrant) (cell, bool) {
	var (
		x, y int
		c    cell
	)
	// at first look current cell
	c = cell{bot.pos.x, bot.pos.y}
	if bot.checkCell(c, 'd') {
		return c, true
	}
	// look cell to left and right
	for y = bot.pos.y - 1; y >= bot.pos.y+1; y += 2 {
		c = cell{bot.pos.x, y}
		if bot.checkQuadrant(q, c) && bot.checkCell(c, 'd') {
			return c, true
		}
	}
	// look at cell up and down
	for x = bot.pos.x - 1; x <= bot.pos.x+1; x += 2 {
		c = cell{x, bot.pos.y}
		if bot.checkQuadrant(q, c) && bot.checkCell(c, 'd') {
			return c, true
		}
	}

	// look cells at diags
	for x = bot.pos.x - 1; x <= bot.pos.x+1; x += 2 {
		for y = bot.pos.y + 1; y >= bot.pos.y-1; y -= 2 {
			c = cell{x, y}
			if bot.checkQuadrant(q, c) && bot.checkCell(c, 'd') {
				return c, true
			}
		}
	}
	return cell{}, false
}

func (bot *theBot) moveToCell(c cell) string {
	switch {
	case bot.pos.y < c.y:
		return "RIGHT"
	case bot.pos.y > c.y:
		return "LEFT"
	case bot.pos.x > c.x:
		return "UP"
	case bot.pos.x < c.x:
		return "DOWN"
	default:
		return "CLEAN"
	}
}

func (bot *theBot) moveToNextQuadrant() cell {
	switch {
	case bot.pos.x < 3 && bot.pos.y < 3:
		if bot.pos.x != 1 {
			return cell{1, 1}
		}
		return cell{1, 3}
	case bot.pos.x < 3 && bot.pos.y >= 3:
		return cell{3, 3}
	case bot.pos.x >= 3 && bot.pos.y >= 3:
		return cell{3, 1}
	}
	return cell{1, 1}
}

func (bot *theBot) checkQuadrant(q quadrant, c cell) bool {
	return c.x >= q.xmin && c.x <= q.xmax && c.y >= q.ymin && c.y <= q.ymax
}

func (bot *theBot) checkCell(c cell, v byte) bool {
	return bot.board[c.x][c.y] == v
}
