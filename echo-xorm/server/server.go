package server

/* zeebra */

import (
	"encoding/base64"

	"github.com/go-xorm/xorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/sha3"
)

var signingKey = []byte("supersecret")

// Server is an main application object that shared (read-only) to application modules
type Server struct {
	db *xorm.Engine
}

// NewServer creates ORM-to-DB connect, init schema and fill it with predefined data
func NewServer() (*Server, error) {
	var err error
	s := new(Server)
	// init engine
	s.db, err = xorm.NewEngine("sqlite3", "/tmp/echo-xorm-test.sqlite.db")
	if err != nil {
		return nil, err
	}

	//init schema
	s.db.ShowSQL(true)
	err = s.db.Sync(new(Reminder), new(User))
	if err != nil {
		return nil, err
	}
	err = s.fillDbByPredefinedData()
	if err != nil {
		return nil, err
	}
	return s, nil
}

// Run registers API and starts http-server
func (s *Server) Run() {

	// Echo instance
	e := echo.New()

	// Global Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	var (
		authHandler     = authHandler{Orm: s.db}
		reminderHandler = reminderHandler{Orm: s.db}
		userHandler     = userHandler{Orm: s.db}
	)

	// Register routes
	// accessible routes
	e.GET("/", authHandler.getVersion)
	e.GET("/version", authHandler.getVersion)
	e.POST("/auth", authHandler.getToken)

	jwtConfig := middleware.JWTConfig{
		SigningKey: signingKey,
	}

	// restricted routes
	r := e.Group("/rest")
	r.Use(middleware.JWTWithConfig(jwtConfig)) // group middleware
	// reminders
	r.POST("/reminders", reminderHandler.CreateReminder)
	r.GET("/reminders", reminderHandler.FindAllReminders)
	r.GET("/reminders/:id", reminderHandler.FindReminder)
	r.PUT("/reminders/:id", reminderHandler.UpdateReminder)
	r.DELETE("/reminders/:id", reminderHandler.DeleteReminder)
	// users
	r.POST("/users", userHandler.CreateUser)
	r.GET("/users", userHandler.FindAllUsers)
	r.GET("/users/:id", userHandler.FindUser)
	r.PUT("/users/:id", userHandler.UpdateUser)
	r.DELETE("/users/:id", userHandler.DeleteUser)

	log.Info("server started at localhost:11110")
	log.Fatal(e.Start(":11110"))
}

func (s *Server) fillDbByPredefinedData() error {
	const adminName = "admin"
	adminsAmount, err := s.db.Count(&User{Login: adminName})
	if err != nil || adminsAmount != 0 {
		return err
	}
	if adminsAmount != 0 {
		return nil
	}
	// encrypt password
	passHash, err := bcrypt.GenerateFromPassword([]byte(adminName), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	hash := getSHA3Hash(adminName)
	if err != nil {
		return err
	}
	_, err = s.db.InsertOne(
		&User{
			Login:    adminName,
			Password: string(passHash),
			Hash:     hash,
		})
	return err
}

// support utility
func getSHA3Hash(data string) string {
	h := make([]byte, 64)
	sha3.ShakeSum256(h, []byte(data))
	return base64.StdEncoding.EncodeToString(h)
}
