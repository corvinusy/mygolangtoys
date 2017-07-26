package main

import (
	"fmt"
	"sort"
)

func main() {
	var s, n, m int
	fmt.Scan(&s, &n, &m)
	kbs := make([]int, n)
	ups := make([]int, m)
	for i := range kbs {
		fmt.Scan(&kbs[i])
	}
	for i := range ups {
		fmt.Scan(&ups[i])
	}
	sort.Ints(kbs)
	sort.Ints(ups)
	fmt.Println(getBest(s, kbs, ups))
}

func getBest(s int, kbs, ups []int) int {
	max := -1
	for i := 0; i < len(kbs) && kbs[i] < s; i++ {
		for j := len(ups) - 1; j >= 0; j-- {
			if kbs[i]+ups[j] > s {
				continue
			}
			if kbs[i]+ups[j] > max {
				max = kbs[i] + ups[j]
				break
			}
		}
	}
	return max
}
