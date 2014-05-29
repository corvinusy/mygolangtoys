package main

import (
	"fmt"
	"math/big"
)

/*-----------------------------------------------------------------------------*/
/*
D = diophantine multiplier, or a radical for square root
a[0] = uint(sqrt(D)) // 2
P[0] = 0
Q[0] = 1
q[0] = 1
p[0] = a[0]

P[1] = a[0]
Q[1] = D - a[0] * a[0]

P[n] = a[n-1]*Q[n-1] - P[n-1]
Q[n] = (D - P[n]**2) / Q[n-1]

a[n] = uint((a[0]+P[n])/Q[n])
p[1] = a[0]*a[1] + 1
q[1] = a[1]

p[n] = a[n] * p[n-1] + p[n-2]
q[n] = a[n] * q[n-1] + q[n-2]

*/
/*-----------------------------------------------------------------------------*/
const LIMIT = 1e3

func main() {

	var (
		n     int64
		count int64 = 0
	)

	a := make([]int64, 0)
	p := make([]*big.Int, 0)
	q := make([]*big.Int, 0)

	a = append(a, 2, 1)

	q = append(q, big.NewInt(1), big.NewInt(a[1]))
	p = append(p, big.NewInt(a[0]), big.NewInt(a[0]*a[1]+1))

	for n = 2; n <= 100; n++ {

		if (n+1)%3 == 0 {
			a = append(a, 2*(n+1)/3)
		} else {
			a = append(a, 1)
		}

		zp := big.NewInt(a[n])
		zp.Mul(zp, p[n-1])
		zp.Add(zp, p[n-2])
		p = append(p, zp)

		zq := big.NewInt(a[n])
		zq.Mul(zq, q[n-1])
		zq.Add(zq, q[n-2])
		q = append(q, zq)

		fmt.Println(n, a[n-1], p[n-1], "/", q[n-1])
	}

	for _, s := range p[99].String() {
		count += int64(s - '0')
	}
	fmt.Println(count)

}

/*-----------------------------------------------------------------------------*/
func is_periodic(a, pb, qb []int64) (bool, int64) {
	for i := 2; i < len(a); i++ {
		if (a[1] == a[i]) && (qb[1] == qb[i]) && (pb[1] == pb[i]) {
			return true, int64(i - 1)
		}
	}
	return false, 0
}
