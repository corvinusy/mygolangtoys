package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	var (
		n int
		s string
	)
	fmt.Scan(&n)
	rd := bufio.NewReader(os.Stdin)
	bigs := make([]string, n)
	for i := range bigs {
		s, _ = rd.ReadString('\n')
		bigs[i] = strings.TrimSpace(s)
	}
	sort.Slice(bigs, func(i, j int) bool {
		if len(bigs[i]) == len(bigs[j]) {
			return strings.Compare(bigs[i], bigs[j]) == -1
		}
		return len(bigs[i]) < len(bigs[j])
	})
	for i := range bigs {
		fmt.Println(bigs[i])
	}
}
