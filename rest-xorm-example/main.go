package main

import (
	"flag"
	"log"
	"runtime"

	"github.com/corvinusy/rest-xorm-example/restxorm"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()

	server, err := restxorm.NewServer()
	if err != nil {
		log.Fatalf("Server creation error: %s\n", err.Error())
	}
	server.Run()
}
