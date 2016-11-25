package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {

	rd := bufio.NewReader(os.Stdin)
	a, _ := rd.ReadBytes('\n')
	a = bytes.TrimSpace(a)

	b, _ := rd.ReadBytes('\n')
	b = bytes.TrimSpace(b)

	fmt.Println(toAnagrams(a, b))

}

func toAnagrams(a, b []byte) int {
	var c int
	for i := range a {
		c = bytes.IndexByte(b, a[i])
		if c != -1 {
			a[i] = 0
			b[c] = 0
		}
	}
	return nonZeros(a) + nonZeros(b)
}

func nonZeros(a []byte) int {
	var count int
	for i := range a {
		if a[i] != 0 {
			count++
		}
	}
	return count
}
