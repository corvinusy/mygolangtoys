package main

import (
	"bytes"
	"fmt"

	"github.com/gchaincl/chanson"
)

func main() {
	ch := make(chan int)
	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		ch <- 4
		close(ch)
	}()

	buf := bytes.NewBuffer(nil)
	cs := chanson.New(buf)
	cs.Array(func(a chanson.Array) {
		for i := range ch {
			a.Push(i)
		}
	})

	fmt.Printf("%v", buf.String())
}
