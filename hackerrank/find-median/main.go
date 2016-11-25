package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var s string
	rd := bufio.NewReader(os.Stdin)
	s, _ = rd.ReadString('\n')
	var n int
	fmt.Sscan(s, &n)
	s, _ = rd.ReadString('\n')
	a := make([]int, n)
	ff := make([]interface{}, n)
	for i := range a {
		ff[i] = &a[i]
	}
	fmt.Sscanln(s, ff...)

	fmt.Println(findMedian(a, 0, n-1))
}

func partition(a []int, lo, hi int) int {
	p := a[hi]
	i := lo
	for j := lo; j < hi; j++ {
		if a[j] <= p {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], a[hi] = a[hi], a[i]
	return i
}

func findMedian(a []int, lo, hi int) int {
	p := partition(a, lo, hi)
	ll := len(a) / 2
	for p != ll {
		if p > ll {
			hi = p - 1
		} else {
			lo = p
		}
		p = partition(a, lo, hi)
		fmt.Printf("%d %d %d\n", p, lo, hi)
	}

	return a[p]
}
