package main

import "fmt"

func main() {
	var t uint64
	fmt.Scan(&t)
	nn := make([]uint64, t)
	for i := range nn {
		fmt.Scan(&nn[i])
	}
	for i := range nn {
		fmt.Println(playGame(nn[i]))
	}
}

func playGame(n uint64) string {
	var count uint64 = 1
	for n != 1 {
		//fmt.Printf("%03d:%064b\n", count, n)
		if bitCount(n) == 1 {
			n >>= 1
		} else {
			n = n - (1 << maxBit(n))
		}
		count++
	}
	if count%2 == 0 {
		return "Louise"
	}
	return "Richard"
}

func bitCount(n uint64) uint64 {
	var bits uint64
	for n > 0 {
		if n%2 == 1 {
			bits++
		}
		n >>= 1
	}
	return bits
}

func maxBit(n uint64) uint64 {
	var mask, maxBit uint64 = 1 << 63, 63
	for n&mask == 0 {
		maxBit--
		n <<= 1
	}
	return maxBit
}
