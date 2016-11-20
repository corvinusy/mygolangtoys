package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
)

func main() {

	rd := bufio.NewReader(os.Stdin)
	s, _ := rd.ReadString('\n')
	var t int
	fmt.Sscan(s, &t)
	ss := make([][]byte, t)
	var b []byte
	for i := range ss {
		b, _ = rd.ReadBytes('\n')
		ss[i] = bytes.TrimSpace(b)
	}

	for i := range ss {
		fmt.Println(getCountToAnagram(ss[i]))
	}
}

func getCountToAnagram(s []byte) int {
	if len(s)%2 != 0 {
		return -1
	}

	var ok bool
	m := make(map[byte]int)
	for i := range s {
		if _, ok = m[s[i]]; !ok {
			m[s[i]] = 0
		}
		if i < len(s)/2 {
			m[s[i]]++
		} else {
			m[s[i]]--
		}
	}
	var count int
	for i := range m {
		count += int(math.Abs(float64(m[i])))
	}
	return count / 2
}
