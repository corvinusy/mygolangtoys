/* channel_test01.go
 * Tests how go-routines interact with channels
 * Call: channel_test01 --help
 * Pls, do not use name "channel_test", because this name always is used by go-pkg-system
 */

package main

import (
	"flag"
	"fmt"
	"runtime"
	"time"
)

// flag support for program
var MAXPROCS int
var LOAD_CYCLES int //internal burden cycle

func init() {
	flag.IntVar(&MAXPROCS, "maxprocs", 1, "maxprocs for testing. From 1 to 256 ")
	flag.IntVar(&LOAD_CYCLES, "cycles", 1000, "burden internal cycle for testing. From 1 to 1000000 and more ")
}

func main() {

	flag.Parse() //get MAXPROCS and LOAD_CYCLES from flags

	// runtime.GOMAXPROCS() returns previous max_procs
	max_procs := runtime.GOMAXPROCS(MAXPROCS)
	// second call to get real state
	max_procs = runtime.GOMAXPROCS(MAXPROCS)
	fmt.Println("MaxProcs = ", max_procs)

	ch1 := make(chan int)
	ch2 := make(chan float64)

	go chan_filler(ch1, ch2)
	go chan_extractor(ch1, ch2)

	fmt.Println("Total:", <-ch2, <-ch2)

}

func chan_filler(ch1 chan int, ch2 chan float64) {

	const CHANNEL_SIZE = 1000000

	for i := 0; i < CHANNEL_SIZE; i++ {
		for j := 0; j < LOAD_CYCLES; j++ {
			i++
		}
		//thus we avoid optimizer influence
		i = i - LOAD_CYCLES
		ch1 <- i

	}
	ch1 <- -1
	ch2 <- 0.0
}

func chan_extractor(ch1 chan int, ch2 chan float64) {

	const PORTION_SIZE = 100000
	total := 0.0

	for {
		t1 := time.Now().UnixNano()
		for i := 0; i < PORTION_SIZE; i++ {
			// burden cycle
			for j := 0; j < LOAD_CYCLES; j++ {
				i++
			}
			i = i - LOAD_CYCLES

			m := <-ch1
			if m == -1 {
				ch2 <- total
			}
		}

		t2 := time.Now().UnixNano()
		dt := float64(t2-t1) / 1e9 //nanoseconds ==> seconds
		total += dt
		fmt.Println(dt)
	}

}
