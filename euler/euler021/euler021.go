package main

import (
    "fmt"
)

func main() {
	var (
		sum int64 = 0
		i int64
		di int64
	)
	for i = 2; i < 10000; i++ {
		di = divsum(i)
		if (di != i) && ((divsum(di)) == i) {
			sum = sum + i
			fmt.Println(i, di)
		}
	}
	fmt.Println(sum)
}


func divsum(num int64) int64 {

	var (
		sum int64 = 1
		acc int64 = 1
		tmp, i int64
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
	return sum - num;
}
