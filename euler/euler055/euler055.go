package main

import (
    "fmt"
	"math/big"
)

func main() {
	
	const (
		LIMIT = 10000
		ITERS = 50
	)

	var (
		i int64
	)

	z1 := new(big.Int)
	z2 := new(big.Int)

	count := 0
	count_rev := 0

	for i = 1; i < LIMIT; i++ {
		z1.SetInt64(i)
		count++
		for j := 0; j <= ITERS; j++ {

			z2.SetString(reverse(z1.String()), 10)
			count_rev++
			z1.Add(z1, z2)

			if is_palindrome(z1.String()) {
				count--
				break
			}
		}
	}

	fmt.Println(count, count_rev)
}
/*-----------------------------------------------------------------------------*/
func is_palindrome(s string) bool {

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}

	return true
}
/*-----------------------------------------------------------------------------*/
func reverse(s string) string {

    runes := []rune(s)

    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }

    return string(runes)
}
