package main

import (
	log "github.com/sirupsen/logrus"
	"os"

	"github.com/corvinusy/mygolangtoys/echo-xorm/server"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	log.SetOutput(os.Stderr)
	server, err := server.NewServer()
	if err != nil {
		log.Fatalf("Server creation error: %s\n", err.Error())
	}
	server.Run()
}
