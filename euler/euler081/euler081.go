package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
)

const (
	mxSize = 80
	maxInt = int(^uint(0) >> 1)
)

type Node struct {
	value  int
	weight int
}

func main() {

	// get path of binary

	programPath, err := exec.LookPath(os.Args[0])

	if err != nil {
		panic(err)
	}

	dataFileName := path.Dir(programPath) + "/matrix.txt"

	//read file into source

	content, err := ioutil.ReadFile(dataFileName)
	if err != nil {
		panic("File can not be read")
	}

	source := strings.Trim(string(content), "\n")

	// init matrix of Nodes
	mx := make([][]Node, mxSize)

	for i := range mx {
		mx[i] = make([]Node, mxSize)
	}

	// parse file content into matrix

	lines := strings.Split(source, "\n")

	for i, line := range lines {
		numbers := strings.Split(line, ",")
		for j, number := range numbers {
			if value, err := strconv.ParseInt(number, 10, 0); err == nil {
				mx[i][j].value = int(value)
			}
			mx[i][j].weight = 0
		}
	}

	/*
		// init start member weight
		mx[0][0].weight = mx[0][0].value

		// init start member row and col weight
		for i := 1; i < mxSize; i++ {
			mx[0][i] := mx[0][i-1].weight + mx[0][i].value
			mx[i][0] := mx[i-1][0].weight + mx[i][0].value
		}
	*/
	// low row
	for j := mxSize - 2; j >= 0; j-- {
		mx[mxSize-2][j].weight = mx[mxSize-2][j+1].value + mx[mxSize-2][j+1].weight
	}

	// right col
	for i := mxSize - 2; i >= 0; i-- {
		mx[i][mxSize-2].weight = mx[i+1][mxSize-2].value + mx[i+1][mxSize-2].weight
	}

	// other members
	for i := mxSize - 3; i >= 0; i-- {
		for j := mxSize - 3; j >= 0; j-- {

			t1 := mx[i+1][j].value + mx[i+1][j].weight
			t2 := mx[i][j+1].value + mx[i][j+1].weight
			if t1 < t2 {
				mx[i][j].weight += t1
			} else {
				mx[i][j].weight += t2
			}
		}
	}

	fmt.Println(mx[0][0].weight + mx[0][0].value)

}

/*----------------------------------------------------------------------------*/
func dijkstraPath(mx [][]int, i_start, j_start, i_end, j_end int) int {
	return 0
}
