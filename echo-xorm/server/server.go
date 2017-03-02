package server

/* zeebra */

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/binary"

	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/go-xorm/xorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
	"golang.org/x/crypto/bcrypt"
	//"golang.org/x/crypto/sha3"
)

var signingKey = []byte("supersecret")

// Server is an main application object that shared (read-only) to application modules
type Server struct {
	engine *xorm.Engine
}

// NewServer creates ORM-to-DB connect, init schema and fill it with predefined data
func NewServer() (*Server, error) {
	var err error
	s := new(Server)
	// init engine
	s.engine, err = xorm.NewEngine("sqlite3", "/tmp/echo-xorm-test.sqlite.db")
	if err != nil {
		return nil, err
	}

	//init schema
	s.engine.ShowSQL(true)
	err = s.engine.Sync(new(Reminder), new(User))
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
		authHandler     = authHandler{Orm: s.engine}
		reminderHandler = reminderHandler{Orm: s.engine}
		userHandler     = userHandler{Orm: s.engine}
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

	log.Println("server started at localhost:11110")
	e.Logger.Fatal(e.Start(":11110"))
}

func (s *Server) fillDbByPredefinedData() error {
	const adminName = "admin"
	adminsAmount, err := s.engine.Count(&User{Login: adminName})
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

	hash, err := getMd5Hash([]byte(adminName))
	if err != nil {
		return err
	}
	_, err = s.engine.InsertOne(
		&User{
			Login:    adminName,
			Password: string(passHash),
			Hash:     hash,
		})
	return err
}

// support utility
func getMd5Hash(data []byte) (string, error) {
	var err error
	seed := time.Now().Unix()
	seedBytes := make([]byte, binary.Size(seed))
	binary.PutVarint(seedBytes, seed)

	hasher := md5.New()
	_, err = hasher.Write(seedBytes)
	if err != nil {
		return "", err
	}
	_, err = hasher.Write(data)
	if err != nil {
		return "", err
	}

	hash := hasher.Sum(nil)
	return base64.StdEncoding.EncodeToString(hash), nil
}
