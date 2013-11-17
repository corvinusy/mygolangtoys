package main

import (
    "fmt"
	"math"
)

/*-----------------------------------------------------------------------------*/
/*
D = 5
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
const LIMIT = 1e4

func main() {

	var (
		d uint
		b bool
		period int
	)

	ds := make([]int, 2)

	for d = 2; d <= LIMIT; d++ {

		sqd := uint(math.Sqrt(float64(d)))

		if sqd*sqd == d {
			ds = append(ds, 0)
			continue
		}

		a := make([]uint, 0)
		pb := make([]uint, 0)
		qb := make([]uint, 0)
		//	p := make([]uint, 0)
		//	q := make([]uint, 0)

		a = append(a, sqd)
		pb = append(pb, 0, a[0])
		qb = append(qb, 0, d - a[0]*a[0])
		a = append(a, (a[0] + pb[1])/qb[1])

		//	q = append(q, 1, a[1])
		//	p = append(p, a[0], a[0]*a[1] + 1)

		for n := 2; ; n++ {
			pb = append(pb,  a[n-1]*qb[n-1] - pb[n-1] )
			qb = append(qb, (d - pb[n]*pb[n]) / qb[n-1] )
			
			a = append(a, (a[0] + pb[n])/qb[n] )
			//		p = append(p, a[n] * p[n-1] + p[n-2] )
			//		q = append(q, a[n] * q[n-1] + q[n-2] )
			b, period = is_periodic(a, pb, qb)
			if b {
				ds = append(ds, period)
				break
			}
		}

	}

	odds := 0

	for _, d := range ds {
		if d % 2 != 0 {
			odds++
		}
	}


	fmt.Println(odds)
	
}
/*-----------------------------------------------------------------------------*/
func is_periodic(a, pb, qb []uint) (bool, int) {
	for i:=2; i < len(a); i++ {
		if (a[1] == a[i]) && (qb[1] == qb[i]) && (pb[1] == pb[i]) {
			return true, i-1
		}
	}
	return false, 0
}
