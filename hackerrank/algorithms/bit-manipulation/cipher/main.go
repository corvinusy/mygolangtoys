package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		k, n int
		s    string
	)
	fmt.Scan(&n, &k)
	fmt.Scan(&s)
	s = strings.TrimSpace(s)
	b := stringToBytes(s)
	c := restoreMessage(b, k, n)
	fmt.Println(bytesToString(c))
}

func stringToBytes(s string) []byte {
	b := make([]byte, len(s))
	for i := range s {
		b[i] = s[i] - '0'
	}
	return b
}

func bytesToString(b []byte) string {
	for i := range b {
		b[i] += '0'
	}
	return string(b)
}

func restoreMessage(b []byte, k, n int) []byte {
	c := make([]byte, n)
	c[0] = b[0]
	for i := 1; i < k && i < n; i++ {
		c[i] = b[i] ^ b[i-1]
	}

	for i := k; i < n; i++ {
		c[i] = b[i] ^ b[i-1] ^ c[i-k]
	}
	return c
}
