package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	// read input
	rd := bufio.NewReader(os.Stdin)
	s, _ := rd.ReadString('\n')
	r := strings.NewReplacer(" ", "")
	s = r.Replace(strings.TrimSpace(s))
	// perform encryption
	ss := performEncryption(s)
	for i := range ss {
		fmt.Print(ss[i] + " ")
	}
}

func performEncryption(s string) []string {
	sqlen := math.Sqrt(float64(len(s)))
	cols := int(math.Ceil(sqlen))
	scol := make([]string, cols)
	for i := 0; i < len(s); i++ {
		scol[i%cols] += string(s[i])
	}
	return scol
}
