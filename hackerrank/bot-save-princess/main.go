package main

import (
	"bufio"
	"fmt"
	"os"
)

type data struct {
	bot   [2]int
	peach [2]int
	found bool
}

func main() {
	var (
		gs int // grid size
		d  data
	)
	fmt.Scan(&gs)
	rd := bufio.NewReader(os.Stdin)
	for i := 0; i < gs; i++ {
		tmpStr, _ := rd.ReadString('\n')
		for k := 0; k < gs; k++ {
			if tmpStr[k] == 'm' {
				d.bot = [2]int{i, k}
			}
			if tmpStr[k] == 'p' {
				d.peach = [2]int{i, k}
			}
		}
	}

	for !d.found {
		d.moveBot()
	}
}

func (d *data) planMove() string {
	dx := d.peach[0] - d.bot[0]
	dy := d.peach[1] - d.bot[1]
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
	d.found = true
	return ""
}

func (d *data) moveBot() {
	pl := d.planMove()
	switch pl {
	case "DOWN":
		d.bot[0]++
		fmt.Println(pl)
	case "UP":
		d.bot[0]--
		fmt.Println(pl)
	case "RIGHT":
		d.bot[1]++
		fmt.Println(pl)
	case "LEFT":
		d.bot[1]--
		fmt.Println(pl)
	default:
		if !d.found {
			panic("No way")
		}
	}
}
