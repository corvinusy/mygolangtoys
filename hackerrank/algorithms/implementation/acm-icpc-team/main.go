package main

import (
	"fmt"
)

func main() {
	var (
		n, m, count, cc int
		max, c          []byte
	)
	fmt.Scan(&n, &m)
	a := make([][]byte, n)
	for i := range a {
		a[i] = make([]byte, m)
		fmt.Scanln(&a[i])
		for j := range a[i] {
			a[i][j] -= '0'
		}
	}

	max = byteOr(a[0], a[1])
	count = 0

	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			c = byteOr(a[i], a[j])
			cc = countOnes(c)
			if cc == countOnes(max) {
				count++
			} else if cc > countOnes(max) {
				max = c
				count = 1
			}
		}
	}
	fmt.Printf("%d\n%d\n", countOnes(max), count)
}

func countOnes(a []byte) int {
	ones := 0
	for i := 0; i < len(a); i++ {
		if a[i] == 1 {
			ones++
		}
	}
	return ones
}

func byteOr(a, b []byte) []byte {
	c := make([]byte, len(a))
	for i := range a {
		c[i] = a[i] | b[i]
	}
	return c
}
