package main

import (
	"github.com/corvinusy/mygolangtoys/facade-example/source"
)

func main() {
	s := source.New(2)
	s.Run()
	return
}
