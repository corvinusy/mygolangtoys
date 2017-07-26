package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	cps := make([]int, n)
	for i := range cps {
		fmt.Scan(&cps[i])
	}
	fmt.Println(countLuckyProblems(k, cps))
}

func countLuckyProblems(k int, cps []int) int {
	count := 0
	// organize pages
	pages := make(map[int][]int, 0)
	cur := 0 // current page
	for i := range cps {
		cur++ // new chapter
		pages[cur] = make([]int, 0, k)
		for j := 1; j <= cps[i]; j++ {
			if len(pages[cur]) < k {
				pages[cur] = append(pages[cur], j)
			} else {
				cur++
				pages[cur] = make([]int, 0, k)
				pages[cur] = append(pages[cur], j)
			}
			// count lucky problems
			if cur == j {
				count++
			}
		}
	}
	return count
}
