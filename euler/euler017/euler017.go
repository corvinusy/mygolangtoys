package main

import (
	"fmt"
)

func main() {

	var let [1001]int
	let[0] = 0
	let[1] = len("one")
	let[2] = len("two")
	let[3] = len("three")
	let[4] = len("four")
	let[5] = len("five")
	let[6] = len("six")
	let[7] = len("seven")
	let[8] = len("eight")
	let[9] = len("nine")
	let[10] = len("ten")

	let[11] = len("eleven")
	let[12] = len("twelve")
	let[13] = len("thirteen")
	let[14] = len("fourteen")
	let[15] = len("fifteen")
	let[16] = len("sixteen")
	let[17] = len("seventeen")
	let[18] = len("eighteen")
	let[19] = len("nineteen")

	let[20] = len("twenty")
	let[30] = len("thirty")
	let[40] = len("forty")
	let[50] = len("fifty")
	let[60] = len("sixty")
	let[70] = len("seventy")
	let[80] = len("eighty")
	let[90] = len("ninety")

	lethundred := len("hundred")
	letand := len("and")

	for i := 21; i < 100; i++ {

		if i%10 != 0 {
			let[i] = let[i-i%10] + let[i%10]
		}
	}

	for i := 100; i < 1000; i++ {

		if i%100 == 0 {
			let[i] = let[i/100] + lethundred
		} else {
			let[i] = let[i/100] + lethundred + letand + let[i%100]
		}

	}

	let[1000] = len("onethousand")

	var sum int = 0
	for i := 1; i <= 1000; i++ {

		sum = sum + let[i]
	}

	fmt.Println(sum)

}
