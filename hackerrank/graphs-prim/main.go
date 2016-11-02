package main

import (
	"fmt"
)

const max = 1e5 + 1

func main() {

	// read count
	var (
		ns, vs       int
		row, col, wt int
	)
	fmt.Scanf("%d %d", &ns, &vs)
	g := make([][]int, ns)
	for i := range g {
		g[i] = make([]int, ns)
		for j := range g[i] {
			g[i][j] = max
		}
	}
	// read graph data
	for i := 0; i < vs; i++ {
		fmt.Scanf("%d %d %d", &row, &col, &wt)
		g[row-1][col-1] = wt
		g[col-1][row-1] = wt
	}
	// read start node
	var st int
	fmt.Scanf("%d", &st)

	mst := findMST(g, st)

	fmt.Println(mst)
}

func findMST(g [][]int, st int) int {
	visited := make([]int, len(g))
	nodes := len(g)
	node := 1
	var (
		min        int
		i, j       int
		a, b, u, v int
		mincost    int
	)

	visited[st-1] = 1

	for node < nodes {
		for i, min = 0, max; i < nodes; i++ {
			for j = 0; j < nodes; j++ {
				if g[i][j] < min {
					if visited[i] != 0 {
						min = g[i][j]
						a, b = i, j
						u, v = i, j
					}
					if visited[u] == 0 || visited[v] == 0 {
						node++
						mincost += min
						visited[b] = 1
					}
					g[a][b], g[b][a] = max, max
				}
			}
		}
	}
	return mincost
}
