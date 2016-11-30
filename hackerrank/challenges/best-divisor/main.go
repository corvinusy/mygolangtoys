package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	divisors := getDivisors(n)
	bestDivisor := getBestDivisor(divisors)
	fmt.Println(bestDivisor)
}

func getDivisors(n int) []int {
	divs := make([]int, 2)
	divs[0] = 1
	divs[1] = n
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			divs = append(divs, i)
		}
	}
	return divs
}

func getBestDivisor(divs []int) int {
	var maxDivs []int
	var dsum, max int
	// get maxlist
	for i := range divs {
		dsum = digSum(divs[i])
		if max == dsum {
			maxDivs = append(maxDivs, divs[i])
		}
		if max < dsum {
			maxDivs = make([]int, 1)
			maxDivs[0] = divs[i]
			max = dsum
		}
	}
	// find best of Equals
	min := 1000000
	for i := range maxDivs {
		if maxDivs[i] < min {
			min = maxDivs[i]
		}
	}
	return min
}

func digSum(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10)
		n /= 10
	}
	return sum
}
