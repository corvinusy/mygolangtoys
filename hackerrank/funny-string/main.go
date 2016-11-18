package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	ts, _ := rd.ReadString('\n')
	var t int
	fmt.Sscan(ts, &t)

	ss := make([][]byte, t)
	for i := range ss {
		s, _ := rd.ReadBytes('\n')
		ss[i] = bytes.TrimSpace(s)
	}

	for i := range ss {
		fmt.Println(isFunny(ss[i]))
	}
}

func isFunny(s []byte) string {
	var d1, d2 int
	if len(s) <= 2 {
		return "Funny"
	}

	for i := 0; i < len(s)-1; i++ {
		d1 = abs(int(s[i+1]) - int(s[i]))
		d2 = abs(int(s[len(s)-i-2]) - int(s[len(s)-i-1]))
		if d1 != d2 {
			return "Not Funny"
		}
	}
	return "Funny"
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
