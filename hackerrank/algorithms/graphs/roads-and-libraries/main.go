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
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
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
		roadList := list.New()
		for ; rNum > 0; rNum-- {
			// read roads and normalize input
			_, err = fmt.Fscan(in, &a, &b)
			if err != nil {
				panic(err)
			}
			if a > b {
				a, b = b, a
			}
			roadList.PushBack(road{a, b})
		}
		// find and print result
		fmt.Println(findCost(roadList, cNum, lCost, rCost))
	}
}

func findCost(roadList *list.List, cNum, lCost, rCost int) int {
	// trivial case
	if lCost <= rCost {
		return lCost * cNum
	}
	// fill cities list
	cityList := list.New()
	for ; cNum > 0; cNum-- {
		cityList.PushFront(cNum)
	}
	// split roads on segments
	cities := splitRoadsToSegments(roadList, cityList)
	// accumulate total cost
	total := 0
	for i := range cities {
		total += lCost + rCost*(cities[i].Len()-1)
	}
	return total
}

func splitRoadsToSegments(roadList, cityList *list.List) [](*list.List) {
	var segments [](*list.List)
	for e := cityList.Front(); e != nil; e = e.Next() {
		if e.Value.(int) == 0 {
			continue
		}
		segment := list.New()
		segment.PushBack(e.Value.(int))
		e.Value = 0
		// find connected cities (dfs)
		fillConnects(segment, roadList)
		reduceCities(segment, cityList)
		//printList(segment)
		// store segment
		segments = append(segments, segment)
	}
	return segments
}

func fillConnects(c, r *list.List) {
	var v int
	var t *list.Element
	for e := c.Front(); e != nil; e = e.Next() {
		v = e.Value.(int)
		for f := r.Front(); f != nil; f = t {
			t = f.Next()
			if v == f.Value.(road)[0] {
				pushBackUnique(c, f.Value.(road)[1])
				r.Remove(f)
			} else if v == f.Value.(road)[1] {
				pushBackUnique(c, f.Value.(road)[0])
				r.Remove(f)
			}
		}
	}
}

func reduceCities(segment, cityList *list.List) {
	var v int

	for es := segment.Front(); es != nil; es = es.Next() {
		v = es.Value.(int)
		for cs := cityList.Front(); cs != nil; cs = cs.Next() {
			if v == cs.Value.(int) {
				cs.Value = 0
			}
		}
	}
}

func pushBackUnique(c *list.List, k int) {
	for e := c.Front(); e != nil; e = e.Next() {
		if e.Value.(int) == k {
			return
		}
	}
	c.PushBack(k)
}

func printList(l *list.List) {
	fmt.Print("list: ")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v ", e.Value)
	}
	fmt.Println()
}
