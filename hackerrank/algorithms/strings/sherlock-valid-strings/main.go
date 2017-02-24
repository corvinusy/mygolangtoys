package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

type byRepeats []int // for sorting slice

func (r byRepeats) Len() int {
	return len(r)
}
func (r byRepeats) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
func (r byRepeats) Less(i, j int) bool {
	return countRepeats(r, i) < countRepeats(r, j)
}

func countRepeats(a []int, k int) int {
	c := 0
	for i := range a {
		if a[i] == a[k] {
			c++
		}
	}
	return c
}

func main() {
	rd := bufio.NewReader(os.Stdin)
	b, _ := rd.ReadBytes('\n')
	b = bytes.TrimSpace(b)
	a := make([]int, len(b))
	for i := range b {
		a[i] = int(b[i])
	}
	sort.Ints(a)
	r := getSliceOfRepeats(a)
	sort.Sort(byRepeats(r)) // sort repeats-slice by it own repeats
	if countDelta(r) == 0 { // always valid
		fmt.Println("YES")
		return
	}
	d := countDelta(r[1:])
	if d == 0 && r[0] == 1 { // can remove 1 extra symbol
		fmt.Println("YES")
		return
	}
	if d == 0 && abs(r[0]-r[1]) == 1 { // can "adopt" 1 extra symbol
		fmt.Println("YES")
		return
	}

	fmt.Println("NO")
}

func getSliceOfRepeats(a []int) []int {
	r := make([]int, 1)
	c := a[0]
	for i := 0; i < len(a); i++ {
		if a[i] == c {
			r[len(r)-1]++
		} else {
			c = a[i]
			r = append(r, 1)
		}
	}
	return r
}

func countDelta(a []int) int {
	d := 0
	for i := range a {
		d += abs(a[i] - a[len(a)-1])
	}
	return d
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
