package main

import (
	"log"
	"runtime"

	"github.com/corvinusy/mygolangtoys/echo-xorm-example/server"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	server, err := server.NewServer()
	if err != nil {
		log.Fatalf("Server creation error: %s\n", err.Error())
	}
	server.Run()
}
