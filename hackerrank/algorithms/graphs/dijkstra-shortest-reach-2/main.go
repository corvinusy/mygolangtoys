package main

import (
	"bufio"
	"fmt"
	"os"
)

type edge struct {
	from, to, wei int
}

const maxDistance = 1 << 31 //(2^31)

func main() {
	// read input
	scanner := bufio.NewScanner(os.Stdin)
	var q, numNodes, numEdges, start int
	var edges []edge
	var paths map[int]int
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d", &q)
	for i := 0; i < q; i++ {
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &numNodes, &numEdges)
		edges = make([]edge, numEdges)
		for j := 0; j < numEdges; j++ {
			scanner.Scan()
			edges[j].from, edges[j].to, edges[j].wei = toInt3(scanner.Bytes())
			// normalize input
			edges[j].from--
			edges[j].to--
		}
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &start)
		start-- //normalize start
		// get paths
		paths = getPathsDijkstra(start, edges)

		// output
		for j := 0; j < numNodes; j++ {
			if j == start {
				continue
			}
			if paths[j] < maxDistance {
				fmt.Printf("%d ", paths[j])
			} else {
				fmt.Printf("%d ", -1)
			}
		}
		fmt.Println()
	}
}

func getPathsDijkstra(s int, es []edge) map[int]int {
	var u int
	dists := make(map[int]int)
	nodes := make(map[int]struct{})
	neis := make(map[int]map[edge]struct{})
	// Init all distances to max
	for i := range es {
		dists[es[i].from] = maxDistance
		nodes[es[i].from] = struct{}{}
		if neis[es[i].from] == nil {
			neis[es[i].from] = make(map[edge]struct{})
		}
		neis[es[i].from][es[i]] = struct{}{}
		// back direction
		dists[es[i].to] = maxDistance
		nodes[es[i].to] = struct{}{}
		if neis[es[i].to] == nil {
			neis[es[i].to] = make(map[edge]struct{})
		}
		neis[es[i].to][es[i]] = struct{}{}
	}
	dists[s] = 0

	for len(nodes) != 0 {
		// find node with minimum distance
		min := maxDistance
		for i := range nodes {
			if dists[i] <= min {
				min = dists[i]
				u = i
			}
		}
		// remove node with minimum distance
		delete(nodes, u)
		// define minimum edge distances
		for k := range neis[u] {
			if k.from == u {
				if dists[k.to] > dists[u]+k.wei {
					dists[k.to] = dists[u] + k.wei
				}
			}
			if k.to == u {
				if dists[k.from] > dists[u]+k.wei {
					dists[k.from] = dists[u] + k.wei
				}
			}
		}
	}

	return dists
}

func toInt3(buf []byte) (int, int, int) {
	var a [3]int
	i := 0
	for _, v := range buf {
		if v != ' ' {
			a[i] = a[i]*10 + int(v-'0')
		} else {
			i++
		}
	}
	return a[0], a[1], a[2]
}
