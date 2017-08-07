package main

import (
	"fmt"
	"strconv"
)

func main() {
	var p, q int
	fmt.Scan(&p, &q)
	a := findKaprekars(p, q)
	if len(a) == 0 {
		fmt.Print("INVALID RANGE")
		return
	}
	for i := range a {
		fmt.Printf("%d ", a[i])
	}
}

func findKaprekars(start, end int) []int {
	a := make([]int, 0)
	for i := start; i <= end; i++ {
		if isKaprekar(i) {
			a = append(a, i)
		}
	}
	return a
}

func isKaprekar(n int) bool {
	var r, l int
	s := strconv.Itoa(n * n)
	for i := len(s) / 2; i < len(s)/2+1; i++ {
		l, _ = strconv.Atoi(s[i:])
		if l == 0 {
			continue
		}
		r, _ = strconv.Atoi(s[:i])
		if r+l == n {
			return true
		}
	}
	return false
}
