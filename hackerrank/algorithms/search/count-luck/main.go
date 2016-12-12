package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type point struct {
	x, y int
}

type path []point

func main() {
	var (
		t int
		s string
	)

	rd := bufio.NewReader(os.Stdin)
	s, _ = rd.ReadString('\n')
	fmt.Sscanf(s, "%d", &t)
	for k := 0; k < t; k++ {
		var n, m, guess int
		s, _ = rd.ReadString('\n')
		fmt.Sscanf(s, "%d %d", &n, &m)
		// read scheme
		var start, end point
		g := make([][]int, n)
		for i := range g {
			s, _ = rd.ReadString('\n')
			s = strings.TrimSpace(s)
			g[i] = make([]int, m)
			for j := range g[i] {
				switch s[j] {
				case 'X':
					g[i][j] = 1
				case '*':
					end = point{i, j}
				case 'M':
					start = point{i, j}
				}
			}
		}
		// read guess
		s, _ = rd.ReadString('\n')
		fmt.Sscanf(s, "%d", &guess)
		// worker
		if guess == waveCount(g, start, end) {
			fmt.Println("Impressed")
		} else {
			fmt.Println("Oops!")
		}
	}
	return
}

func waveCount(g [][]int, start, end point) int {
	var count int
	pa := findPath(g, start, end)
	if pa == nil {
		return -1
	}
	for i := range pa {
		fmt.Print(" ", pa[i])
		if isAmbigous(g, start, pa[i]) {
			fmt.Print("*")
			count++
		}
	}
	fmt.Println("\n", count)
	return count
}

func planMove(g [][]int, p point) path {
	var pa path
	n := len(g)
	m := len(g[0])
	if p.x < n-1 && g[p.x+1][p.y] == 0 {
		pa = append(pa, point{p.x + 1, p.y})
	}
	if p.x > 0 && g[p.x-1][p.y] == 0 {
		pa = append(pa, point{p.x - 1, p.y})
	}
	if p.y < m-1 && g[p.x][p.y+1] == 0 {
		pa = append(pa, point{p.x, p.y + 1})
	}
	if p.y > 0 && g[p.x][p.y-1] == 0 {
		pa = append(pa, point{p.x, p.y - 1})
	}
	return pa
}

func isAmbigous(g [][]int, start, p point) bool {
	var dof int
	n := len(g)
	m := len(g[0])
	if p.x < n-1 && g[p.x+1][p.y] != 1 {
		dof++
	}
	if p.x > 0 && g[p.x-1][p.y] != 1 {
		dof++
	}
	if p.y < m-1 && g[p.x][p.y+1] != 1 {
		dof++
	}
	if p.y > 0 && g[p.x][p.y-1] != 1 {
		dof++
	}
	if p == start {
		return dof > 1
	}
	return dof > 2
}

type ant struct {
	pos   point
	track []point
	alive bool
}

func findPath(g [][]int, start, end point) path {
	var pa path
	// create ant
	ants := []ant{{pos: start, alive: true}}
	// walk ant
	for {
		for i := range ants {
			if !ants[i].alive {
				continue
			}
			if ants[i].pos == end {
				return ants[i].track
			}
			ants[i].track = append(ants[i].track, ants[i].pos) // add pos to track
			g[ants[i].pos.x][ants[i].pos.y] = -1               // mark pos
			pa = planMove(g, ants[i].pos)
			switch len(pa) {
			case 0:
				// kill ant
				ants[i].alive = false
			case 1:
				// make move
				ants[i].pos = pa[0]
			case 2, 3, 4:
				// fork
				ants[i].pos = pa[0]
				for j := 1; j < len(pa); j++ {
					a := ant{pos: pa[j], alive: true}
					a.track = append(a.track, ants[i].track...)
					ants = append(ants, a)
				}
			}
		}
	}

}
