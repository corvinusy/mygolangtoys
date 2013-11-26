package main

import (
    "fmt"
	"math"
	"math/big"
)

/*-----------------------------------------------------------------------------*/
/*
D = diophantine multiplier
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
		d, n int64
		b bool
		period int64
	)

	ds := make([]int64, 2)
	zres := big.NewInt(0)

	for d = 2; d <= LIMIT; d++ {

		sqd := int64(math.Sqrt(float64(d)))

		if sqd*sqd == d {
			ds = append(ds, 0)
			continue
		}

		period = 0
		b = false

		a := make([]int64, 0)
		pb := make([]int64, 0)
		qb := make([]int64, 0)
		p := make([]*big.Int, 0)
		q := make([]*big.Int, 0)

		a = append(a, sqd)
		pb = append(pb, 0, a[0])
		qb = append(qb, 0, d - a[0]*a[0])
		a = append(a, (a[0] + pb[1])/qb[1])

		q = append(q, big.NewInt(1), big.NewInt(a[1]))
		p = append(p, big.NewInt(a[0]), big.NewInt(a[0]*a[1] + 1))

		for n = 2; ; n++ {
			pb = append(pb,  a[n-1]*qb[n-1] - pb[n-1] )
			qb = append(qb, (d - pb[n]*pb[n]) / qb[n-1] )

			a = append(a, (a[0] + pb[n])/qb[n] )

//			p = append(p, a[n] * p[n-1] + p[n-2] )
//			q = append(q, a[n] * q[n-1] + q[n-2] )

			zp := big.NewInt(a[n])
			zp.Mul(zp, p[n-1])
			zp.Add(zp, p[n-2])
			p = append(p, zp )

			zq := big.NewInt(a[n])
			zq.Mul(zp, p[n-1])
			zq.Add(zp, p[n-2])
			q = append(q, zq )

			if period == 0 {
				b, period = is_periodic(a, pb, qb)
			}
			if b && period % 2 == 0 {
				if zres.Cmp(p[n-2]) == -1 {
					fmt.Println(d, p[n-2], q[n-2])
					zres = p[n-2]
				}
				break
			} 
			if b && (period % 2 != 0) && n == period*2  {
				if zres.Cmp(p[n-1]) == -1 {
					fmt.Println(d, p[n-1], q[n-1])
					zres = p[n-1]
				}
				break
			}

		}
		
	}
	fmt.Println("result:", zres)
}
/*-----------------------------------------------------------------------------*/
func is_periodic(a, pb, qb []int64) (bool, int64) {
	for i:=2; i < len(a); i++ {
		if (a[1] == a[i]) && (qb[1] == qb[i]) && (pb[1] == pb[i]) {
			return true, int64(i-1)
		}
	}
	return false, 0
}
