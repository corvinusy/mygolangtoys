package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

func main() {

	var (
		i, num, q        int64
		str              string
		count, magnitude int
	)

	const LIMIT = 1e6

	primes := make([]int64, 0, LIMIT)
	prmap := make(map[int64]bool, LIMIT)

	time1 := time.Now()

	create_primes_atkin(2*LIMIT, &primes, &prmap)
	//	create_primes_daisy(2*LIMIT, &primes, prmap)

	time2 := time.Since(time1)
	fmt.Println("primes created for ", time2)

	count = 4

	for i = 4; primes[i] < LIMIT; i++ {
		str = strconv.FormatInt(primes[i], 10)
		if strings.ContainsAny(str, "024568") {
			continue
		}

		magnitude = len(str)
		num = primes[i]
		for j := 0; j < magnitude; j++ {
			q = num % 10
			num = (num / 10) + q*int64(math.Pow10(magnitude-1))
			if prmap[num] == false {
				count--
				break
			}

		}
		count++
	}

	time3 := time.Since(time1)
	fmt.Println(count, "\n", time3)

}

/*-----------------------------------------------------------------------------*/
func create_primes_atkin(limit int64, primes *([]int64), prmap *(map[int64]bool)) int64 {

	var sqr_lim int64 = int64(math.Sqrt(float64(limit)))

	var sieve_nums = make([]bool, limit+1)

	var i, x, y, n, count int64

	for i = 5; i <= limit; i++ {
		sieve_nums[i] = false
	}

	sieve_nums[2] = true
	sieve_nums[3] = true

	for x = 1; x <= sqr_lim; x++ {
		for y = 1; y <= sqr_lim; y++ {

			n = 4*x*x + y*y
			if (n <= limit) && ((n%12 == 1) || (n%12 == 5)) {
				sieve_nums[n] = !sieve_nums[n]
			}

			n = n - x*x
			if (n <= limit) && (n%12 == 7) {
				sieve_nums[n] = !sieve_nums[n]
			}

			n = n - 2*y*y
			if (x > y) && (n <= limit) && (n%12 == 11) {
				sieve_nums[n] = !sieve_nums[n]
			}
		}
	}

	for i = 5; i <= sqr_lim; i++ {
		if sieve_nums[i] {
			n = i * i
			for j := n; j <= limit; j += n {
				sieve_nums[j] = false
			}
		}
	}

	count = 0
	for i = 0; i <= limit; i++ {
		if sieve_nums[i] {
			*primes = append(*primes, i)
			(*prmap)[i] = true
			count++
		}
	}

	return count
}

/*-----------------------------------------------------------------------------*/

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func generate(ch chan<- int64) {
	var i int64
	for i = 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func filter(in <-chan int64, out chan<- int64, prime int64) {
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}

// The prime sieve: Daisy-chain Filter processes.
func create_primes_daisy(LIMIT int64, primes *[]int64, prmap map[int64]bool) {

	var i int64
	ch1 := make(chan int64) // Create a new channel.
	go generate(ch1)        // Launch Generate goroutine.
	for i = 0; i < LIMIT; i++ {
		prime := <-ch1
		prmap[prime] = true
		*primes = append(*primes, prime)
		//		print(prime, "\n")
		ch2 := make(chan int64)
		go filter(ch1, ch2, prime)
		ch1 = ch2 //stop generate pattern
	}
}
