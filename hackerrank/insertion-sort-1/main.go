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

	sorting := tmp

	for e := l.Back(); e != nil; e = e.Prev() {
		if e.Prev() != nil && e.Prev().Value.(int) > sorting {
			e.Value = e.Prev().Value
			printList(l)
		} else {
			e.Value = sorting
			break
		}
	}
	printList(l)
}

func printList(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%d ", e.Value)
	}
	fmt.Println()
}
