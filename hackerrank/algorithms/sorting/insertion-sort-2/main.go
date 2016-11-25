package main

import (
	"container/list"
	"fmt"
)

func main() {
	var n, tmp int
	fmt.Scan(&n)
	l := list.New()
	for i := 0; i < n; i++ {
		fmt.Scan(&tmp)
		l.PushBack(tmp)
	}

	for e := l.Front(); e.Next() != nil; e = e.Next() {
		if e.Next().Value.(int) < e.Value.(int) {
			insSortElem(e.Next())
		}
		printList(l)
	}
}

func insSortElem(elem *list.Element) {
	value := elem.Value.(int)
	for e := elem; e != nil; e = e.Prev() {
		if e.Prev() != nil && e.Prev().Value.(int) > value {
			e.Value = e.Prev().Value
		} else {
			e.Value = value
			break
		}
	}

}

func printList(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%d ", e.Value)
	}
	fmt.Println()
}
