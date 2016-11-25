package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var n, down, up int
	fmt.Scan(&n)
	x := make([][]int, n)
	for i := range x {
		x[i] = make([]int, n)
		for j := range x[i] {
			fmt.Scan(&x[i][j])
			if i == j {
				down += x[i][j]
			}
			if i == n-1-j {
				up += x[i][j]
			}
		}
	}
	fmt.Print(x)

	if (up - down) < 0 {
		fmt.Println(down - up)
	} else {
		fmt.Println(up - down)
	}
}
