package main

import (
	"fmt"
	"strconv"
	"strings"
)

const SIZE = 15

type Node struct {
	row    int
	col    int
	value  int
	weight int
}

func main() {

	var (
		source string = `75
95 64
17 47 82
18 35 87 10
20 04 82 47 65
19 01 23 75 03 34
88 02 77 73 07 63 67
99 65 04 28 06 16 70 92
41 41 26 56 83 40 80 70 33
41 48 72 33 47 32 37 16 94 29
53 71 44 65 25 43 91 52 97 51 14
70 11 33 28 77 73 17 78 39 68 17 57
91 71 52 38 17 14 91 43 58 50 27 29 48
63 66 04 68 89 53 67 30 73 16 69 87 40 31
04 62 98 27 23 09 70 98 73 93 38 53 60 04 23`

		strs    []string
		strnums []string
		tmp     int64
		t1, t2  int
	)

	//create tree and init with 0
	tree := make([][]Node, SIZE)

	for i := range tree {
		tree[i] = make([]Node, SIZE)
		for j, node := range tree[i] {
			node.row = i
			node.col = j
			node.value = 0
			node.weight = 0
		}
	}

	strs = strings.SplitAfter(source, "\n")

	for i, s := range strs {
		strnums = strings.Fields(s)
		for j, sn := range strnums {
			tmp, _ = strconv.ParseInt(sn, 10, 0)
			tree[i][j].value = int(tmp)
		}
	}

	for i := SIZE - 2; i >= 0; i-- {
		for j := 0; j < SIZE-1; j++ {
			t1 = tree[i+1][j].value + tree[i+1][j].weight
			t2 = tree[i+1][j+1].value + tree[i+1][j+1].weight
			if t1 > t2 {
				tree[i][j].weight += t1
			} else {
				tree[i][j].weight += t2
			}
		}
	}

	for i := range tree {
		for j := range tree[i] {
			fmt.Printf("%2d/%3d ", tree[i][j].value, tree[i][j].weight)
		}
		fmt.Println()
	}

	fmt.Println(tree[0][0].weight + tree[0][0].value)

}
