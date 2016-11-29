package main

import "fmt"

func main() {
	var trials int
	var l, r int
	fmt.Scan(&trials)
	results := make([]int, trials)
	for t := 0; t < trials; t++ {
		fmt.Scan(&l, &r)
		results[t] = seqXor(l-1) ^ seqXor(r)
	}
	for t := 0; t < trials; t++ {
		fmt.Println(results[t])
	}
}

func seqXor(r int) int {
	switch r % 8 {
	case 0, 1:
		return r
	case 2, 3:
		return 2
	case 4, 5:
		return r + 2
	case 6, 7:
		return 0
	}
	panic("unreachable")
}
