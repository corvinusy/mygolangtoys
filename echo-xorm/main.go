package main

import (
	log "github.com/Sirupsen/logrus"

	"github.com/corvinusy/mygolangtoys/echo-xorm/server"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	server, err := server.NewServer()
	if err != nil {
		log.Fatalf("Server creation error: %s\n", err.Error())
	}
	server.Run()
}
