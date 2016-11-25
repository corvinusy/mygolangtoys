package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var (
		t      int
		s, res string
	)
	rd := bufio.NewReader(os.Stdin)
	s, _ = rd.ReadString('\n')
	fmt.Sscan(s, &t)
	input := make([][2]string, t)
	for i := range input {
		s, _ = rd.ReadString('\n')
		input[i][0] = strings.TrimSpace(s)
		s, _ = rd.ReadString('\n')
		input[i][1] = strings.TrimSpace(s)
	}

	for i := range input {
		amap := reduce(&input[i][0])
		bmap := reduce(&input[i][1])
		res = "NO"
		for k := range amap {
			if bmap[k] {
				res = "YES"
				break
			}
		}
		fmt.Println(res)
	}
}

func reduce(ps *string) map[byte]bool {
	s := *ps
	m := make(map[byte]bool)
	for i := range s {
		m[s[i]] = true
	}
	return m
}
