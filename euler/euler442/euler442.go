package main

import (
    "fmt"
	"strconv"
	"strings"
	"time"
	"runtime"
)

const LIMIT = 5e5

func main() {

	const THREADS = 4

	runtime.GOMAXPROCS(THREADS)

	slice11n := make([]string,0)
	
	var i, n, prod uint64

	ch := make(chan uint64, THREADS)

	for n, prod = 1, 1; prod < 1.8e19; n++ {
		prod *=11
		slice11n = append(slice11n, strconv.FormatUint(prod, 10))
	}

	t1 := time.Now()

	for i = 1; i <= THREADS; i++{
		go gorush(i, THREADS, ch, slice11n);
	}

	prod = 0

	for i = 1; i <= THREADS; i++ {
		prod += <-ch
	}

	t2 := time.Since(t1)
	fmt.Println(LIMIT,":", prod, "time:", t2)
}
/*-----------------------------------------------------------------------------*/
func gorush(start uint64, delta uint64, result chan uint64, slice11n []string) {
	var (
		n, count uint64
		str string
	)
	count = 0
	for n = start; count <= LIMIT/delta; n = n + delta {
		str = strconv.FormatUint(n, 10)
		for _, s := range slice11n {
			if len(s) > len(str) {
				break
			}
			if strings.Contains(str, s) {
				count--
				break
			}
		}
		count++
	}

	result <- (n + 4 - start )/ delta
}
/*-----------------------------------------------------------------------------*/
/*func is_eleven_free(n uint64, slice11n []string) bool {
	str := strconv.FormatUint(n, 10)
	for _, s := range slice11n {
		if len(s) > len(str) {
			break
		}
		if strings.Contains(str, s) {
			return false
		}
	}
	return true
}
*/
