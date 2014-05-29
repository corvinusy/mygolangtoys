package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

// x ** y > a ** b
// y > b * log_x(a)

// log_x(a) = log10(a)/log10(x)

// y * log10(x) > b * log10(a)

const SIZE = 1000

type Node struct {
	row    int
	col    int
	value  int
	weight int
}

func main() {

	var (
		strs    []string
		strnums []string
		tmp     int64
	)

	//read file into source

	content, err := ioutil.ReadFile("base_exp.txt")
	if err != nil {
		panic("File not read")
	}

	source := strings.Trim(string(content), "\n")

	strs = strings.Split(source, "\n")

	pair := make([][2]int64, SIZE)

	for i, s := range strs {
		strnums = strings.Split(s, ",")
		for j, sn := range strnums {
			tmp, _ = strconv.ParseInt(sn, 10, 0)
			pair[i][j] = int64(tmp)
		}
	}
	/*
	       //check data
	   	for i, _ := range pair {
	   		for j, _ := range pair[i] {
	   			fmt.Printf("%7d", pair[i][j])
	   		}
	   		fmt.Println()
	   	}
	*/

	//	nMax := a * log10(b)

	nMax := float64(pair[0][1]) * math.Log(float64(pair[0][0]))

	for i, _ := range pair {
		n := float64(pair[i][1]) * math.Log(float64(pair[i][0]))
		if nMax < n {
			nMax = n
			fmt.Println(i, pair[i][0], pair[0][1])
		}
	}

	return

}
