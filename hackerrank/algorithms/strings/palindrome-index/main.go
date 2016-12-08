package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		s string
		t int
	)
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		fmt.Scan(&s)
		s = strings.TrimSpace(s)
		if isPalyndrome(s) {
			fmt.Println(-1)
			continue
		}
		idx := indexPalyndrome(s)
		fmt.Println(idx)
	}
}

func isPalyndrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func indexPalyndrome(s string) int {
	var sp string
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			sp = s[:i]
			if i < len(s)-1 {
				sp = sp + s[i+1:]
			}
			if isPalyndrome(sp) {
				return i
			}

			sp = s[:j]
			if j < len(s)-1 {
				sp = sp + s[j+1:]
			}
			if isPalyndrome(sp) {
				return j
			}

			return -1
		}
	}
	return -1

}
