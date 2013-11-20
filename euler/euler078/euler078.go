package main

import (
    "fmt"
    "math/big"
)

// Partition Fuction P
// P(1, 1) = 1
// P(n, k) = P(n, k-1) + P(n - k, k)
// P(n, n) = 1
// P(n, 0) = 0
// P(n, 1) = n
// P(n, k > n) = 0
// P(n) = Sum((-1)**(k+1) * P(n - k*(3*k - 1)/2) + P(n - k*(3*k + 1)/2) ) for k = 1..n
// p(k)=p(k−1)+p(k−2)−p(k−5)−p(k−7)+p(k−12)+p(k−15)−p(k−22)−....

const LIMIT = 100001

func main() {

    cache := make([]*big.Int, LIMIT+1)

    cache[0] = big.NewInt(0)
    cache[1] = big.NewInt(1)

    for i := 100; i <= LIMIT; i++ {

        s := pn(i, cache).String()

        if s[len(s)-6:] == "000000" {
            fmt.Println(i-1, s)
            break
        }

    }

}

/*-----------------------------------------------------------------------------*/
func pent(k int) int {
    return k * (3*k - 1) / 2 // 1, 5, 12
}

/*-----------------------------------------------------------------------------*/
func pn(n int, cache []*big.Int) *big.Int {

    if n <= 0 {
        return big.NewInt(0)
    }

    if cache[n] != nil {
        return cache[n]
    }

    z := big.NewInt(0)

    for k := 1; pent(k) < n; k += 2 {

        z.Add(z, pn(n-pent(k), cache))
        z.Add(z, pn(n-pent(-k), cache))
        z.Sub(z, pn(n-pent(k+1), cache))
        z.Sub(z, pn(n-pent(-k-1), cache))

    }

    cache[n] = z
    return z
}
