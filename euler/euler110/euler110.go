package main

import (
	"fmt"
	"github.com/cznic/mathutil"
)

const (
	limit = 1e18

	base = uint64(2 * 3 * 5 * 7 * 11 * 13 * 17 * 19 * 23 * 29 * 31)
)

type factor struct {
	prime uint64
	power uint64
}

type divisor struct {
	number uint64
	factors []factor
}

func main() {

	d := new(divisor)

	d.number = base
	d.updateFactors()

	for ;d.number < limit; d.getNext()  {

		dr := d.getReciprocals()

		if dr > 4e6 {

			fmt.Printf("n = %4d : a(n) = %d\n", d.number, dr)
			break
		}
	}
	
	return
}
/*----------------------------------------------------------------------------*/
func (d *divisor) getNumberOfDivisors() uint64 {

	//calculate sigma function
	result := uint64(1)

	for i := range d.factors {
		result *= 2 * d.factors[i].power + 1
	}

	return result
}
/*----------------------------------------------------------------------------*/
func (d *divisor) getReciprocals() uint64 {

	return (d.getNumberOfDivisors() + 1) / 2

}
/*----------------------------------------------------------------------------*/
func (d *divisor) getNext() {

	d.number += base
	d.updateFactors()

}
/*----------------------------------------------------------------------------*/
func (d *divisor) updateFactors() {

	if mathutil.IsPrimeUint64(d.number) {
		d.factors = []factor{{d.number,1}}
		return
	}

	fs := make([]factor,0)
	n := d.number

	prime32 := uint32(0)

	for {
		prime32, _ = mathutil.NextPrime(prime32);
		if prime32 > 5e1 {
			break
		}
		
		prime := uint64(prime32)

		if prime*prime > n {
			break
		}
		
		power := uint64(0)
		
		for n % prime == 0 {
			n /= prime
			power++
		}

		if power != 0 {
			fs = append(fs, factor{prime, power} )
		}

		if n == 1 {
			break
		}
	}

	if n != 1 {
		fs = append(fs, factor{n, 1} )		
	}

	d.factors = fs

}
/*----------------------------------------------------------------------------*/
