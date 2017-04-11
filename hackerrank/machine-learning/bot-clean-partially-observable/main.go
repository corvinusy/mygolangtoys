package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Agent ...
type Agent struct {
	pos   cell
	board []string
}

type cell struct {
	x, y int
}

const boardSize = 5

var referPath = []cell{{0, 1}, {1, 1}, {2, 1}, {3, 1}, {4, 1}, {4, 2}, {4, 3}}

//-----------------------------------------------------------------------------
func main() {
	bot := new(Agent)
	err := bot.getWorld()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(bot.nextMove())
}

//-----------------------------------------------------------------------------
func (a *Agent) getWorld() error {
	var (
		err error
		s   string
	)
	rd := bufio.NewReader(os.Stdin)
	s, err = rd.ReadString('\n')
	if err != nil {
		return err
	}
	_, err = fmt.Sscanf(s, "%d %d", &a.pos.x, &a.pos.y)
	if err != nil {
		return err
	}
	a.board = make([]string, boardSize)
	for i := range a.board {
		s, _ = rd.ReadString('\n')
		a.board[i] = s[:boardSize]
	}
	return nil
}

//-----------------------------------------------------------------------------
func (a Agent) nextMove() string {
	if a.isOnDirtyCell() {
		return "CLEAN"
	}
	// findTarget
	c := a.findTarget()
	// generateCmd
	return a.generateCmd(c)
}

//-----------------------------------------------------------------------------
func (a Agent) isOnDirtyCell() bool {
	return a.board[a.pos.x][a.pos.y] == 'd'
}

//-----------------------------------------------------------------------------
func (a Agent) findTarget() cell {
	switch {
	case !a.isOnMainPath():
		return a.goMainPath()
	case a.isDirtyToLeft():
		return a.goLeft()
	case a.isDirtyToRight():
		return a.goRight()
	default:
		return a.goNextRefer()
	}
}

//-----------------------------------------------------------------------------
func (a Agent) isOnMainPath() bool {
	if a.pos.y > 2 {
		return true
	}
	for i := range referPath {
		if a.pos == referPath[i] {
			return true
		}
	}
	return false
}

//-----------------------------------------------------------------------------
func (a Agent) goMainPath() cell {
	var c cell
	c.x = a.pos.x
	if a.pos.y < 3 {
		c.y = 1
	} else {
		c.y = 3
	}
	return c
}

//-----------------------------------------------------------------------------
func (a Agent) isDirtyToLeft() bool {
	if a.pos.y == 0 || a.pos.y == 3 {
		return false
	}
	return a.board[a.pos.x][a.pos.y-1] == 'd'
}

//-----------------------------------------------------------------------------
func (a Agent) isDirtyToRight() bool {
	if a.pos.y == 4 {
		return false
	}
	return a.board[a.pos.x][a.pos.y+1] == 'd'
}

//-----------------------------------------------------------------------------
func (a Agent) goLeft() cell {
	if a.pos.y == 0 {
		return cell{a.pos.x, a.pos.y}
	}
	return cell{a.pos.x, a.pos.y - 1}
}

//-----------------------------------------------------------------------------
func (a Agent) goRight() cell {
	if a.pos.y == 4 {
		return cell{a.pos.x, a.pos.y}
	}
	return cell{a.pos.x, a.pos.y + 1}
}

//-----------------------------------------------------------------------------
func (a Agent) goNextRefer() cell {
	for i := 0; i < len(referPath)-1; i++ {
		if a.pos == referPath[i] {
			return referPath[i+1]
		}
	}
	if a.pos.y > 2 && a.pos.x > 0 {
		return cell{a.pos.x - 1, a.pos.y}
	}
	return referPath[0]
}

//-----------------------------------------------------------------------------
func (a Agent) generateCmd(c cell) string {
	switch {
	case a.pos.x > c.x:
		return "UP"
	case a.pos.x < c.x:
		return "DOWN"
	case a.pos.y < c.y:
		return "RIGHT"
	case a.pos.y > c.y:
		return "LEFT"
	default:
		return "CLEAN"
	}
}
