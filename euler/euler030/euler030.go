package main

import (
    "fmt"
)

func main() {
	
	const LIMIT = 2e5

	var (
		i, d, sum, num, finsum  int64
	)

	numslice := make([]int64, 0)
	finsum = 0

	for i = 2; i <= LIMIT; i++ {
		num = i
		sum = 0
		for num > 0  {
			d = num % 10
			sum = sum + d*d*d*d*d
			if sum > i {
				break
			}
			num = num / 10
		}

		if i == sum {
			numslice = append(numslice, i)
			finsum += i
		}
	}
	fmt.Println(numslice, finsum)
}
