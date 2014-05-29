package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Node struct {
	value int
	path  int
}

const SIZE = 80

func main() {

	var (
		strs               []string
		strnums            []string
		tmp                int64
		stable, gen_stable bool
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
		for _, node := range tree[i] {
			node.value = 0
			node.path = 0
		}
	}

	strs = strings.Split(source, "\n")

	for i, s := range strs {
		strnums = strings.Split(s, ",")
		for j, sn := range strnums {
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

	//first column init

	tree[0][0].path = tree[0][0].value

	// left column init

	for i := 1; i < SIZE; i++ {
		tree[i][0].path = tree[i-1][0].path + tree[i][0].value
	}

	for i := range tree {
		for j := 1; j < SIZE; j++ {
			tree[i][j].path = tree[i][j-1].path + tree[i][j].value
		}
	}

	// horizon paths ready

	gen_stable = false

	for !gen_stable {

		// scan from upper row to lower

		gen_stable = true
		stable = false
		for !stable {
			stable = true
			for i := 1; i < SIZE; i++ {
				for j := 0; j < SIZE; j++ {
					if tree[i][j].path > tree[i-1][j].path+tree[i][j].value {
						tree[i][j].path = tree[i-1][j].path + tree[i][j].value
						stable = false
						gen_stable = false
					}
				}
			}
		}

		// scan from left column to right

		stable = false
		for !stable {
			stable = true
			for j := 1; j < SIZE; j++ {
				for i := 0; i < SIZE; i++ {
					if tree[i][j].path > tree[i][j-1].path+tree[i][j].value {
						tree[i][j].path = tree[i][j-1].path + tree[i][j].value
						stable = false
						gen_stable = false
					}
				}
			}
		}

		// scan from lower row to upper

		stable = false
		for !stable {
			stable = true
			for i := SIZE - 2; i >= 0; i-- {
				for j := 0; j < SIZE; j++ {
					if tree[i][j].path > tree[i+1][j].path+tree[i][j].value {
						tree[i][j].path = tree[i+1][j].path + tree[i][j].value
						stable = false
						gen_stable = false
					}
				}
			}
		}

		// scan from right column to left

		stable = false
		for !stable {
			stable = true
			for j := SIZE - 2; j >= 0; j-- {
				for i := 0; i < SIZE; i++ {
					if tree[i][j].path > tree[i][j+1].path+tree[i][j].value {
						tree[i][j].path = tree[i][j+1].path + tree[i][j].value
						stable = false
						gen_stable = false
					}
				}
			}
		}
		fmt.Println(tree[SIZE-1][SIZE-1].path)
	}
	fmt.Println("result= ", tree[SIZE-1][SIZE-1].path)

}

/*-----------------------------------------------------------------------------*/
