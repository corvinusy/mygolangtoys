package main

import (
	"fmt"
	"math"
)


const limit = 1e7

var primes []int

func main() {

	primes = createPrimesAtkin(limit/10)

	for n := int(1); n < limit; n++ {
		primeVec := getFactors(n)
		units := 1

		for i := range primeVec {
			units *= 2*primeVec[i] + 1
		}

		units += 1
		units /= 2
		if units > 1e3 {
			fmt.Println(n, units)
			break
		}
	}
}
/*----------------------------------------------------------------------------*/
func getFactors(num int) []int {
	sqrnum := int(math.Sqrt(float64(num)))
	vector := make([]int, 0)

	for i := 0; primes[i] <= sqrnum; i++ {
		if num%primes[i] == 0 {
			vector = append(vector, 0)
		}
		for num%primes[i] == 0 {
			vector[len(vector)-1]++
			num /= primes[i]
		}
	}
	return vector
}
/*-----------------------------------------------------------------------------*/
func createPrimesAtkin(limit int) []int {

	var (
		i, x, y, n int
	)

	sqrLim := int(math.Sqrt(float64(limit)))
	sieveNums := make([]bool, limit+1)
	primes := make([]int, 0)

	for i = 5; i <= limit; i++ {
		sieveNums[i] = false
	}

	sieveNums[2] = true
	sieveNums[3] = true

	for x = 1; x <= sqrLim; x++ {
		for y = 1; y <= sqrLim; y++ {

			n = 4*x*x + y*y
			if (n <= limit) && ((n%12 == 1) || (n%12 == 5)) {
				sieveNums[n] = !sieveNums[n]
			}

			n = n - x*x
			if (n <= limit) && (n%12 == 7) {
				sieveNums[n] = !sieveNums[n]
			}

			n = n - 2*y*y
			if (x > y) && (n <= limit) && (n%12 == 11) {
				sieveNums[n] = !sieveNums[n]
			}
		}
	}

	for i = 5; i <= sqrLim; i++ {
		if sieveNums[i] {
			n = i * i
			for j := n; j <= limit; j += n {
				sieveNums[j] = false
			}
		}
	}

	for i = 0; i <= limit; i++ {
		if sieveNums[i] {
			primes = append(primes, i)
		}
	}

	return primes
}
/*----------------------------------------------------------------------------*/
