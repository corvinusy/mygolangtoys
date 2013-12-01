package main

import (
    "fmt"
)

const LIMIT = 6000 // maxlen 

func main() {

	cache := prepare_cache() //cache = map of squares

	var l, c int

	count := 0

	for size := 2; ; size++ {

		for a := 1; a <= size; a++ {
			for b := a; b <= size; b++ {
				c = size

				l = c * c + (b+a) * (b+a) // if a <= b <= c

				if cache[l] {
					count++
				}
			}
		}

//		fmt.Println("size =", size, "count=", count)
		if count > 1e6 {
			fmt.Println("result = ", size)
			return
		}
	}
	
}
/*-----------------------------------------------------------------------------*/
func prepare_cache() map[int]bool {

	cache := make(map[int]bool, 0)
	for i := 1; i <= LIMIT; i++ {
		cache[i*i] = true
	}
	return cache
}
