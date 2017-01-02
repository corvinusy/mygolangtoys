package main

import (
	"container/list"
	"fmt"
)

type point struct {
	x, y int
}

func main() {
	var n, m, c int
	fmt.Scan(&n, &m)
	g := list.New()
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&c)
			if c == 1 {
				g.PushBack(point{i, j})
			}
		}
	}
	//printPoints(g)
	res := max(findForestSizes(g))
	fmt.Println(res)
}

func max(a *list.List) int {
	if a.Len() == 0 {
		return -1 << 63 // math.MinInt64
	}
	mx := a.Front().Value.(int)
	for e := a.Front(); e != nil; e = e.Next() {
		if e.Value.(int) > mx {
			mx = e.Value.(int)
		}
	}
	return mx
}

func findForestSizes(g *list.List) *list.List {
	forest := makeForest(g)
	//printForest(forest)
	sizes := list.New()
	for e := forest.Front(); e != nil; e = e.Next() {
		sizes.PushBack(e.Value.(*list.List).Len())
	}
	return sizes
}

func makeForest(g *list.List) *list.List {
	var neighbors, rc *list.List
	forest := list.New()
	for e := g.Front(); e != nil; e = e.Next() {
		p := list.New()
		p.PushBack(e.Value)
		forest.PushBack(p)
	}
	for e := g.Front(); e != nil; e = e.Next() {
		neighbors = findNeighbors(forest, e.Value.(point))
		if neighbors.Len() != 0 {
			// unite
			rc = list.New()
			for r := neighbors.Front(); r != nil; r = r.Next() {
				rc.PushBackList(r.Value.(*list.List))
				removeNeighbor(forest, r.Value.(*list.List))
			}
			// clear forest
			forest.PushBack(rc)
		}
	}
	return forest
}

func removeNeighbor(forest, neighbor *list.List) {
	for f := forest.Front(); f != nil; f = f.Next() {
		if isListsEqual(f.Value.(*list.List), neighbor) {
			forest.Remove(f)
		}
	}
}

func isListsEqual(a, b *list.List) bool {
	if a.Len() != b.Len() {
		return false
	}
	for e, f := a.Front(), b.Front(); e != nil && f != nil; e, f = e.Next(), f.Next() {
		if e.Value != f.Value {
			return false
		}
	}
	return true
}

func findNeighbors(forest *list.List, c point) *list.List {
	neighbors := list.New()
	for f := forest.Front(); f != nil; f = f.Next() {
		for e := f.Value.(*list.List).Front(); e != nil; e = e.Next() {
			if abs(e.Value.(point).x-c.x) <= 1 && abs(e.Value.(point).y-c.y) <= 1 {
				neighbors.PushBack(f.Value.(*list.List))
				break
			}
		}
	}
	return neighbors
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

/*
func printForest(forest *list.List) {
	for f := forest.Front(); f != nil; f = f.Next() {
		printPoints(f.Value.(*list.List))
		fmt.Print(";")
	}
	fmt.Println()
}

func printPoints(g *list.List) {
	for e := g.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value.(point), " ")
	}
}
*/
