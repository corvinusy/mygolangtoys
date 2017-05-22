package main

import (
	"bufio"
	"fmt"
	"os"
)

type road [2]int

func main() {
	var (
		err error
		q   int
	)

	in := bufio.NewReader(os.Stdin)
	_, err = fmt.Fscan(in, &q)
	if err != nil {
		panic(err)
	}

	for ; q > 0; q-- {
		var (
			cNum, rNum   int
			lCost, rCost int
			a, b         int
		)
		// take input
		_, err = fmt.Fscan(in, &cNum, &rNum, &lCost, &rCost)
		if err != nil {
			panic(err)
		}
		roads := make([]road, rNum)
		for ; rNum > 0; rNum-- {
			// read roads and normalize input
			_, err = fmt.Fscan(in, &a, &b)
			if err != nil {
				panic(err)
			}
			if a > b {
				a, b = b, a
			}
			roads[rNum-1] = road{a, b}
		}
		// find and print result
		fmt.Println(findCost(roads, cNum, lCost, rCost))
	}
}

func findCost(roads []road, cNum, lCost, rCost int) int {
	// trivial case
	if lCost <= rCost {
		return lCost * cNum
	}
	// split roads on segments
	segments := splitCitiesToSegments(roads, cNum)
	// accumulate total cost
	total := 0
	for i := range segments {
		total += lCost + rCost*(len(segments[i])-1)
	}
	return total
}

func splitCitiesToSegments(roads []road, cNum int) map[int]map[int]bool {
	var (
		k, r0, r1 int
	)
	segmentOfCity := make([]int, cNum+1) // cities-to-segments tracker
	segments := map[int]map[int]bool{}   // segments
	for i := 1; i <= cNum; i++ {
		segmentOfCity[i] = i
		segments[i] = map[int]bool{}
		segments[i][i] = true
	}
	// cycle
	for i := range roads {
		// found segment connection - move segment 1 to segment 0
		if segmentOfCity[roads[i][0]] != segmentOfCity[roads[i][1]] {
			r0 = segmentOfCity[roads[i][0]]
			r1 = segmentOfCity[roads[i][1]]
			for k = range segments[r1] {
				segments[r0][k] = true
				segmentOfCity[k] = r0
			}
			delete(segments, r1) // delete segment
			//fmt.Printf("%+v\n", segments) // DEBUG
		}
	}
	return segments
}
