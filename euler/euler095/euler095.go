package main

import (
    "fmt"
)

const LIMIT = 1e6

const CHAINLIMIT = 100

func main() {

	var (
		n, i, minmapmember, maxmaplen, minresmember int
	)

	cache := propdivsum2(LIMIT)

	maxmaplen = 0
	minresmember = LIMIT+1

	for i = 1; i <= LIMIT; i++ {

		chmap := make(map[int]bool)

		minmapmember = LIMIT+1

		n = i

		for len(chmap) < CHAINLIMIT {

			n = cache[n]

			if n > LIMIT {break} 

			if n < minmapmember {
				minmapmember = n
			}

			if n == i {
				if maxmaplen < len(chmap) {
					maxmaplen = len(chmap)
					minresmember = minmapmember
				}
				break
			}

			if chmap[n] { break }

			chmap[n] = true

		}
	}

	fmt.Println("maxmaplen =", maxmaplen, "minmapmember =", minresmember)
	


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
		return num
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
/*-----------------------------------------------------------------------------*/
func propdivsum2(lim int) []int {

	prop := make([]int, lim+1)
	prop[0] = 1

	for i:=1; i < lim+1; i++ {
		prop[i] = 1
	}

	for i := 2; i <= 1e3; i++ {
		prop[i*i] += i
		for j := i + 1; i * j <= LIMIT; j++ {
			prop[i*j] += i + j 
		}
	}
	return prop

}
