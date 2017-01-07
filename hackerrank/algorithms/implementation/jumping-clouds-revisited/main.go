package main

import (
	"container/ring"
	"fmt"
)

func main() {
	var n, k, tmp int
	fmt.Scan(&n, &k)
	clouds := ring.New(n)
	for i := 0; i < n; i++ {
		fmt.Scan(&tmp)
		if tmp == 1 {
			tmp++
		}
		clouds.Value = tmp + 1
		clouds = clouds.Next()
	}

	endPoint := clouds
	energy := 100 - clouds.Value.(int) // first jump included

	for clouds = clouds.Move(k); endPoint != clouds; clouds = clouds.Move(k) {
		energy -= clouds.Value.(int)
	}

	fmt.Println(energy)
}
