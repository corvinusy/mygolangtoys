package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	var t int
	rd := bufio.NewReader(os.Stdin)
	tstr, _ := rd.ReadString('\n')
	fmt.Sscan(tstr, &t)

	ss := make([][]byte, t)
	for i := range ss {
		s, _ := rd.ReadBytes('\n')
		ss[i] = bytes.TrimSpace(s)
	}

	for i := range ss {
		fmt.Println(anagramCountdown(ss[i]))
	}
}

func anagramCountdown(s []byte) int {
	var count int
	if len(s) == 1 {
		return 0
	}

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] > s[j] {
			count += int(s[i] - s[j])
		} else {
			count += int(s[j] - s[i])
		}
	}
	return count
}
