package main

import (
	"fmt"
	"sort"
)

var cache map[int]map[int]int

func main() {
	var sum, n int
	fmt.Scan(&sum, &n)
	coins := make([]int, n)
	for i := range coins {
		fmt.Scan(&coins[i])
	}
	sort.Ints(coins)
	result := countWays(coins, sum)
	// cache = make(map[int]map[int]int)
	// result := countWays2(coins, sum, 0)
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

func countWays2(coins []int, sum, cur int) int {
	if _, ok := cache[sum][cur]; ok {
		return cache[sum][cur]
	}
	cache[sum] = make(map[int]int)
	var accu int

	for i := cur; i < len(coins); i++ {
		if sum-coins[i] > 0 {
			accu += countWays2(coins, sum-coins[i], i)
		} else if sum-coins[i] == 0 {
			accu++
		}
	}
	cache[sum][cur] = accu
	return accu
}
