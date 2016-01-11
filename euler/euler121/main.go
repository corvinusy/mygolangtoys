package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

/*
for 4 turns
variants 2 * 3 * 4 * 5 = 120
wins = (4 + 3 + 2 + 1) + 1 = 11

120/11 = 10

for 2 turns
1/2*1/3 = 1/6

for 3 turns
vars = 2 * 3 * 4 = 24
wins = (3 + 2 + 1) + 1 = 7

for 5 turns
variants = 6! = 720
wins = 5*(4 + 3 + 2 + 1) + 4*(3 + 2 + 1) + 3*(2+1) + 2*(1) + (5 + 4 + 3 + 2 + 1) + 1 = 101
wins = 1 + Sum(1..n) + (2)*Sum(1) + ... + n*(sum(1..n-1))

for 7 turns
variants = 7! = 40320
wins =
result = 1 + Sum(1..7) + Sum

1/2*1/3 + 1/2*1/4 + 1/2*1/5 + 1/3*1/4 + 1/3*1/5 + 1/4*1/5
20/120 + 15/120 + 12/120 + 10/120 + 8/120 + 6/120 = 71/120


*/

func main() {

	// boilerplate
	var games = 360360
	var rounds = 4
	var err error

	if len(os.Args) > 1 {
		rounds, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
	}
	if len(os.Args) > 2 {
		games, err = strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
	}

	//let's check it by many cycles of game
	wins := 0
	for i := 0; i < games; i++ {
		if playGame(rounds) {
			wins++
		}
	}

	fmt.Println(wins)
}

func playGame(rounds int) bool {
	blues := 0
	for i := 0; i < rounds; i++ {
		if roundUp(i + 2) {
			blues++
		}
	}
	return blues > rounds/2
}

func roundUp(quo int) bool {
	rand.Seed(time.Now().UTC().UnixNano())
	if rand.Intn(quo) == 0 {
		return true
	}
	return false
}
