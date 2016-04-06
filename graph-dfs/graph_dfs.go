package main

import (
    "fmt"
)

const LIMIT = 12

func main() {

	vertex := make([][]int, LIMIT)

	vertex[0] = append(vertex[0], 0, 0, 1)

	for i:=1; i < LIMIT; i++ {
		vertex[i] = append(vertex[i], 0, i-1, i+1)
	}

	vertex[4] = append(vertex[4], 8)
	vertex[7][2] = 6
	vertex[8][1] = 4

	for i, e := range vertex {
		fmt.Println(i, e)
	}

	fmt.Println("Graph initialized\n");

	path := make([]int,0)

	dfs(vertex, path, 0, 0, 11)

}
/*-----------------------------------------------------------------------------*/
func dfs(vertex [][]int, path []int, v, from, finish int) []int {

	if vertex[v][0] != 0 {
		return path
	}

	vertex[v][0] = 1

	path = append(path, v)

	if v == finish {
		fmt.Println("Hooray! The path was found!", path)
		return path
	}

	for i := 1; i < len(vertex[v]); i++ {
		dfs(vertex, path, vertex[v][i], v, finish)
	}

	return nil
	
}
