package main

import (
    "fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

const SIZE = 81

type Node struct {
	row int
	col int
	value int
	weight int
}


func main() {

	var (
		strs []string
		strnums []string
		tmp int64
	)

	//read file into source

	content, err := ioutil.ReadFile("matrix.txt")
	if err != nil {
		panic("File not read")
	}

	source := strings.Trim(string(content), "\n")

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

	strs = strings.Split(source,"\n")

	for i, s := range strs {
		strnums = strings.Split(s, ",")
		for j, sn := range strnums  {
			tmp, _ = strconv.ParseInt(sn, 10, 0)
			tree[i][j].value = int(tmp)
		}
	}

/*  //check data
	for i, _ := range tree {
		for j, _ := range tree[i] {
			fmt.Printf("%5d", tree[i][j].value)
		}
		fmt.Println()
	}
*/	

	// low row
	for j := SIZE-2; j >= 0; j-- {
		tree[SIZE-2][j].weight = tree[SIZE-2][j+1].value + tree[SIZE-2][j+1].weight
	}

	// right col
	for i := SIZE-2; i >= 0; i-- {
		tree[i][SIZE-2].weight = tree[i+1][SIZE-2].value + tree[i+1][SIZE-2].weight
	}


	// other members
	for i := SIZE-3; i >= 0; i-- {
		for j := SIZE-3; j >= 0; j-- {

			t1 := tree[i+1][j].value + tree[i+1][j].weight
			t2 := tree[i][j+1].value + tree[i][j+1].weight
			if t1 < t2 {
				tree[i][j].weight += t1
			} else {
				tree[i][j].weight += t2
			}
		}
	}

	fmt.Println(tree[0][0].weight + tree[0][0].value)

}
