package main

import (
	"fmt"
	"math"
)

func main() {
	var n, delta int
	fmt.Scan(&n, &delta)
	var p, c, count int
	t := make(map[int]int)
	for i := 0; i < n-1; i++ {
		fmt.Scan(&p, &c)
		t[c] = p
	}
	fmt.Println(len(t))
	for k := range t {
		for z, ok := t[k]; ok; z, ok = t[z] {
			//fmt.Printf(" %d:%d %d\n", k, z, t[z])
			if int(math.Abs(float64(k-z))) <= delta {
				count++
			}
		}
	}
	fmt.Printf("%d\n", count)
}
