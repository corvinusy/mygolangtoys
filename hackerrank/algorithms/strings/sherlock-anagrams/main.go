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
	for ; t > 0; t-- {
		fmt.Scanln(&s)
		s = strings.TrimSpace(s)
		fmt.Println(countAnagrams(s))
	}
}

func countAnagrams(s string) int {
	var count int
	for slen := 1; slen < len(s); slen++ {
		for i := 0; i < len(s)-slen; i++ {
			for j := i + 1; j <= len(s)-slen; j++ {
				if isAnagram(s[i:i+slen], s[j:j+slen]) {
					count++
				}
			}
		}
	}
	return count
}

func isAnagram(s1, s2 string) bool {
	var abc [26]int
	for i := range s1 {
		abc[s1[i]-'a']++
	}
	for i := range s2 {
		abc[s2[i]-'a']--
	}
	for i := range abc {
		if abc[i] != 0 {
			return false
		}
	}
	return true
}
