package main

import (
    "fmt"
)

func main() {

	const LIMIT =  0x02000000000000
	
	subj := []uint64  {1,2,5,10,20,50,100}

	//count = 100p-byte + 50p-byte + 20p-byte + 10p-byte + 5p-byte + 2p-byte + 1p-byte

	var n, count, pensum uint64

	fmt.Println(subj, LIMIT)

	for n, count = 0, 1; n <= LIMIT; {

		pensum = pens_sum(n, subj)

		switch {
		case pensum == 200: {
			count ++
			n = next_rank(n)
		}
		case pensum > 200: 
			n = next_rank(n)

			default :
			n++
		}

	}
	fmt.Println(count)
}
/*-----------------------------------------------------------------------------*/
func pens_sum(n uint64, s []uint64) uint64 {
	i := 0
	var (
		sum uint64 = 0
		b uint64
	)
	for n > 0 {
		b = n & 0xFF
		sum += s[i] * b
		i++
		n = n >> 8
	}
	return sum
}
/*-----------------------------------------------------------------------------*/
func next_rank(n uint64) uint64 {
	var i uint64 = 1
	for n & 0xFF == 0 {
		n >>= 8
		i++
	}
	n >>= 8
	n++
	n <<= (8*i)

	return n
}
