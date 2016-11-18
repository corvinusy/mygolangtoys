package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	rd := bufio.NewReader(os.Stdin)
	s, _ := rd.ReadBytes('\n')
	s = s[:len(s)-1] // trimRight '\r'
	ns, _ := rd.ReadString('\n')
	fmt.Sscan(ns, &n)
	fmt.Println(len(s), s, n)

	a := countA(s)
	if a == 0 {
		fmt.Println(0)
		return
	}

	rem := s[:n%len(s)]
	as := a*(n/len(s)) + countA(rem)
	fmt.Println(as)
}

func countA(s []byte) int {
	var count int
	for i := range s {
		if s[i] == 'a' {
			count++
		}
	}
	return count
}
