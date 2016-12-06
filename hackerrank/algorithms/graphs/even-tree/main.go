package main

import (
	"container/list"
	"fmt"
)

type edge struct {
	from, to int
}

type graph struct {
	edges *list.List
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	var g graph
	g.edges = list.New()
	var e edge
	for i := 0; i < m; i++ {
		fmt.Scan(&e.from, &e.to)
		g.edges.PushBack(e)
	}
	fmt.Println(g.getMaxTreePartitions())
}

func (g *graph) getMaxTreePartitions() int {
	var (
		count int
		left  *graph
		right *graph
		ee    edge
	)
	que := list.New()
	que.PushBack(g)
	for que.Len() != 0 {
		g = que.Front().Value.(*graph)
		for e := g.edges.Front(); e != nil; e = e.Next() {
			ee = e.Value.(edge)
			left, right = g.getPartition(ee)
			if left.edges.Len()%2 == 1 && right.edges.Len()%2 == 1 {
				//left.Print(); right.Print(); fmt.Println(" edge = ", ee)
				count++
				que.PushBack(left)
				que.PushBack(right)
				break
			}
		}
		que.Remove(que.Front())
	}
	return count
}

func (g *graph) getPartition(cut edge) (*graph, *graph) {
	gg := g.deepCopy()
	left := gg.newSubTree(cut, cut.to)
	right := gg.newSubTree(cut, cut.from)
	if left.edges.Len() == 0 || right.edges.Len() == 0 {
		return left, right
	}
	var e edge
	// remove cutting edge
	for el := gg.edges.Front(); el != nil; el = el.Next() {
		e = el.Value.(edge)
		if e.from == cut.from && e.to == cut.to {
			gg.edges.Remove(el)
			break
		}
	}
	// separate edges by sides
	for gg.edges.Len() != 0 {
		for el := gg.edges.Front(); el != nil; el = el.Next() {
			e = el.Value.(edge)
			if left.connectedTo(e.to) || left.connectedTo(e.from) {
				left.edges.PushBack(e)
				gg.edges.Remove(el)
				continue
			}
			if right.connectedTo(e.to) || right.connectedTo(e.from) {
				right.edges.PushBack(e)
				gg.edges.Remove(el)
			}
		}
	}
	return left, right
}

func (g *graph) deepCopy() *graph {
	newGraph := new(graph)
	newGraph.edges = list.New()
	var e edge
	for el := g.edges.Front(); el != nil; el = el.Next() {
		e = el.Value.(edge)
		newGraph.edges.PushBack(e)
	}
	return newGraph
}

func (g *graph) newSubTree(cut edge, node int) *graph {
	newGraph := new(graph)
	newGraph.edges = list.New()
	var e edge
	for el := g.edges.Front(); el != nil; el = el.Next() {
		e = el.Value.(edge)
		if e.from == cut.from && e.to == cut.to { // skip cutting edge
			continue
		}
		if e.to == node || e.from == node {
			newGraph.edges.PushBack(e)
			g.edges.Remove(el)
		}
	}
	return newGraph
}

func (g *graph) connectedTo(node int) bool {
	var e edge
	for el := g.edges.Front(); el != nil; el = el.Next() {
		e = el.Value.(edge)
		if e.to == node || e.from == node {
			return true
		}
	}
	return false
}

func (g *graph) Print() {
	for el := g.edges.Front(); el != nil; el = el.Next() {
		fmt.Print(" ", el.Value.(edge))
	}
	fmt.Print(",")
}
