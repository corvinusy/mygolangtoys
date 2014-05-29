package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {

	var (
		a, b, c, start, end byte
		num                 uint64
		res, r              string
	)

	raw_bytes, err := ioutil.ReadFile("cipher1.txt")
	if err != nil {
		panic("File not read")
	}

	source := strings.TrimSpace(string(raw_bytes)) // kill last "\n"

	strs := strings.Split(source, ",")

	nums := make([]byte, 0, len(strs))

	for _, s := range strs {
		num, err = strconv.ParseUint(string(s), 10, 8)
		if err != nil {
			panic("Byte not recognised")
		}
		nums = append(nums, byte(num))
	}

	start = 97 // 'a'
	end = 116  // 'z'

	for a = start; a <= end; a++ {
		for b = start; b <= end; b++ {
			for c = start; c <= end; c++ {
				key := [...]byte{a, b, c}
				r = cipher_decode(nums, key)
				if check_res(r) {
					fmt.Println(a, b, c, r)
					res = r
				}
			}
		}
	}

	sum := 0
	for _, a := range res {
		sum += int(a)
	}
	fmt.Println(sum)
}

/*-----------------------------------------------------------------------------*/
func cipher_decode(s []byte, key [3]byte) string {
	r := make([]byte, len(s))
	for i, a := range s {
		r[i] = a ^ key[i%3]
	}
	return string(r)
}

/*-----------------------------------------------------------------------------*/
func check_res(s string) bool {
	if strings.Contains(s, "the") && strings.Contains(s, "this") &&
		strings.Contains(s, "and") && strings.Contains(s, "not") {
		return true
	}
	return false
}
