package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	var (
		s    string
		t, n int
		b    []byte
	)
	s, _ = rd.ReadString('\n')
	fmt.Sscan(s, &t)
	input := make([][][]int, t)
	for i := range input {
		s, _ = rd.ReadString('\n')
		fmt.Sscan(s, &n)
		input[i] = make([][]int, n)
		for j := range input[i] {
			b, _ = rd.ReadBytes('\n')
			b = bytes.TrimSpace(b)
			input[i][j] = make([]int, n)
			for k := range b {
				input[i][j][k] = int(b[k])
			}

		}
	}
	// sort grid
	for i := range input {
		for j := range input[i] {
			sort.Ints(input[i][j])
		}
	}
	// check condition
	for i := range input {
		if isSortable(input[i]) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}

}

func isSortable(a [][]int) bool {
	for i := 0; i < len(a)-1; i++ {
		for j := range a[i] {
			if a[i][j] > a[i+1][j] {
				return false
			}
		}
	}
	return true
}
