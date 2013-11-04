package main

import (
	"fmt"
	"math"
	"time"
)

func main() {

	var i, num, divs int64

	t1 := time.Now()
	for i = 1; i < 15000; i++ {
		num = num + i
		if num % (30) != 0 {
			continue
		}
		divs = get_div_num(num)
		if divs >= 500 {
			t2 := time.Since(t1)
			fmt.Printf("i = %d, num = %d, divs = %d\n%v\n", i, num, divs, t2)
			break
		}
	}
}

func get_div_num (num int64) int64 {

	var count int64 = 2;
	var i int64;
	sqrnum := int64(math.Sqrt(float64(num)))
	
	for i=2; i <= sqrnum; i ++ {
		if num % i == 0 {
			count +=2
		}
	}
	return count
}


/*
func main() {

	primes := make([]int64, 0, 1e8)
	var i, n, res, num, count, subcount int64;

	time1 := time.Now();

	prime_list(2e8, &primes)

	for n = 100; ; n++ {
		num = (n + 1) * n / 2
		res = num
		count = 1 //itself and 1

		for i = 0; (i <= num) && (primes[i] * primes[i] <= num); i++ {
			//find prime divisors
			subcount = 0
			for ;num % primes[i] == 0; {
				subcount++
				num /= primes[i]
			}
			count *= (subcount + 1)
		}
		if count >= 500 {
			break
		}
		if n % 10000 == 0 {
			fmt.Println("n=", n);
		}
	}

	time2 := time.Since(time1);
	fmt.Println(n, " ", res, "\n", time2)
}


func prime_list (limit int64, primes *([]int64)) {

	var sqr_lim int64 = int64(math.Sqrt(float64(limit)))

	var sieve = make([]bool, limit+1)

	var i, x, y, n int64;

	for i = 5; i <= limit ; i++ {
		sieve[i] = false;
	}

	sieve[2] = true
	sieve[3] = true
	
	for x = 1; x <= sqr_lim; x++ {
		for y = 1; y <= sqr_lim; y++ {

			n = 4 * x * x + y * y
			if (n <= limit) && ( (n % 12 == 1) || (n % 12 == 5) ) {
				sieve[n] = !sieve[n]
			}

			n = n - x * x
			if (n <= limit) && (n % 12 == 7) {
				sieve[n] = !sieve[n]
			}

			n = n - 2 * y * y
			if (x > y) && (n <= limit) && (n % 12 == 11) {
				sieve[n] = !sieve[n]
			}
    	}
	}

	for i = 5; i <= sqr_lim; i++ {
		if sieve[i] {
			n = i * i
			for j := n; j <= limit; j += n {
				sieve[j] = false
			}
		}
	}
	
	for i = 0; i <= limit; i++ {
		if sieve[i] {
			*primes = append(*primes, i)
		}
	}

	return
}
*/
