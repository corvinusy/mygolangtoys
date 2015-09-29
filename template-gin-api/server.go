package main

import (
	"flag"
	"vv-gin-api/application"
	"vv-gin-api/router"

	"github.com/gin-gonic/gin"
)

func main() {
	cfgFilename := flag.String("config",
		"/home/corvinus/ws/go/local/src/vv-gin-api/config.json", "Path to configuration file")

	var app = &application.Application{}

	app.Init(cfgFilename)
	app.ConnectToStorage()
	app.ConnectToDatabase()
	app.ConnectToCache()

	var engine *gin.Engine

	engine = gin.New()

	// Apply middleware
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	engine.Use(app.ApplyDatabase)
	engine.Use(app.ApplyStorage)
	//engine.Use(app.ApplyCache)

	// Apply routes
	var router = &router.Router{}
	router.ApplyRoutes(engine)

	// server Run
	engine.Run(":8000")

	// how to finalize?
}
