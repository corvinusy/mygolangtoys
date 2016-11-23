package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var k int

func main() {
	var (
		s string
		n int
	)
	// get input
	rd := bufio.NewReader(os.Stdin)
	s, _ = rd.ReadString('\n')
	fmt.Sscanln(s, &n, &k)
	a := make([]int, n)
	ff := make([]interface{}, n)
	for i := range ff {
		ff[i] = &a[i]
	}
	s, _ = rd.ReadString('\n')
	fmt.Sscanln(s, ff...)

	// remove dupl
	amap := make(map[int]bool)
	for i := range a {
		amap[a[i]] = true
	}
	a = []int{}
	for key := range amap {
		a = append(a, key)
	}
	sort.Ints(a)

	if len(a) < 2 {
		fmt.Println(1)
		return
	}

	var t, count int
	for h := 0; h != -1; h = nextHouse(a, t) {
		t = nextTrans(a, h)
		count++
	}
	// print
	fmt.Println(count)
}

func nextTrans(a []int, idx int) int {
	for i := idx; i < len(a); i++ {
		if a[i]-a[idx] < k {
			continue
		}
		if a[i]-a[idx] == k {
			return i
		}
		if a[i]-a[idx] > k {
			return i - 1
		}
	}
	return idx
}

func nextHouse(a []int, idx int) int {
	for i := idx; i < len(a); i++ {
		if a[i]-a[idx] > k {
			return i
		}
	}
	return -1
}
