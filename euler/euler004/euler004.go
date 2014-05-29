package main

import (
	"fmt"
	"strconv"
)

func main() {

	for i := 999; i > 900; i-- {
		for j := 999; j > 900; j-- {
			if is_palindrome(i * j) {
				fmt.Println(i * j)
				break
			}
		}
	}
}

func is_palindrome(p int) bool {

	s := strconv.Itoa(p)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
