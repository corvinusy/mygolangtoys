package main

import (
	"fmt"
)

func main() {

	var (
		fact [10]int64
		sum  int64 = 0
		c    int64
	)

	for i := 0; i < 10; i++ {
		fact[i] = factorial(i)
	}

	for c = 3; c <= 1e7; c++ {
		if is_curious(c, fact) {
			sum += c
			fmt.Println(c)
		}
	}

	fmt.Println(sum)

}

/*-----------------------------------------------------------------------------*/
func factorial(n int) int64 {
	if n == 0 {
		return 1
	}
	f := n
	for ; n > 1; n-- {
		f = f * (n - 1)
	}
	return int64(f)
}

/*-----------------------------------------------------------------------------*/
func is_curious(num int64, fact [10]int64) bool {

	var (
		d   int
		sum int64 = 0
	)

	for n := num; n > 0; n /= 10 {
		d = int(n % 10)
		sum = sum + fact[d]
	}
	return sum == num
}
