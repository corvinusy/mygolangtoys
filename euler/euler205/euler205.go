package main

import (
    "fmt"
	"math/rand"
	"time"
)

const LIMIT = 1e9

func main() {

	r := rand.New(rand.NewSource(time.Now().UnixNano()));

	count := 0

	for i := 0; i < LIMIT; i++ {
		if dicesPete(r) > dicesNic(r) {
			count += 1
		}
	}

	fmt.Println(float64(count)/float64(LIMIT))
}
/*----------------------------------------------------------------------------*/
func dicesPete(r *rand.Rand) int {

	result := 9

	for i := 0; i < 9; i++ {
		result += r.Intn(4)
	}

	return result
}
/*----------------------------------------------------------------------------*/
func dicesNic(r *rand.Rand) int {

	result := 6

	for i := 0; i < 6; i++ {
		result += r.Intn(6)
	}

	return result
}
