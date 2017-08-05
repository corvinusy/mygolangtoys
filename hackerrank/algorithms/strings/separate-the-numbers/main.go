package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var (
		q int
		s string
	)
	for fmt.Scan(&q); q > 0; q-- {
		fmt.Scan(&s)
		s = strings.TrimSpace(s)
		fmt.Println(solveSeparateNums(s))
	}
}

func solveSeparateNums(s string) string {
	for i := 1; i <= len(s)/2; i++ {
		if checkStep(s, i) {
			return "YES " + s[:i]
		}
	}
	return "NO"
}

func checkStep(s string, step int) bool {
	var (
		cur, i  int
		nextStr string
	)
	for i <= len(s)-step-step {
		if s[i] == '0' {
			return false
		}
		cur, _ = strconv.Atoi(s[i : i+step])
		nextStr = strconv.Itoa(cur + 1)
		if !strings.HasPrefix(s[i+step:], nextStr) {
			return false
		}
		i += step
		step = len(nextStr)
	}

	if i+len(nextStr) != len(s) {
		return false
	}
	return true
}
