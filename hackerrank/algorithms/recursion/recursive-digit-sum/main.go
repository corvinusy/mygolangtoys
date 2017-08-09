package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		s string
		k int
	)
	fmt.Scan(&s, &k)
	fmt.Println(superdigit(s, k))
}

func superdigit(s string, k int) string {
	if k*len(s) == 1 {
		return s
	}
	var n int
	for i := range s {
		n += int(s[i] - '0')
	}
	n *= k
	return superdigit(strconv.Itoa(n), 1)
}
