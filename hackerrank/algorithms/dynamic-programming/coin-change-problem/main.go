package main

import (
	"fmt"
	"sort"
)

func main() {
	var sum, n int
	fmt.Scan(&sum, &n)
	coins := make([]int, n)
	for i := range coins {
		fmt.Scan(&coins[i])
	}
	sort.Ints(coins)
	result := countWays(coins, sum)
	fmt.Println(result)
}

func countWays(coins []int, sum int) int {
	ways := make([]int, sum+1)
	ways[0] = 1
	for i := range coins {
		for j := coins[i]; j <= sum; j++ {
			ways[j] += ways[j-coins[i]]
		}
	}
	return ways[sum]
}
