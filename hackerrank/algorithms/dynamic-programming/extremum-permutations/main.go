package main

import (
	"fmt"

	plib "github.com/fighterlyt/permutation"
)

func main() {
	var n, k, l uint
	fmt.Scan(&n, &k, &l)
	a := make([]uint, k)
	for i := range a {
		fmt.Scan(&a[i])
		a[i]--
	}
	b := make([]uint, l)
	for i := range b {
		fmt.Scan(&b[i])
		b[i]--
	}

	for n = 11; n < 13; n++ {
		c := countPermutations(n, a, b)
		fmt.Println(c, float64(fac(n))/float64(c))
	}
}

func countPermutations(n uint, a, b []uint) uint {
	st := make([]uint, n)
	for i := range st {
		st[i] = uint(i)
	}

	p, _ := plib.NewPerm(st, nil)

	var count uint
	for c, err := p.Next(); err == nil; c, err = p.Next() {
		if isMagic(a, b, c.([]uint)) {
			count++
		}
	}
	return count
}
func fac(n uint) uint {
	p := uint(1)
	for n > 0 {
		p *= n
		n--
	}
	return p
}

func isMagic(a, b, p []uint) bool {
	for _, v := range a {
		if p[v-1] < p[v] || p[v+1] < p[v] {
			return false
		}
	}
	for _, v := range b {
		if p[v-1] > p[v] || p[v+1] > p[v] {
			return false
		}
	}
	return true
}
