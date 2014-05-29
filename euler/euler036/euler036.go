package main

import (
	"fmt"
	"strconv"
)

func main() {

	const LIMIT = 1e6

	var (
		i, finsum int64
		s1, s2    string
	)

	numslice := make([]int64, 0)

	finsum = 0

	for i = 1; i <= LIMIT; i++ {
		s1 = strconv.FormatInt(i, 10)
		s2 = strconv.FormatInt(i, 2)

		if is_palindrome(s1) && is_palindrome(s2) {
			numslice = append(numslice, i)
			finsum += i
		}
	}
	fmt.Println(numslice, finsum)
}

/*-----------------------------------------------------------------------------*/
func is_palindrome(s string) bool {

	last := len(s) - 1
	i := 0

	for i < last-i {
		if s[i] != s[last-i] {
			return false
		}
		i++

	}
	return true
}
