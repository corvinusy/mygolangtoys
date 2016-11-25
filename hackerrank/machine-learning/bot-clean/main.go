package main

import (
	"bufio"
	"fmt"
	"os"
)

type point struct {
	x int
	y int
}

type data struct {
	bot   point
	dirty []point
	moves int
}

func main() {
	const gs = 5 // grid size
	var d data
	rd := bufio.NewReader(os.Stdin)
	bs, _ := rd.ReadString('\n')
	fmt.Sscan(bs, &d.bot.x, &d.bot.y)
	for i := 0; i < gs; i++ {
		s, _ := rd.ReadBytes('\n')
		for k := 0; k < gs; k++ {
			if s[k] == 'd' {
				d.dirty = append(d.dirty, point{i, k})
			}
		}
	}

	d.moveBot()
	/*for d.moveBot() {
	}*/
}

func (d *data) moveBot() bool {
	g := d.findNearest()
	pl := d.planMove(g)
	fmt.Println(pl)
	switch pl {
	case "DOWN":
		d.bot.x++
		d.moves++
	case "UP":
		d.bot.x--
		d.moves++
	case "RIGHT":
		d.bot.y++
		d.moves++
	case "LEFT":
		d.bot.y--
		d.moves++
	case "CLEAN":
		d.cleanCell(g)
		d.moves++
	default:
		panic("bad state")
	}
	return len(d.dirty) != 0
}

func (d *data) cleanCell(g point) {
	for i := range d.dirty {
		if d.dirty[i].x == g.x && d.dirty[i].y == g.y {
			//			if i+1 < len(d.dirty) {
			d.dirty = append(d.dirty[:i], d.dirty[i+1:]...)
			//			}
			break
		}
	}
}

func (d *data) findNearest() point {
	var ds int
	minDs := 5 * 2
	var dirty point
	for i := range d.dirty {
		ds = d.getDist(d.dirty[i])
		if ds < minDs {
			minDs = ds
			dirty = d.dirty[i]
		}
	}
	return dirty
}

func (d *data) getDist(g point) int {
	return abs(g.x-d.bot.x) + abs(g.y-d.bot.y)
}

func (d *data) planMove(g point) string {
	dx := g.x - d.bot.x
	dy := g.y - d.bot.y
	if dx > 0 {
		return "DOWN"
	} else if dx < 0 {
		return "UP"
	} else if dy > 0 {
		return "RIGHT"
	} else if dy < 0 {
		return "LEFT"
	}
	// else
	return "CLEAN"
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
