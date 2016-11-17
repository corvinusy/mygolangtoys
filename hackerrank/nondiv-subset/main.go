package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	m := make(map[int]int, k)
	for i := range a {
		m[a[i]%k]++
	}

	sum := 0
	if m[0] != 0 {
		sum++
	}
	if k%2 == 0 {
		sum++
	}

	for i, j := 1, k-1; i < j; i, j = i+1, j-1 {
		if m[i] >= m[j] {
			sum += m[i]
		} else {
			sum += m[j]
		}
	}

	fmt.Println(sum)
}
