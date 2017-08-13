package main

import (
	"container/list"
	"fmt"
)

func main() {
	var t int
	for fmt.Scan(&t); t > 0; t-- {
		doRepl()
	}
}

type position struct {
	square, move int
}

func doRepl() {
	// read
	var n, x, y int
	a := make(map[int]int)
	for fmt.Scan(&n); n > 0; n-- {
		fmt.Scan(&x, &y)
		a[x] = y
	}
	for fmt.Scan(&n); n > 0; n-- {
		fmt.Scan(&x, &y)
		a[x] = y
	}
	// evaluate
	result := eval(a, 100)
	// print
	fmt.Println(result)
}

func eval(a map[int]int, d int) int {
	adj := buildAdjacency(a)
	visited := make([]bool, 101)
	finish := 100
	queue := list.New()
	queue.PushBack(position{1, 0}) // zero move: square = start, move = 0
	return doBfs(adj, visited, queue, finish)
}

func buildAdjacency(a map[int]int) [][]int {
	var (
		x  int
		ok bool
	)
	adj := make([][]int, 101)
	for i := 1; i <= 100; i++ {
		adj[i] = []int{}
		for j := 1; i+j <= 100 && j <= 6; j++ {
			if x, ok = a[i+j]; ok {
				adj[i] = append(adj[i], x)
			} else {
				adj[i] = append(adj[i], i+j)
			}
		}
	}
	return adj
}

func doBfs(adj [][]int, visited []bool, queue *list.List, finish int) int {
	if queue.Len() == 0 {
		return -1
	}
	pos := queue.Front().Value.(position)
	sq := pos.square
	if sq == finish {
		return pos.move
	}
	queue.Remove(queue.Front())
	for i := range adj[sq] {
		// push next moves
		if !visited[adj[sq][i]] {
			visited[adj[sq][i]] = true
			queue.PushBack(position{adj[sq][i], pos.move + 1})
		}
	}
	return doBfs(adj, visited, queue, finish)
}
