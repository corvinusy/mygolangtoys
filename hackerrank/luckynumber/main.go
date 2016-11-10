package main

import "fmt"

func main() {
	var (
		trials int
	)
	fmt.Scan(&trials)
	// find prime map
	primes := getPrimes(9 * 9 * 18)

	results := make([]int, trials)
	for t := 0; t < trials; t++ {
		d := getInput()
		results[t] = getLucky(d[0], d[1], primes)
	}
	for t := 0; t < trials; t++ {
		fmt.Println(results[t])
	}
}

func getInput() [2]int {
	var d [2]int
	fmt.Scanln(&d[0], &d[1])
	return d
}

func getLucky(start, end int, primes map[int]bool) int {
	//fmt.Println(primes)
	var count int
	for n := start; n <= end; n++ {
		if isLucky(n, primes) {
			fmt.Println("lucky =", n)
			count++
		}
	}
	return count
}

func getPrimes(max int) map[int]bool {
	primes := make(map[int]bool, 2)
	primes[2] = true
	sieve := []int{2}

Loop:
	for i := 3; i <= max; i += 2 {
		for _, v := range sieve {
			if i%v == 0 {
				continue Loop
			}
		}
		sieve = append(sieve, i)
		primes[i] = true
	}

	return primes
}

func isLucky(n int, primes map[int]bool) bool {
	return isPrime(getDigitSum(n), primes) && isPrime(getDigitSqSum(n), primes)
}

func getDigitSum(n int) int {
	var sum int
	for n > 0 {
		sum += n % 10
		n = n / 10
	}
	return sum
}

func getDigitSqSum(n int) int {
	var sum, x int
	for n > 0 {
		x = n % 10
		sum += x * x
		n = n / 10
	}
	return sum
}

func isPrime(n int, primes map[int]bool) bool {
	_, ok := primes[n]
	return ok
}
