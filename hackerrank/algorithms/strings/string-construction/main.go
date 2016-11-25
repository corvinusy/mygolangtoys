package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	var n int
	rd := bufio.NewReader(os.Stdin)
	nstr, _ := rd.ReadString('\n')
	fmt.Sscan(nstr, &n)

	ss := make([][]byte, n)
	for i := range ss {
		s, _ := rd.ReadBytes('\n')
		ss[i] = bytes.TrimSpace(s)
	}

	for i := range ss {
		fmt.Println(construct(ss[i]))
	}
}

func construct(s []byte) int {
	var (
		cost int
		p    []byte
	)

	for i := 0; len(p) < len(s); i++ {
		if !bytes.Contains(p, []byte{s[i]}) {
			cost++
		}
		p = append(p, s[i])
	}
	return cost
}
