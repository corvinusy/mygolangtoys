package main

import (
	"fmt"
	"strconv"
	"strings"
	//	"runtime"
)

const LIMIT = 1e5

func main() {

	//	runtime.GOMAXPROCS(2)

	c := make(chan string)
	res := make(chan int64)

	go generate(c)
	go process(c, res)

	fmt.Println(<-res)
}

/*-----------------------------------------------------------------------------*/
func generate(c chan<- string) {
	var (
		i, n int64
		str  string
	)

	for n = 1; n < LIMIT; n++ {
		str = strconv.FormatInt(n, 10)

		for i = 2; len(str) < 9; i++ {
			str = str + strconv.FormatInt(n*i, 10)
		}

		if len(str) == 9 {
			c <- str
		}
	}
	c <- "stop"
	return
}

/*-----------------------------------------------------------------------------*/
func process(c <-chan string, res chan<- int64) {
	var (
		sum, finsum int64
		str         string
	)
	finsum = 0
	for {
		str = <-c

		if str == "stop" {
			break
		}

		if is_pandigital(&str) {
			sum, _ = strconv.ParseInt(str, 10, 0)
			if sum > finsum {
				finsum = sum
			}
		}
	}
	res <- finsum
}

/*-----------------------------------------------------------------------------*/
func is_pandigital(s *string) bool {
	for b := '1'; b <= '9'; b++ {
		if strings.IndexRune(*s, b) == -1 {
			return false
		}
	}
	return true
}
