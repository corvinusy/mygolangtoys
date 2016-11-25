package main

import (
	"container/list"
	"fmt"
)

var n, qswaps, iswaps int

func main() {
	fmt.Scan(&n)
	d := make([]int, n)
	l := list.New()
	for i := range d {
		fmt.Scan(&d[i])
		l.PushBack(d[i])
	}

	insSort(l)
	quickSort(d, 0, n-1)
	fmt.Println(iswaps - qswaps)
}

func quickSort(a []int, lo, hi int) {
	if lo < hi {
		p := partition(a, lo, hi)
		quickSort(a, lo, p-1)
		quickSort(a, p+1, hi)
	}
}

func partition(a []int, lo, hi int) int {
	p := a[hi]
	i := lo
	for j := lo; j < hi; j++ {
		if a[j] <= p {
			a[i], a[j] = a[j], a[i]
			qswaps++
			i++
		}
	}
	a[i], a[hi] = a[hi], a[i]
	qswaps++
	return i
}

func insSort(l *list.List) {
	for e := l.Front(); e.Next() != nil; e = e.Next() {
		if e.Next().Value.(int) < e.Value.(int) {
			insSortElem(e.Next())
		}
	}
}

func insSortElem(elem *list.Element) {
	value := elem.Value.(int)
	for e := elem; e != nil; e = e.Prev() {
		if e.Prev() != nil && e.Prev().Value.(int) > value {
			e.Value = e.Prev().Value
			iswaps++
		} else {
			e.Value = value
			break
		}
	}
}

/*
func printArray(a []int) {
	for i := range a {
		fmt.Printf("%d ", a[i])
	}
	fmt.Println()
}
*/
