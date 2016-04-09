/* channels.go
 * Tests how go-routines interact with channels
 * Call: channels --help
 */

package main

import (
	"flag"
	"fmt"
	"runtime"
	"time"
)

// flag support for program
var maxProcs int
var cycleNumber int //internal burden cycle

func init() {
	flag.IntVar(&maxProcs, "maxprocs", 1, "maxprocs for testing. From 1 to 256 ")
	flag.IntVar(&cycleNumber, "cycles", 1000, "burden internal cycle for testing. From 1 to 1000000 and more ")
}

func main() {

	flag.Parse()

	procs := runtime.GOMAXPROCS(maxProcs)
	// runtime.GOMAXPROCS() returns previous procs
	// so make second call to get actual procs
	procs = runtime.GOMAXPROCS(maxProcs)
	fmt.Println("MaxProcs = ", procs)

	ch1 := make(chan int)     // this is testing channel
	ch2 := make(chan float64) // this is result collecting channel

	go filler(ch1, ch2)
	go extractor(ch1, ch2)

	fmt.Println("Total:", <-ch2, <-ch2)

}

// fill ch1 with garbage, after test fill ch2 with elapsed time
func filler(ch1 chan int, ch2 chan float64) {

	const channelSize = 1000000

	for i := 0; i < channelSize; i++ {
		for j := 0; j < cycleNumber; j++ {
			i++
		}
		// thus we avoid optimizer influence
		i = i - cycleNumber
		ch1 <- i

	}
	// signal "portion end"
	ch1 <- -1
	ch2 <- 0.0
}

func extractor(ch1 chan int, ch2 chan float64) {

	const portionSize = 100000
	total := 0.0

	for {
		t1 := time.Now().UnixNano()
		for i := 0; i < portionSize; i++ {
			// burden cycle
			for j := 0; j < cycleNumber; j++ {
				i++
			}
			i = i - cycleNumber

			m := <-ch1
			if m == -1 { // catched signal "portion end" - time to load result
				ch2 <- total
			}
		}

		// collect results
		t2 := time.Now().UnixNano()
		dt := float64(t2-t1) / 1e9 //nanoseconds ==> seconds
		total += dt
		fmt.Println(dt)
	}

}
