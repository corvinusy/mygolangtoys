package main

import (
	"bytes"
	"fmt"
)

func main() {
	var (
		n int
		s string
	)
	fmt.Scan(&n)
	fmt.Scanln(&s)
	a := alphabet(s)
	max, r := 0, 0
	for c1 := range a {
		for c2 := range a {
			if c1 != c2 {
				r = reduce(c1, c2, s)
				if max < r {
					max = r
				}
			}
		}
	}
	fmt.Println(max)
}

func alphabet(s string) map[byte]bool {
	a := make(map[byte]bool, 0)
	var ok bool
	for i := range s {
		if _, ok = a[s[i]]; !ok {
			a[s[i]] = true
		}
	}
	return a
}

func reduce(c1, c2 byte, s string) int {
	var (
		buf  bytes.Buffer
		prev byte
	)
	for i := range s {
		if s[i] == c1 || s[i] == c2 {
			if s[i] != prev {
				buf.WriteByte(s[i])
				prev = s[i]
			} else {
				return 0
			}
		}
	}
	return buf.Len()
}
