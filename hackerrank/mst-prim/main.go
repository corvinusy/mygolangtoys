package main

import (
	"fmt"
)

const max = 1e6 / 2

func main() {

	// read count
	var (
		ns, vs       int
		row, col, wt int
	)
	fmt.Scan(&ns, &vs)
	g := make([][]int, ns)
	for i := range g {
		g[i] = make([]int, ns)
		for j := range g[i] {
			g[i][j] = max
		}
	}
	// read graph data
	for i := 0; i < vs; i++ {
		fmt.Scan(&row, &col, &wt)
		g[row-1][col-1] = wt
		g[col-1][row-1] = wt
	}
	// read start node
	var st int
	fmt.Scan(&st)

	mst := findMST(g, st-1, ns)

	fmt.Println(mst)
}

func findMST(g [][]int, st, ns int) int {
	dists := make([]int, len(g))
	visited := make([]bool, len(g))

	for i := range dists {
		dists[i] = max
	}
	dists[st] = 0

	for {
		v := -1
		for nv := 0; nv < ns; nv++ { // iterate over all nodes
			if !visited[nv] && (dists[nv] < max) && (v == -1 || dists[nv] < dists[v]) { // not visited && existing && closest node
				v = nv
			}
		}
		if v == -1 {
			break // not found
		}
		visited[v] = true // mark node as always visited
		for nv := 0; nv < ns; nv++ {
			if !visited[nv] && g[v][nv] < max { // for all not visited neighрbors
				if g[v][nv] < dists[nv] {
					dists[nv] = g[v][nv] // update dists values
				}
			}
		}
	}
	ret := 0 // count MST from dists
	for v := 0; v < ns; v++ {
		ret += dists[v]
	}
	return ret
}
