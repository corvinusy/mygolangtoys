package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", n)
	if n < 1 || n > 100 {
		panic("wrong input")
	}

	if (n%2 != 0) || (n >= 6 && n <= 20) {
		fmt.Println("Weird")
	} else {
		fmt.Println("Not Weird")
	}
}
