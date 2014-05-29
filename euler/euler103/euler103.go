package main

import (
	"fmt"
)

/*
 S(B) != S(C)
 if {B} > {A} then S(B) > S(A)
*/

const LIMIT = 7

func main() {

	rslice := make([][]int, LIMIT)

	rslice[0] = make([]int, 1)
	rslice[0][0] = 1

	rslice[1] = make([]int, 2)
	rslice[1][0] = 1
	rslice[1][1] = 2

	base := 1
	newbase := 1

	for i := 2; i < LIMIT; i++ {
		rslice[i] = make([]int, i+1)

		base = newbase

		for j := 0; j < len(rslice[i]); j++ {
			rslice[i][j] = base + j
		}

		// seed the sum
		sum := 10000

		for rslice[i][i] < sum/4 {

			rslice[i] = advance(rslice[i], base)
			if isSatisfies(rslice[i]) {
				sSum := sliceSum(rslice[i])
				fmt.Println(i+1, base, rslice[i], sSum)
				if sSum < sum {
					sum = sSum
					newbase = rslice[i][i/2]
				}
			}
		}

	}

}

/*----------------------------------------------------------------------------*/
func isSatisfies(a []int) bool {

	// verify rule 1

	// enough to compare "low half of slice" and "high part minus 1 element"

	mid := len(a) / 2
	if sliceSum(a[0:mid+len(a)%2]) <= sliceSum(a[mid+1:]) {
		return false
	}

	//verify rule 2

	// build subset-sums map

	seen := make(map[int]uint)

	var (
		lenA uint = uint(len(a))
		bf   uint
	)

	for bf = 1; bf < 1<<lenA; bf++ {
		// count sum of bit elements
		sum := 0
		for q, j := bf, 0; q != 0; q, j = q>>1, j+1 {
			if q&1 != 0 {
				sum += a[j]
			}
		}
		if seen[sum] == 0 {
			seen[sum] = bf
		} else {
			if seen[sum]&bf == 0 {
				return false
			} else {
				seen[sum] = seen[sum] | bf
			}
		}
	}

	// if map was successfully builded, that sequence is ok

	return true
}

/*----------------------------------------------------------------------------*/
func advance(a []int, base int) []int {

	b := make([]int, len(a))
	copy(b, a)

	for i := 0; i < len(b)-1; i++ {
		if a[i] < a[i+1]-1 {
			b[i] += 1
			for j := 0; j < i; j++ {
				b[j] = base + j
			}
			return b
		}
	}

	// if not advanced by search than increment last and return

	b[len(b)-1] += 1

	for j := 0; j < len(b)-1; j++ {
		b[j] = base + j
	}

	return b
}

/*----------------------------------------------------------------------------*/
func sliceSum(a []int) int {

	result := 0
	for _, v := range a {
		result += v
	}
	return result
}
