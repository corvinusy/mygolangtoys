package main

import (
	"container/list"
	"fmt"
)

func main() {
	var n, t int
	fmt.Scan(&n, &t)
	seq := make([]*list.List, n)
	for i := 0; i < n; i++ {
		seq[i] = list.New()
	}
	var qtype, x, y, lastAns int
	for i := 0; i < t; i++ {
		fmt.Scan(&qtype, &x, &y)
		switch qtype {
		case 1:
			query1(seq, x, y, lastAns)
		case 2:
			lastAns = query2(seq, x, y, lastAns)
			fmt.Println(lastAns)
		}
	}
}

func query1(seq []*list.List, x, y, lastAns int) {
	k := (x ^ lastAns) % len(seq)
	seq[k].PushBack(y)
}

func query2(seq []*list.List, x, y, lastAns int) int {
	k := (x ^ lastAns) % len(seq)
	i := y % seq[k].Len()
	var e *list.Element
	for e = seq[k].Front(); i > 0; i-- {
		e = e.Next()
	}
	return e.Value.(int)
}
