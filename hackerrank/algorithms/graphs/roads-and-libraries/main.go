package main

import (
	"bufio"
	"container/list"
	"fmt"
)

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"
)

type road [2]int

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	var (
		err error
		q   int
		f   *os.File
	)

	flag.Parse()
	if *cpuprofile != "" {
		f, err = os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		err = pprof.StartCPUProfile(f)
		if err != nil {
			panic(err)
		}
		defer pprof.StopCPUProfile()
	}

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
		roadMap := map[road]bool{}
		for ; rNum > 0; rNum-- {
			// read roads and normalize input
			_, err = fmt.Fscan(in, &a, &b)
			if err != nil {
				panic(err)
			}
			if a > b {
				a, b = b, a
			}
			roadMap[road{a, b}] = true
		}
		// find and print result
		fmt.Println(findCost(roadMap, cNum, lCost, rCost))
	}
}

func findCost(roadMap map[road]bool, cNum, lCost, rCost int) int {
	// trivial case
	if lCost <= rCost {
		return lCost * cNum
	}
	// fill cities list
	cityMap := map[int]bool{}
	for ; cNum > 0; cNum-- {
		cityMap[cNum] = true
	}
	// split roads on segments
	cities := splitCitiesToSegments(roadMap, cityMap)
	// accumulate total cost
	total := 0
	for i := range cities {
		total += lCost + rCost*(cities[i].Len()-1)
	}
	return total
}

func splitCitiesToSegments(roadMap map[road]bool, cityMap map[int]bool) [](*list.List) {
	var (
		segments [](*list.List)
	)
	for k := range cityMap {
		if !cityMap[k] {
			continue
		}
		segment := list.New()
		segMap := map[int]bool{}
		segment.PushBack(k)
		cityMap[k] = false
		segMap[k] = true
		// find connected cities (dfs)
		fillConnects(segment, roadMap, segMap)
		reduceCities(segMap, cityMap)
		// store segment
		segments = append(segments, segment)

	}
	return segments
}

func fillConnects(seg *list.List, roadMap map[road]bool, segMap map[int]bool) {
	var v int
	for e := seg.Front(); e != nil; e = e.Next() {
		v = e.Value.(int)
		for road := range roadMap {
			if v == road[0] {
				delete(roadMap, road)
				if !segMap[road[1]] {
					seg.PushBack(road[1])
					segMap[road[1]] = true
				}
			} else if v == road[1] {
				delete(roadMap, road)
				if !segMap[road[0]] {
					seg.PushBack(road[0])
					segMap[road[0]] = true
				}
			}
		}
	}
}

func reduceCities(segMap, cityMap map[int]bool) {
	for k := range segMap {
		cityMap[k] = false
	}
}

func printList(l *list.List) {
	fmt.Print("list: ")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v ", e.Value)
	}
	fmt.Println()
}
