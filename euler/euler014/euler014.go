package main

import (
	"fmt"
)

func main() {

	var i, n, cur_i, cur_count, count int64

	cur_i = 1
	cur_count = 0

	for i = 2; i < 1e6; i++ {
		count = 0
		n = i
		for n != 1 {
			if n%2 == 0 {
				n /= 2
			} else {
				n = 3*n + 1
			}
			count++
		}
		if cur_count < count {
			cur_count = count
			cur_i = i
		}
	}
	fmt.Println(cur_i)
}
