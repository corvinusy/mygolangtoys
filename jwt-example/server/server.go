package server

import (
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
)

// Server is methods container
type Server struct {
	engine *xorm.Engine
}

// NewServer is a constructor
func NewServer() (*Server, error) {
	var err error
	s := new(Server)
	// init engine
	s.engine, err = xorm.NewEngine("sqlite3", "/tmp/rest-xorm.sqlite.db")
	if err != nil {
		return nil, err
	}

	//init schema
	s.engine.ShowSQL(true)
	s.engine.Sync2(new(User))

	return s, nil

}

// Run registers API and starts http-server
func (s *Server) Run() {
	var (
		err    error
		router rest.App
	)

	// register API
	userHandler := UserHandler{Engine: s.engine}

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err = rest.MakeRouter(
		rest.Get("/users", userHandler.GetAllUsers),
		rest.Post("/users", userHandler.PostUser),
		rest.Get("/users/:id", userHandler.GetUser),
		rest.Put("/users/:id", userHandler.PutUser),
		rest.Delete("/users/:id", userHandler.DeleteUser),
	)
	if err != nil {
		log.Fatalf("server run error: %v\n", err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8090", api.MakeHandler()))
}
