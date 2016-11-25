package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var (
		s    string
		n, m int
		ok   bool
	)
	// get input
	rd := bufio.NewReader(os.Stdin)
	s, _ = rd.ReadString('\n')
	fmt.Sscan(s, &n)
	a := make([]int, n)
	ff := make([]interface{}, n)
	for i := range ff {
		ff[i] = &a[i]
	}
	s, _ = rd.ReadString('\n')
	fmt.Sscanln(s, ff...)

	s, _ = rd.ReadString('\n')
	fmt.Sscan(s, &m)
	b := make([]int, m)
	ff = make([]interface{}, m)
	for i := range ff {
		ff[i] = &b[i]
	}
	s, _ = rd.ReadString('\n')
	fmt.Sscanln(s, ff...)

	// build map
	bmap := make(map[int]int)
	for i := range b {
		if _, ok = bmap[b[i]]; !ok {
			bmap[b[i]] = 1
		} else {
			bmap[b[i]]++
		}
	}

	// update map
	for i := range a {
		bmap[a[i]]--
	}

	// extract result
	var r []int
	for k := range bmap {
		if bmap[k] != 0 {
			r = append(r, k)
		}
	}

	// sort result
	sort.Ints(r)

	// print
	for i := range r {
		fmt.Printf("%d ", r[i])
	}
	fmt.Println()
}
