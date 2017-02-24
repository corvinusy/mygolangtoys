package main

import (
	"bufio"
	"fmt"
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

var mainLine = []cell{{0, 1}, {1, 1}, {2, 1}, {3, 1}, {4, 1}}

func main() {
	bot := new(Agent)
	bot.getWorldState()
	println(bot.nextMove())
}

func (a *Agent) getWorldState() error {
	var (
		err error
		s   string
	)
	rd := bufio.NewReader(os.Stdin)
	s, err = rd.ReadString('\n')
	if err != nil {
		return err
	}
	fmt.Sscanf(s, "%d %d", &a.pos.x, &a.pos.y)
	a.board = make([]string, boardSize)
	for i := range a.board {
		s, err = rd.ReadString('\n')
		if err != nil {
			return err
		}
		a.board[i] = s[:boardSize]
	}
	return nil
}

func (a *Agent) nextMove() string {
	// if isDirty then CLEAN
	if a.isOnDirtyCell() {
		return "CLEAN"
	}
	// else
	// if not on MainLine then move to closest Main Line
	if !a.isOnMainLine() {
		return a.moveTo(a.closestMainLine())
	}
	// if on MainLine then find dirty
	if dirtyCell, found := a.findDirtyCell(); found {
		return a.moveTo(dirtyCell)
	}
	return ""
}

func (a *Agent) isOnDirtyCell() bool {
	return a.board[a.pos.x][a.pos.y] == 'd'
}

func (a *Agent) isOnMainLine() bool {
	for i := range mainLine {
		if a.pos == mainLine[i] {
			return true
		}
	}
	return false
}

func (a *Agent) findDirtyCell() (cell, bool) {
	return cell{}, false
}

func (a *Agent) closestMainLine() cell {
	switch a.pos.x {
	case 0:
		return cell{1, a.pos.y}
	case 2, 4:
		return cell{3, a.pos.y}
	default:
		return cell{a.pos.x, a.pos.y}
	}
}

func (a *Agent) moveTo(c cell) string {
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
