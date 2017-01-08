package main

import (
	"fmt"
)

func main() {
	var (
		s, t string
		k, r int
	)
	fmt.Scanln(&s)
	fmt.Scanln(&t)
	fmt.Scan(&k)
	m := max(len(s), len(t))
	for i := 0; i < m; i++ {
		if i >= len(t) || i >= len(s) {
			r = abs(len(s) - len(t))
			break
		} else {
			if s[i] != t[i] {
				r = len(s) - i + len(t) - i
				break
			}
		}
	}
	fmt.Println(r)
	switch {
	case r == k,
		k >= len(t)+len(s),
		r < k && r%2 == k%2:
		fmt.Println("Yes")
	default:
		fmt.Println("No")
	}

}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
