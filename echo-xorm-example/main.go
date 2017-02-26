package main

import (
	"flag"
	"log"
	"runtime"

	"github.com/corvinusy/mygolangtoys/echo-xorm-example/server"
	_ "github.com/mattn/go-sqlite3"
)

var (
	testFlag = flag.Bool("test", true, "-test for test-mode")
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()

	server, err := server.NewServer()
	if err != nil {
		log.Fatalf("Server creation error: %s\n", err.Error())
	}
	server.Run()
}
