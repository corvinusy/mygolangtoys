package main

import (
	"container/list"
	"fmt"
)

type edge struct {
	from, to, wei int
}

func main() {
	// read input
	var q, numNodes, numEdges, start int
	var edges []edge
	var paths map[int]int
	fmt.Scan(&q)
	for i := 0; i < q; i++ {
		fmt.Scan(&numNodes, &numEdges)
		edges = make([]edge, numEdges)
		for j := 0; j < numEdges; j++ {
			fmt.Scan(&edges[j].from, &edges[j].to)
			// normalize input
			edges[j].from--
			edges[j].to--
			edges[j].wei = 6
		}
		fmt.Scan(&start)
		start-- //normalize start
		// get paths
		paths = getPathsBFS(start, edges)
		// output
		for j := 0; j < numNodes; j++ {
			if j == start {
				continue
			}
			if _, ok := paths[j]; ok {
				fmt.Printf("%d ", paths[j])
			} else {
				fmt.Printf("%d ", -1)
			}
		}
		fmt.Println()
	}
}

func getPathsBFS(s int, es []edge) map[int]int {
	dists := make(map[int]int)
	visited := make(map[int]bool)
	dists[s] = 0
	visited[s] = true
	l := list.New()
	l.PushFront(s)
	for l.Len() != 0 {
		node := l.Back().Value.(int)
		l.Remove(l.Back())
		for i := range es {
			if es[i].from == node && !visited[es[i].to] {
				l.PushFront(es[i].to)
				visited[es[i].to] = true
				dists[es[i].to] = dists[es[i].from] + 6
			}
			// back direction check
			if es[i].to == node && !visited[es[i].from] {
				l.PushFront(es[i].from)
				visited[es[i].from] = true
				dists[es[i].from] = dists[es[i].to] + 6
			}

		}
	}
	return dists
}
