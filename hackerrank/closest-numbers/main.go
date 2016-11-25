package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	sort.Ints(a)
	min := a[1] - a[0]
	var mins []int
	for i := 1; i < len(a); i++ {
		if a[i]-a[i-1] < min {
			min = a[i] - a[i-1]
			mins = []int{}
		}
		if a[i]-a[i-1] == min {
			mins = append(mins, a[i-1], a[i])
		}
	}

	for i := range mins {
		fmt.Printf("%d ", mins[i])
	}
	fmt.Println()
}
