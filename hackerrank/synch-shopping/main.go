package main

import (
	"fmt"
)

func main() {
	var (
		shopNum, roadNum, sortNum int
		sortLen                   int
		x, y, z                   int
	)
	fmt.Scanln(&shopNum, &roadNum, &sortNum)
	// read sorts
	sorts := make([][]int, sortNum)
	for i := 0; i < shopNum; i++ {
		fmt.Scanf("%d", &sortLen)
		sorts[i] = make([]int, sortLen)
		for k := range sorts[i] {
			fmt.Scanf("%d", &sorts[i][k])
		}
	}
	// read graph
	graph := make([][]int, shopNum)
	for i := range graph {
		graph[i] = make([]int, shopNum)
	}
	for i := 0; i < roadNum; i++ {
		fmt.Scanf("%d %d %d", &x, &y, &z)
		graph[x-1][y-1] = z
		graph[y-1][x-1] = z
	}
	fmt.Printf("%+v\n%+v\n", sorts, graph)
}
