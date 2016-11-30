package main

import (
	"fmt"
	"sort"
)

type edge struct {
	from, to, wei int
}

func main() {
	// read input
	var nn, ne int
	fmt.Scan(&nn, &ne)

	nodes := make([]int, nn)
	for i := range nodes {
		nodes[i] = -i - 1
	}
	edges := make([]edge, ne)
	for i := range edges {
		fmt.Scan(&edges[i].from, &edges[i].to, &edges[i].wei)
		// normalize edges to zero-indexing values
		edges[i].from--
		edges[i].to--
	}
	mst := findMstKruskal(edges, nodes)

	fmt.Println(mst)
}

// sorting boilerplate
type byWeight []edge

func (a byWeight) Len() int           { return len(a) }
func (a byWeight) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byWeight) Less(i, j int) bool { return a[i].wei < a[j].wei }

var last int

// get node color.
// if nodes[n] < 0 then color = nodes[n]
// else color = nodes[nodes[n]] // recursive
func getColor(nodes []int, n int) int {
	if nodes[n] < 0 {
		last = n
		return nodes[last]
	}
	c := getColor(nodes, nodes[n])
	nodes[n] = last
	return c
}

// worker
func findMstKruskal(edges []edge, nodes []int) int {
	// sort edges
	sort.Sort(byWeight(edges))
	// traversing graph
	var c, dist int
	for i := range edges {
		c = getColor(nodes, edges[i].to)
		if c != getColor(nodes, edges[i].from) {
			nodes[last] = edges[i].to
			// print mst-edges list
			//fmt.Printf("%d %d %d\n", edges[i].from, edges[i].to, edges[i].wei)
			dist += edges[i].wei
		}
	}

	return dist
}
