package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	const maxLen = 100
	var n int
	rd := bufio.NewReader(os.Stdin)
	ns, _ := rd.ReadString('\n')
	fmt.Sscan(ns, &n)

	rs := make([]string, n)
	minLen := maxLen + 1
	minInd := 0
	for i := range rs {
		s, _ := rd.ReadString('\n')
		rs[i] = strings.TrimSpace(s)
		if minLen > len(rs[i]) {
			minLen = len(rs[i])
			minInd = i
		}
	}

	resMap := make(map[string]bool)

	var f bool
	for i := range rs[minInd] {
		f = true
		for k := range rs {
			if k == minInd {
				continue
			}
			if !strings.Contains(rs[k], string(rs[minInd][i])) {
				f = false
				break
			}
		}
		if f {
			resMap[string(rs[minInd][i])] = true
		}
	}

	fmt.Println(len(resMap))
}
