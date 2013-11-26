package main

import (
    "fmt"
)

func main() {

	for k := 7; k <= 7; k++ {
		vec := make([]int, k)
		for i, _ := range vec {
			vec[i] = 1
		}
		
		for /*vsum(vec) != vproduct(vec)*/ {
			inc(vec)
			fmt.Println(vec, vsum(vec), vproduct(vec))
		}
		
		fmt.Println("result", k, vec, vsum(vec))
	}
}
/*-----------------------------------------------------------------------------*/
func vsum(vec []int) int {
	sum := 0
	for i, _ := range vec {
		sum += vec[i]
	}
	return sum
}
/*-----------------------------------------------------------------------------*/
func vproduct(vec []int) int {
	prod := 1
	for i, _ := range vec {
		prod *= vec[i]
	}
	return prod
}
/*-----------------------------------------------------------------------------*/
func inc(vec []int) {

	lim := len(vec)

	i := lim-1

	for ;i > 0; i-- {
		if vec[i] < vec[i-1] {
			vec[i]++
			if i == lim - 1 {
				break
			}
			for j := i+1; j < lim; j++ {
				vec[j] = 1
			}

			return
		}
	}

	vec[0]++

	if vec[0] > lim {
		panic("fuck")
	}

	for j := 1; j < lim; j++ {
		vec[j] = 1
	}

	return
}
/*-----------------------------------------------------------------------------*/
func divsum(num int) int {

	var (
		sum int = 1
		acc int = 1
		tmp, i int
	)
	
	//trivial case
	if num == 1 {
		return num
	}
	// powers of 2
	for num % 2 == 0 {
		acc *= 2
		num /= 2
	}
	acc = acc * 2 - 1
	
	if num == 1 {
		return acc
	} else {
		sum = acc
	}
	
	// brute force others
	for i = 3; i * i <= num; i += 2 {
		acc = 1
		tmp = 1
		for num % i == 0 {
			acc = acc * i
			tmp = tmp + acc
			num = num / i
		}
		if (acc > 1) {
			sum = sum * tmp
		}
	}

	if (num > 1) {
		sum = sum * (num + 1)
	}
	return sum;
}
