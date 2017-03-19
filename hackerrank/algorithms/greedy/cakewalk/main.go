package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n) //nolint
	c := make([]int, n)
	for i := range c {
		fmt.Scan(&c[i]) //nolint
	}
	sort.Ints(c)
	fmt.Println(findMinimumCalories(c))
}

func findMinimumCalories(c []int) int {
	var calories, power int = 0, 1
	for i := 0; i < len(c); i++ {
		calories += power * c[len(c)-1-i]
		power *= 2
	}
	return calories
}
