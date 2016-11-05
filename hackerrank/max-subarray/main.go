package main

import (
	"fmt"
)

func main() {
	var (
		trials int
	)
	fmt.Scan(&trials)
	results := make([][2]int, trials)
	for t := 0; t < trials; t++ {
		d := getInput()
		results[t][0] = getMaxCont(d)
		results[t][1] = getMaxAny(d)
	}
	for t := 0; t < trials; t++ {
		fmt.Println(results[t][0], results[t][1])
	}
}

func getInput() []int {
	var n int
	fmt.Scan(&n)
	d := make([]int, n)
	y := make([]interface{}, n)
	for i := range y {
		y[i] = &d[i]
	}
	fmt.Scanln(y...)
	return d
}

func getMaxCont(d []int) int {
	ans := d[0]
	sum := 0
	for i := 0; i < len(d); i++ {
		sum += d[i]
		ans = max(ans, sum)
		sum = max(sum, 0)
	}
	if ans < 0 {
		return maxElem(d)
	}
	return ans
}

func getMaxAny(d []int) int {
	sum := 0
	for i := range d {
		if d[i] > 0 {
			sum += d[i]
		}
	}
	if sum <= 0 {
		return maxElem(d)
	}
	return sum
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func maxElem(d []int) int {
	m := d[0]
	for i := range d {
		if m < d[i] {
			m = d[i]
		}
	}
	return m
}
