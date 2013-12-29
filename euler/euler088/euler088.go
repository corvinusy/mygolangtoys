package main

import (
    "fmt"
)

const LIMIT = 12000
const UPPER = 13000

func main() {

	tmpmap := make(map[int]int) // temp map for reversed data

	cache := make([][][]int,UPPER) // cache with always found factor lists

	for n := 1; n < UPPER; n++ {

		factorLists := getFactorLists(n, cache) // divSet is [][]int

		for i := range factorLists {

			product := 1
			for j := range factorLists[i] {
				product *= factorLists[i][j]
			}

			sum := 0
			for j := range factorLists[i] {
				sum += factorLists[i][j]
			}

			k := product - sum + len(factorLists[i])

			if tmpmap[k] == 0 {
				tmpmap[k] = n
			} 

		}
	}

	delete (tmpmap, 1) // remove 1 according to task definition

	//prepare map for summation of keys
	summap := make(map[int]bool)

	for k, v := range tmpmap {
		if k <= LIMIT {
			summap[v] = true
		}
	}

	sum := 0

	for k := range summap {
		sum += k
	}

	fmt.Println("sum =", sum)

}
/*-----------------------------------------------------------------------------*/
func getFactorLists(n int, cache [][][]int) [][]int {
	// returns list of factor combinations

	if n < 2 { return nil }

	if cache[n] != nil {
		return cache[n]
	}
	
	// result placeholder
	factorLists := make([][]int, 0)

	// find factors
	factors := make([]int, 0)

	for i := 2; i <= n/2; i++ {
		if n % i == 0 {
			factors = append(factors, i)
		}
	}

	factors = append(factors, n) // add itself to factors

	// find tail-subfactors for all found factors
	for i := range factors {

		tail := getFactorLists(n / factors[i], cache)

		if tail == nil {
			fl := make([]int, 1)
			fl[0] = factors[i]
			factorLists = append(factorLists, fl)
			continue
		}

		for j := range tail {
			fl := make([]int, 1)
			fl[0] = factors[i]
			fl = append(fl, tail[j]...)
			factorLists = append(factorLists, fl)
		}
	}

	cache[n] = factorLists

	return factorLists
}
