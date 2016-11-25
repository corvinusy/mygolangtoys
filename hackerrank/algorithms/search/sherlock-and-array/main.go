package main

import "fmt"

func main() {
	var trials, size int
	fmt.Scan(&trials)
	data := make([][]int, trials)
	results := make([]bool, trials)
	for t := 0; t < trials; t++ {
		fmt.Scan(&size)
		data[t] = make([]int, size)
		for i := range data[t] {
			fmt.Scan(&data[t][i])
		}
		results[t] = reduce(data[t])
	}
	for t := 0; t < trials; t++ {
		if results[t] {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

const (
	greater = 1
	less    = -1
	equal   = 0
)

func reduce(a []int) bool {
	var cur, prev int
	n := len(a) / 2
	dist := n
	prev = 2 // at first time "prev" != cur
	for dist > 0 {
		if cur != prev {
			prev = cur
			if dist > 1 {
				dist++
			}
			dist = dist >> 1
		}
		//fmt.Print(n, dist)
		cur = paritet(a, n)
		if cur == equal {
			return true
		}
		if cur == greater {
			n -= dist
		} else {
			n += dist
		}
	}
	return false
}

func paritet(a []int, n int) int {
	if len(a) == 1 {
		return equal
	}
	if n >= len(a) {
		return greater
	}
	sum1, sum2 := 0, 0
	for i := 0; i < n; i++ {
		sum1 += a[i]
	}
	for i := n + 1; i < len(a); i++ {
		sum2 += a[i]
	}
	//fmt.Println(": ", sum1, sum2)
	if sum1 == sum2 {
		return equal
	} else if sum1 > sum2 {
		return greater
	}
	return less
}
