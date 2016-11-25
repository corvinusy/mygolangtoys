package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	var s string
	var n int
	s, _ = rd.ReadString('\n')
	fmt.Sscan(s, &n)
	a := make([]int, n)
	b := make([]string, n)
	for i := range a {
		s, _ = rd.ReadString('\n')
		fmt.Sscan(s, &a[i], &b[i])
		if i < len(a)/2 {
			b[i] = "-"
		}
	}

	printSlice(sortCount(a, b))
}

func sortCount(a []int, b []string) []string {
	counts := make([]int, len(a))
	// count quotities
	for i := range a {
		counts[(a[i])]++
	}
	// count sums
	for i := 1; i < len(counts); i++ {
		counts[i] += counts[i-1]
	}
	// prepare output
	out := make([]string, len(b))

	// pose the elements
	for i := len(b) - 1; i >= 0; i-- {
		out[counts[a[i]]-1] = b[i]
		counts[a[i]]--
	}
	return out
}

func printSlice(b []string) {
	for i := 0; i < len(b); i++ {
		fmt.Printf("%s ", b[i])
	}
	fmt.Println()
}
