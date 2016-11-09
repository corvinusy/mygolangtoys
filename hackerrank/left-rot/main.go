package main

import (
	"container/ring"
	"fmt"
)

func main() {
	var n, tmp, rot int
	fmt.Scanln(&n, &rot)
	r := ring.New(n)
	for i := 0; i < n; i++ {
		fmt.Scan(&tmp)
		r.Value = tmp
		r = r.Next()
	}

	for i := 0; i < rot; i++ {
		r = r.Next()
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", r.Value)
		r = r.Next()
	}
}
