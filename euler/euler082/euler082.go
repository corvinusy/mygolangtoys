package main

import (
    "fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

const SIZE = 80

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

/*
  //check data
	for i, _ := range tree {
		for j, _ := range tree[i] {
			fmt.Printf("%5d", tree[i][j].value)
		}
		fmt.Println()
	}
*/	
	fmt.Println("start")

	for i := SIZE-11; i <= SIZE-1; i++ {
		for j := SIZE-11; j <= SIZE-1; j++ {
			fmt.Printf("%6d", tree[i][j].value)
		}
		fmt.Println()
	}


	// right col
	for i := SIZE-1; i >= 0; i-- {
		tree[i][SIZE-1].weight = tree[i][SIZE-1].value
	}

	// init all weights going by columns from right to left 
	for j := SIZE-2; j >= 0; j-- {
		for i := 0; i <= SIZE-1; i++ {
			tree[i][j].weight = tree[i][j].value + tree[i][j+1].weight
		}
		// correct weights for downward
		stable := false
		for !stable {
			stable = true

			for i := SIZE-2; i >= 0; i-- {
				// if down is better than replace
				downwt := tree[i+1][j].weight + tree[i][j].value
				if tree[i][j].weight > downwt {
					tree[i][j].weight = downwt
					stable = false
				}
			}
		}
		// correct weights for upward
		stable = false

		for !stable {
			stable = true

			for i := 1; i <= SIZE-1; i++ {
				// if up is better than replace
				upwt := tree[i-1][j].weight + tree[i][j].value
				if tree[i][j].weight > upwt {
					tree[i][j].weight = upwt
					stable = false
				}
			}
		}
	}

	// find result in first column
	result := tree[0][0].weight
	for i := 0; i <= SIZE-1; i++ {
		if result > tree[i][0].weight {
			result = tree[i][0].weight
		}
	}

	fmt.Println(result)

}
/*-----------------------------------------------------------------------------*/
