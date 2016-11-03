package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var n, sum, t int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&t)
		sum += t
	}
	fmt.Println(sum)
}
