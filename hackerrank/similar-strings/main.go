package main

/*
abcde fghij
giggabaj
68660109
*/

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	nstr, _ := rd.ReadString('\n')
	var n, q int
	fmt.Sscan(nstr, &n, &q)
	s, _ := rd.ReadBytes('\n')
	s = s[:len(s)-1]
	dat := make([][2]int, q)
	for i := range dat {
		qstr, _ := rd.ReadString('\n')
		fmt.Sscan(qstr, &dat[i][0], &dat[i][1])
	}
	for i := range dat {
		fmt.Println(getSimilarCount(s, dat[i][0], dat[i][1]))
	}
}

func getSimilarCount(s []byte, start, end int) int {
	ll := end - start
	var count int
	for i := 0; i < len(s)-ll; i++ {
		if i == start-1 {
			count++
			continue
		}
		if isSimilar(s[start-1:end], s[i:i+ll+1]) {
			count++
		}
	}
	return count
}

func isSimilar(s1, s2 []byte) bool {
	for i := range s1 {
		for j := i + 1; j < len(s1); j++ {
			if s1[i] != s1[j] && s2[i] != s2[j] {
				continue
			} else if s1[i] == s1[j] && s2[i] == s2[j] {
				continue
			} else {
				return false
			}
		}
	}
	return true
}
