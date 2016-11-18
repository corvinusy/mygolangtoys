package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t, dels int
	fmt.Scan(&t)
	ss := make([][]byte, t)
	rd := bufio.NewReader(os.Stdin)
	for i := range ss {
		ss[i], _ = rd.ReadBytes('\n')
	}

	for i := range ss {
		dels = countDels(ss[i])
		fmt.Println(dels)
	}
}

func countDels(s []byte) int {
	var count int
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
			s = append(s[:i], s[i+1:]...)
			i--
		}
	}
	return count
}
