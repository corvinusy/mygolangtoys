package main

import (
	"container/list"
	"fmt"
	"sort"
)

type interval struct {
	start int
	end   int
	count int
}

func main() {
	var trials, size int
	fmt.Scan(&trials)
	input := make([][]int, trials)
	for i := 0; i < trials; i++ {
		fmt.Scanln(&size)
		arr := make([]int, size)
		for k := 0; k < size; k++ {
			fmt.Scan(&arr[k])
		}
		input[i] = arr
	}
	for i := range input {
		fmt.Println(playGame(input[i]))
	}

}

func playGame(a []int) int {
	var left, right *interval
	l := list.New()
	l.PushBack(&interval{0, len(a), 0})
	maxCount := 0
	for e := l.Front(); e != nil; e = e.Next() {
		//fmt.Println(e.Value.(*interval))
		left, right = partition(a, e.Value.(*interval))
		if left != nil {
			if maxCount < left.count {
				maxCount = left.count
			}
			l.PushBack(left)
			l.PushBack(right)
		}
	}
	return maxCount
}

func partition(a []int, iv *interval) (*interval, *interval) {
	if iv.end-iv.start < 2 {
		return nil, nil
	}
	// binary search
	aa := a[iv.start:iv.end]
	i := sort.Search(len(aa),
		func(i int) bool {
			return i != 0 && sum(aa[:i]) >= sum(aa[i:])
		})
	if sum(aa[:i]) == sum(aa[i:]) {
		return &interval{iv.start, iv.start + i, iv.count + 1},
			&interval{iv.start + i, iv.end, iv.count + 1}
	}
	return nil, nil
}

func sum(a []int) int {
	// TODO: implement memoize
	var sum int
	for i := range a {
		sum += a[i]
	}
	return sum
}
