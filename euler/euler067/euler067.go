package main

import (
    "fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

const SIZE = 100

type Node struct {
	row int
	col int
	value int
	weight int
}


func main() {

	var (
//		source string
		strs []string
		strnums []string
		tmp int64
		t1, t2 int
	)

	//read file into source

	content, err := ioutil.ReadFile("triangle.txt")
	if err != nil {
		panic("File not read")
	}

	source := string(content)

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

	
	strs = strings.SplitAfter(source,"\n")

	for i, s := range strs {
		strnums = strings.Fields(s)
		for j, sn := range strnums  {
			tmp, _ = strconv.ParseInt(sn, 10, 0)
			tree[i][j].value = int(tmp)
		}
	}
	


	for i:=SIZE-2; i>=0; i-- {
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

	fmt.Println(tree[0][0].weight + tree[0][0].value)

}


