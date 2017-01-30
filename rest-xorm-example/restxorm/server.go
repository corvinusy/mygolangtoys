package restxorm

import (
	"log"
	"log/syslog"
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

	logWriter, err := syslog.New(syslog.LOG_INFO, "rest-xorm-example")
	if err != nil {
		log.Fatalf("Fail to create xorm system logger: %v\n", err)
	}

	logger := xorm.NewSimpleLogger(logWriter)
	logger.ShowSQL(true)

	s.engine.SetLogger(logger)

	// usual log works fine
	logger.Info("initializing rest-xorm-example")

	//init schema
	s.engine.ShowSQL(true)
	s.engine.Sync(new(Reminder))

	return s, nil

}

// Run registers API and starts http-server
func (s *Server) Run() {
	var (
		err    error
		router rest.App
	)

	// register API
	reminders := ReminderImpl{Engine: s.engine}

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err = rest.MakeRouter(
		rest.Get("/reminders", reminders.GetAllReminders),
		rest.Post("/reminders", reminders.PostReminder),
		rest.Get("/reminders/:id", reminders.GetReminder),
		rest.Put("/reminders/:id", reminders.PutReminder),
		rest.Delete("/reminders/:id", reminders.DeleteReminder),
	)
	if err != nil {
		log.Fatalf("server run error: %v\n", err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8090", api.MakeHandler()))
}
