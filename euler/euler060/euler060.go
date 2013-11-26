package main

import (
    "fmt"
    "math"
)

func main() {

    const LIMIT = 1e7

    var n1, n2, n3, n4, n5, result uint64

    primes := make([]uint64, 0, LIMIT)
    prmap := make(map[uint64]bool, 0)

    create_primes_atkin(LIMIT * 10, &primes, prmap)

    result = 1e6

    for n1 = 1; n1 < LIMIT; n1++ {
        for n2 = n1 + 1; n2 < LIMIT; n2++ {
            if !is_satisfies(n1, n2, primes, prmap) { continue }
            for n3 = n2 + 1; n3 < LIMIT; n3++ {
                if !is_satisfies(n1, n3, primes, prmap) { continue }
                if !is_satisfies(n2, n3, primes, prmap) { continue }
                for n4 = n3 + 1; n4 < LIMIT; n4++ {
                    if !is_satisfies(n1, n4, primes, prmap) { continue }
                    if !is_satisfies(n2, n4, primes, prmap) { continue }
                    if !is_satisfies(n3, n4, primes, prmap) { continue }
                    for n5 = n4 + 1; n5 < LIMIT; n5++ {
                        if !is_satisfies(n1, n5, primes, prmap) { continue }
                        if !is_satisfies(n2, n5, primes, prmap) { continue }
                        if !is_satisfies(n3, n5, primes, prmap) { continue }
                        if !is_satisfies(n4, n5, primes, prmap) { continue }
                        sum := primes[n1] + primes[n2] + primes[n3] + primes[n4] + primes[n5]
                        fmt.Println(primes[n1], primes[n2], primes[n3], primes[n4], primes[n5], sum)
                        return
                        if result > sum {
                            result = sum
                            fmt.Println(sum)
                        }

                    }
                }
            }
        }
    }
}
/*-----------------------------------------------------------------------------*/
func is_satisfies(n1, n2 uint64, primes []uint64, prmap map[uint64]bool) bool {

    n1n2 := primes[n2] + primes[n1] * digit10n(primes[n2])
    if !prmap[n1n2] { return false }

    n2n1 := primes[n1] + primes[n2] * digit10n(primes[n1])
    if !prmap[n2n1] { return false }

    return true
    
}
/*-----------------------------------------------------------------------------*/
func create_primes_atkin (limit uint64, primes *([]uint64), prmap map[uint64]bool)  {

    var sqr_lim uint64 = uint64(math.Sqrt(float64(limit)))

    var sieve_nums = make([]bool, limit+1)

    var i, x, y, n uint64;

    for i = 5; i <= limit ; i++ {
        sieve_nums[i] = false;
    }

    sieve_nums[2] = true
    sieve_nums[3] = true
    
    for x = 1; x <= sqr_lim; x++ {
        for y = 1; y <= sqr_lim; y++ {

            n = 4 * x * x + y * y
            if (n <= limit) && ( (n % 12 == 1) || (n % 12 == 5) ) {
                sieve_nums[n] = !sieve_nums[n]
            }

            n = n - x * x
            if (n <= limit) && (n % 12 == 7) {
                sieve_nums[n] = !sieve_nums[n]
            }

            n = n - 2 * y * y
            if (x > y) && (n <= limit) && (n % 12 == 11) {
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
    
    for i = 0; i <= limit; i++ {
        if sieve_nums[i] {
            *primes = append(*primes, i)
            prmap[i] = true
        } 
    }

    return
}
/*-----------------------------------------------------------------------------*/
func digit10n (n uint64) uint64 {
    var count uint64 = 1
    for n > 0 {
        count *= 10
        n /= 10
    }
    return count
}
