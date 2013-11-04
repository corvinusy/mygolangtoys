package main

import (
	"fmt"
)

func main() {
	var sum int = 0
	for i,j := 2,1; i < 4e6; i, j = i + j, i {
		if (i % 2 == 0) {
			sum = sum + i;
		}
	}
	fmt.Println(sum);
}
