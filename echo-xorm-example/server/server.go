package server

import (
	"log"
	"net/http"

	"github.com/go-xorm/xorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
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
	s.engine, err = xorm.NewEngine("sqlite3", "/tmp/echo-xorm.sqlite.db")
	if err != nil {
		return nil, err
	}
	// connect cacher per 1000 elements -
	// CAREFUL: it leads to request time regression (from twice to tenfold)
	//cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	//s.engine.SetDefaultCacher(cacher)

	//init schema
	s.engine.ShowSQL(true)
	s.engine.Sync2(new(Reminder))

	return s, nil
}

// Run registers API and starts http-server
func (s *Server) Run() {

	// register API
	reminderHandler := ReminderHandler{Orm: s.engine}

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.Get("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to 'echo' microweb framework server with 'xorm' db orm engine\n")
	})

	// Route => handler
	e.Post("/reminders", reminderHandler.CreateReminder)
	e.Get("/reminders", reminderHandler.FindAllReminders)
	e.Get("/reminders/:id", reminderHandler.FindReminder)
	e.Put("/reminders/:id", reminderHandler.UpdateReminder)
	e.Delete("/reminders/:id", reminderHandler.DeleteReminder)

	log.Println("server started at localhost:11999")
	e.Run(standard.New(":11999"))
}
