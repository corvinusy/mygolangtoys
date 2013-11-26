package main

import (
    "fmt"
)

//99999999019

const LIMIT = 5e8

//var cache [LIMIT+1]int64

func main() {
/*
	for i:=1; i < len(cache); i++ {
		cache[i] = ave_sum_lcm(int64(i))
	}
	cache[0] = 0
*/


	for i := int64(1); i <= 200; i++ {
		fmt.Println(i,":\t", i*(i-1)/2+1, "\tave =", ave_lcm(i), "\td = ", i*(i-1)/2 +1 - ave_lcm(i) )
	}
	return 


	fmt.Println (sf(100))
	fmt.Println("ave_lcm(",LIMIT,")=", ave_lcm(LIMIT))
//	fmt.Println (sf(99999999019))
	
}
/*-----------------------------------------------------------------------------*/
func lcm(a, b int64) int64 {

	return (a / gcd(a,b)) * b

}
/*-----------------------------------------------------------------------------*/
func gcd(a, b int64) int64 {

	for b != 0 {
		a, b = b, a % b
	}

	return a
}
/*-----------------------------------------------------------------------------*/
func ave_lcm(n int64) int64 {

	var (
		sum int64 = 0
	)

	for i := int64(1); i <= n; i++ {

		sum = (sum + i / gcd(n, i)) // 999999017

		if sum > 1e15 {
			sum %= 1e10
		}
	}
	
	return sum 
}
/*-----------------------------------------------------------------------------*/
func sf(n int64) int64 {

	var (
		sum int64 = 0
		res int64
	)

	for i := int64(1); i <= n; i++ {
		res = ave_lcm(i) 
		sum = sum + res

		if i % 1e4 == 0 {
			fmt.Println(i, sum)
		}

		if sum > 1e15 {sum %= 1e10} //999999017
	}

	return sum % 999999017
}
