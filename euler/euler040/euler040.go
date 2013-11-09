package main

import (
    "fmt"
	"strconv"
)

func main() {

	const LIMIT = 1e6 + 1

	var i int64

	var s string = "."

	for i = 1; len(s) <= LIMIT; i++ {
		s = s + strconv.FormatInt(i, 10)
	}

	fmt.Println((s[1]-'0') * (s[10]-'0') * (s[100]-'0') * (s[1e3]-'0') * (s[1e4]-'0') * (s[1e5]-'0') * (s[1e6]-'0'))

}


