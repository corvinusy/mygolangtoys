package main

import (
	"fmt"
)

type ttrain struct {
	line, start, end int
}

func main() {
	var n, m, k int
	fmt.Scanf("%d %d %d", &n, &m, &k)
	trains := make([]ttrain, k)
	for i := range trains {
		fmt.Scanf("%d %d %d", &trains[i].line, &trains[i].start, &trains[i].end)
		// normalize
		if trains[i].start > trains[i].end {
			trains[i].start, trains[i].end = trains[i].end, trains[i].start
		}
		trains[i].line--
		trains[i].start--
		trains[i].end--
	}
	releaseOverlaps(trains)
	//fmt.Println("debug", trains)
	occupied := getOccupiedPositions(trains)
	fmt.Println(occupied)
	fmt.Println(-occupied + n*m)
}

func getOccupiedPositions(trains []ttrain) int {
	var oc int
	for i := range trains {
		if trains[i].line != -1 {
			oc = oc + (trains[i].end - trains[i].start + 1)
		}
	}
	return oc
}

func releaseOverlaps(trains []ttrain) {
	for i := 0; i < len(trains)-1; i++ {
		//fmt.Println("debug", i, trains)
		if trains[i].line == -1 {
			continue
		}
		for j := range trains {
			if j == i {
				continue
			}
			if trains[j].line == -1 {
				continue
			}
			if trains[j].line == trains[i].line {
				// train-j fully covered by train-i
				if trains[i].start <= trains[j].start && trains[i].end >= trains[j].end {
					trains[j].line = -1
					break
				}
				// train-i fully covered by train-j
				if trains[i].start >= trains[j].start && trains[i].end <= trains[j].end {
					trains[i].line = -1
					break
				}
				// trains overlapped, train-i is left to train-j
				if trains[i].start <= trains[j].start && trains[j].start <= trains[i].end {
					trains[i].end = trains[j].end
					trains[j].line = -1
					i = -1
					break
				}
				// trains overlapped, train-i is right to train-j
				if trains[i].start >= trains[j].start && trains[i].end >= trains[j].start {
					trains[i].start = trains[j].start
					trains[j].line = -1
					i = -1
					break
				}
			}
		}
	}
}
