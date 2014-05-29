package main

import (
	"fmt"
)

func main() {

	greatest := (101 * 55) * (101 * 55)
	least := 0

	for i := 1; i <= 100; i++ {
		least += i * i
	}

	fmt.Println(greatest - least)
}
