package main

import (
    "fmt"
)

const LIMIT = 28123
//const LIMIT = 50

func main() {
	var (
		i, j, pds int
		abundants []int
		absums map[int]bool
	)

	abundants = make([]int, 0)
	absums = make(map[int]bool, LIMIT+1)

	for i = 1; i <= LIMIT; i++ {
		pds = propdivsum(i)
		if pds > i {
			abundants = append(abundants, i)
		}
	}
/*
	for _, a := range abundants {
		fmt.Println(a)
	}
*/
	for _, i = range abundants {
		for _, j = range abundants {
			if i + j <= LIMIT {
				absums[i+j] = true
			}
		}
	}

	sum := 0

	for i = 1; i <= LIMIT; i++ {
		if absums[i] != true {
//			fmt.Printf("%d ", i)
			sum += i
		}
	}

	fmt.Println("\n", sum)

}
/*-----------------------------------------------------------------------------*/

func propdivsum(num int) int {

	var (
		sum int = 1
		acc int = 1
		tmp, i int
	)

	num_stored := num
	
	//trivial case
	if num == 1 {
		return num - num_stored
	}
	// powers of 2
	for num % 2 == 0 {
		acc *= 2
		num /= 2
	}
	acc = acc * 2 - 1
	
	if num == 1 {
		return acc - num_stored
	} else {
		sum = acc
	}
	
	// brute force others
	for i = 3; i * i <= num; i += 2 {
		acc = 1
		tmp = 1
		for num % i == 0 {
			acc = acc * i
			tmp = tmp + acc
			num = num / i
		}
		if (acc > 1) {
			sum = sum * tmp
		}
	}

	if (num > 1) {
		sum = sum * (num + 1)
	}
	return sum - num_stored;
}
