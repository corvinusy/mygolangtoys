package main

import (
	"bufio"
	"fmt"
	"os"
)

type pair [2]int

func main() {
	var (
		err             error
		nPilots, nPairs int
		a, b            int
	)

	in := bufio.NewReader(os.Stdin)
	_, err = fmt.Fscan(in, &nPilots, &nPairs)
	if err != nil {
		panic(err)
	}

	pairs := make([]pair, nPairs)

	for i := 0; i < nPairs; i++ {
		_, err = fmt.Fscan(in, &a, &b)
		if err != nil {
			panic(err)
		}
		// normalize input
		if a > b {
			a, b = b, a
		}
		pairs[i] = pair{a, b}
	}

	countries := splitPilotsByCountries(nPilots, pairs)
	totalPais := 0
	accu := 0
	for k := range countries {
		totalPais += accu * len(countries[k])
		accu += len(countries[k])
	}
	fmt.Println(totalPais)
}

func splitPilotsByCountries(n int, pairs []pair) map[int]map[int]bool {
	var c0, c1, k int
	countries := map[int]map[int]bool{}
	pilotCountries := map[int]int{}
	for i := 0; i < n; i++ {
		pilotCountries[i] = i
		countries[i] = map[int]bool{}
		countries[i][i] = true
	}
	for i := range pairs {
		if pilotCountries[pairs[i][0]] != pilotCountries[pairs[i][1]] {
			c0 = pilotCountries[pairs[i][0]]
			c1 = pilotCountries[pairs[i][1]]
			// move all pilots from country 1 to country 0
			for k = range countries[c1] {
				countries[c0][k] = true
				pilotCountries[k] = c0
			}
			delete(countries, c1)
		}
	}
	return countries
}
