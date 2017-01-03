package main

import (
	"fmt"
	"sort"
)

type contest struct {
	luck int
	imp  int
}

// sort contests interface
type byValue []contest

func (c byValue) Len() int      { return len(c) }
func (c byValue) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c byValue) Less(i, j int) bool {
	// if importances are equal then more luck go first
	if c[i].imp == c[j].imp {
		return c[i].luck > c[j].luck
	}
	return c[i].imp < c[j].imp // less importance go first
}

func main() {
	var n, k, impCount int
	fmt.Scan(&n, &k)
	cs := make([]contest, n)
	for i := range cs {
		fmt.Scan(&cs[i].luck, &cs[i].imp)
		if cs[i].imp == 1 {
			impCount++
		}
	}
	// sort array by importance and luck, so less valuable elements go to tail
	sort.Sort(byValue(cs))
	// sum lucks for positive part of array
	var sum int
	bdry := n - impCount + k
	for i := 0; i < bdry && i < len(cs); i++ {
		sum += cs[i].luck
	}
	// deduct lucks for negative part of array
	for i := bdry; i < len(cs); i++ {
		sum -= cs[i].luck
	}

	fmt.Println(sum)
}
