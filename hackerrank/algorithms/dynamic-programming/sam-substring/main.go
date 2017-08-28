package main

import "fmt"

const mod = 1e9 + 7

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(sumOfSubs(s))
}

func sumOfSubs(s string) int {
	var sum int
	f := 1
	for i := len(s) - 1; i >= 0; i-- {
		sum = (sum + int(s[i]-'0')*f*(i+1)) % mod
		f = (f*10 + 1) % mod
	}
	return sum
}
