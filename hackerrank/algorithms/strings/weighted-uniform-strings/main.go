package main

import "fmt"

func main() {
	var (
		s    string
		q, k int
	)
	fmt.Scan(&s)
	weights := formWeights(s)
	for fmt.Scan(&q); q > 0; q-- {
		fmt.Scan(&k)
		if weights[k] {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

func formWeights(s string) map[int]bool {
	var w int
	weights := make(map[int]bool)
	// init first
	k := int(s[0] - 'a' + 1)
	weights[k] = true
	last := s[0]
	for i := 1; i < len(s); i++ {
		w = int(s[i] - 'a' + 1)
		if s[i] == last {
			k += w
		} else {
			k = w
		}
		weights[k] = true
		last = s[i]
	}
	return weights
}
