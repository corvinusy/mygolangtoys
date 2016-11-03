package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var n, v, i int
	fmt.Scan(&v)
	fmt.Scan(&n)
	x := make([]int, n)
	for i = range x {
		fmt.Scan(&x[i])
	}

	i = n / 2
	for d := i / 2; v != x[i]; d = (d + 1) / 2 {
		fmt.Printf("%d:%d ", i, d)
		if v > x[i] {
			i += d
		} else {
			i -= d
		}
	}
	fmt.Println("res=", i)
}
