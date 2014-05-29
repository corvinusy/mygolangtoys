package main

import (
	"fmt"
)

func main() {

	const LIMIT = 2e7

	subj := []int{1, 0, 2, 3, 4, 5, 6, 7, 8, 9}

	var count int = 0

	fmt.Println("start:", subj)

	for i := 0; getnum(subj) != 9876543210; i++ {
		next_permutation(0, len(subj)-1, subj)
		if is_satisfies(subj) {
			count += getnum(subj)
		}
	}
	fmt.Println(count)
}

/*-----------------------------------------------------------------------------*/
func getnum(s []int) int {
	num := s[9] + 10*s[8] + 100*s[7] + 1000*s[6] + 1e4*s[5] + 1e5*s[4] +
		1e6*s[3] + 1e7*s[2] + 1e8*s[1] + 1e9*s[0]
	return num
}

/*-----------------------------------------------------------------------------*/
func is_satisfies(s []int) bool {

	switch {
	//2
	case s[3]%2 != 0:
		return false
		// 5
	case s[5]%5 != 0:
		return false
		// 3
	case (s[2]+s[3]+s[4])%3 != 0:
		return false
		// 7
	case (s[6]+10*s[5]+100*s[4])%7 != 0:
		return false
		// 11
	case (s[7]+10*s[6]+100*s[5])%11 != 0:
		return false
		// 13
	case (s[8]+10*s[7]+100*s[6])%13 != 0:
		return false
		// 17
	case (s[9]+10*s[8]+100*s[7])%17 != 0:
		return false
	}
	return true
}

/*-----------------------------------------------------------------------------*/
func next_permutation(start int, end int, s []int) {

	if start >= end {
		return
	}

	var (
		i, j, k int
		tmp     int
	)

	i = end

	for {
		j = i
		i--

		if s[i] < s[j] {
			for k = end; s[i] > s[k]; k-- {

			}
			//swap(s[i], s[k])
			tmp = s[i]
			s[i] = s[k]
			s[k] = tmp

			reverse(j, end, s)
			return
		}

		if i == start {
			reverse(start, end, s)
			return
		}
	}
}

/*-----------------------------------------------------------------------------*/
func reverse(start int, end int, s []int) {

	if start >= end {
		return
	}

	var tmp int

	for {
		if start >= end {
			return
		} else {
			tmp = s[start]
			s[start] = s[end]
			s[end] = tmp
			end--
			start++
		}
	}
}
