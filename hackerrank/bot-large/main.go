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
}

func main() {
	var d data
	var h, w int
	rd := bufio.NewReader(os.Stdin)
	bs, _ := rd.ReadString('\n')
	fmt.Sscan(bs, &d.bot.x, &d.bot.y)
	bs, _ = rd.ReadString('\n')
	fmt.Sscan(bs, &h, &w)
	for i := 0; i < h; i++ {
		s, _ := rd.ReadBytes('\n')
		for k := 0; k < w; k++ {
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
	g := d.findNext()
	pl := d.planMove(g)
	fmt.Println(pl)
	switch pl {
	case "DOWN":
		d.bot.x++
	case "UP":
		d.bot.x--
	case "RIGHT":
		d.bot.y++
	case "LEFT":
		d.bot.y--
	default: // "CLEAN":
		d.cleanCell(g)
	}
	return len(d.dirty) != 0
}

func (d *data) cleanCell(g point) {
	for i := range d.dirty {
		if d.dirty[i].x == g.x && d.dirty[i].y == g.y {
			d.dirty = append(d.dirty[:i], d.dirty[i+1:]...)
			break
		}
	}
}

func (d *data) findNext() point { // greedy
	var ds int
	minDs := 50 * 50
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
