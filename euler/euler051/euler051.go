package main

import (
    "fmt"
	"strconv"
)

func main() {

    const LIMIT = 1e3

    var (
		i uint64
	)

	for i = 1e2+1; i < LIMIT; i += 2 {
		if i % 5 == 0 {
			continue
		}
		if is_satisfies(i) { 
			fmt.Println(i)
			break
		} 
	}

}
/*-----------------------------------------------------------------------------*/
func is_satisfies(n uint64) bool {

	var i, pcount, ncount, cn int

	ns := strconv.FormatUint(n, 10)

	// ordering XXX
	// XX***X - none
	// X***XX - none
	// ***XXX - none
	// **XX*X - none
	// *XX**X - none
	// *X*X*X - 121313
	// X*X**X - none

	for i, ncount, pcount = 1, 0, 0; (ncount <= 2) && (i <= 9); i++ {
		cns := strconv.Itoa(i) + string(ns[:1]) + strconv.Itoa(i) + string(ns[1]) + strconv.Itoa(i) + string(ns[2])
		cn, _ = strconv.Atoi(cns)
		if is_mr_prime(uint64(cn)) {
			pcount++
			fmt.Println(cn)
		} else {
			ncount++
		}
		if pcount >=8 {
			return true
		}
	}

	return false

}
/*-----------------------------------------------------------------------------*/
func is_mr_prime(n uint64) bool {
    var i, upper uint64
    
    upper = 2 * log2_n(n) * log2_n(n)
    for i = 3; (i < upper) && (i < n) ; i += 1 + upper/10 {
        if !is_witness(i, n) { return false }
    }

    return true
}
/*-----------------------------------------------------------------------------*/
func is_witness(a, n uint64) bool {
    u := n / 2
    t := 1
    for u % 2 == 0 {
        u /= 2
        t++
    }

    prev := exp_a_n_mod (a, u, n)

    var curr uint64

    for i := 1; i <= t; i++ {
        curr = (prev*prev) % n
        if (curr == 1) && (prev != 1) && (prev != n - 1) {
            return false
        }
        prev = curr
    }

    if curr != 1 {
        return false
    }
    return true
}
/*-----------------------------------------------------------------------------*/
func exp_a_n_mod (a, n, mod uint64) uint64 { // fast (a ** n ) % mod

    var result uint64 = 1

    for n != 0 {
        if n % 2 != 0 {
            result = (result * a) % mod
        }
        a = (a * a) % mod
        n /= 2
    }
    return result
}
/*-----------------------------------------------------------------------------*/
func log2_n (n uint64) uint64 { // fast log2(n)
    var result uint64 = 0

    if (n >= 1<<32) { n >>= 32; result += 32; }
    if (n >= 1<<16) { n >>= 16; result += 16; }
    if (n >= 1<< 8) { n >>=  8; result +=  8; }
    if (n >= 1<< 4) { n >>=  4; result +=  4; }
    if (n >= 1<< 2) { n >>=  2; result +=  2; }
    if (n >= 1<< 1) {           result +=  1; }
    
    return result
}
/*-----------------------------------------------------------------------------*/
