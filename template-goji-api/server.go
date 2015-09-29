package main

import (
	"flag"

	"github.com/corvinusy/vv-goji-api/core"
	"github.com/corvinusy/vv-goji-api/router"

	"github.com/golang/glog"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/graceful"
)

func main() {
	cfgFileName := flag.String("config",
		"/home/corvinus/ws/go/src/github.com/corvinusy/vv-goji-api/config.json", "Path to configuration file")

	flag.Parse()
	defer glog.Flush()

	var application = &core.Application{}

	application.Init(cfgFileName)
	application.ConnectToStorage()
	application.ConnectToDatabase()
	application.ConnectToCache()

	// Apply middleware to framework
	goji.Use(application.ApplyDatabase)
	goji.Use(application.ApplyStorage)
	goji.Use(application.ApplyCache)

	// Apply routes
	var router = &router.Router{}
	router.ApplyRoutes()

	// finalization
	graceful.PostHook(func() {
		application.Close()
	})
	goji.Serve()
}
