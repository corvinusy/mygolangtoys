package main

import "fmt"

func main() {
	var s, t int
	fmt.Scan(&s, &t)
	var a, b int
	fmt.Scan(&a, &b)
	var m, n int
	fmt.Scan(&m, &n)
	var as, acount int
	for i := 0; i < m; i++ {
		fmt.Scan(&as)
		if a+as >= s && a+as <= t {
			acount++
		}
	}
	var bs, bcount int
	for i := 0; i < n; i++ {
		fmt.Scan(&bs)
		if b+bs >= s && b+bs <= t {
			bcount++
		}
	}
	fmt.Println(acount)
	fmt.Println(bcount)
}
