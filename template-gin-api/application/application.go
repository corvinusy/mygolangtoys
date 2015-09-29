package application

import (
	"database/sql"
	"encoding/json"
	"log"

	"gopkg.in/redis.v2"

	"io/ioutil"

	_ "github.com/lib/pq"
)

type Application struct {
	Configuration *Configuration
	Database      *sql.DB
	Storage       *redis.Client
	Cache         *redis.Client
	DbVersion     string
	ApiVersion    string
}

func (app *Application) Init(filename *string) {

	data, err := ioutil.ReadFile(*filename)

	if err != nil {
		log.Printf("Can't read configuration file: %s\n", err)
		panic(err)
	}

	app.Configuration = &Configuration{}

	err = json.Unmarshal(data, &app.Configuration)

	if err != nil {
		log.Printf("Can't parse configuration file: %s\n", err)
		panic(err)
	}

	app.ApiVersion = app.Configuration.ApiVersion

	//	app.Store = sessions.NewCookieStore([]byte(app.Configuration.Secret))
}

func (app *Application) ConnectToDatabase() {
	var (
		err error
		dsn string
		ver string
	)
	//postgresql://[user[:password]@][netloc][:port][/dbname][?param1=value1&...]
	dsn = "postgres://" + app.Configuration.Database.User +
		":" + app.Configuration.Database.Password +
		"@" + app.Configuration.Database.Host +
		":" + app.Configuration.Database.Port +
		"/" + app.Configuration.Database.Dbname +
		"?sslmode=" + app.Configuration.Database.Sslmode

	app.Database, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Printf("Can't connect to the database: %v\n", err)
		panic("app.ConnectToDatabase()\n" + err.Error())
	}

	qtst := "SELECT constant_value FROM msc.constants WHERE constant_name = 'SCHEMA_VERSION';"
	err = app.Database.QueryRow(qtst).Scan(&ver)
	if err != nil {
		log.Printf("Can't query schema version: %v\n", err)
		panic("app.ConnectToDatabase()\n" + err.Error())
	}
	app.DbVersion = ver
}

func (app *Application) ConnectToStorage() {
	var err error
	app.Storage = redis.NewTCPClient(&redis.Options{
		Addr:     app.Configuration.Storage.Addr,
		Password: app.Configuration.Storage.Password,
		DB:       int64(app.Configuration.Storage.Db)})
	_, err = app.Storage.Ping().Result()
	if err != nil {
		log.Printf("Can't connect to the storage: %v\n", err)
		panic("app.ConnectToStorage()\n" + err.Error())
	}
}

func (app *Application) ConnectToCache() {
	var err error
	app.Cache = redis.NewTCPClient(&redis.Options{
		Addr:     app.Configuration.Cache.Addr,
		Password: app.Configuration.Cache.Password,
		DB:       int64(app.Configuration.Cache.Db)})
	_, err = app.Storage.Ping().Result()
	if err != nil {
		log.Printf("Can't connect to the cache: %v\n", err)
		panic("app.ConnectToCache()\n" + err.Error())
	}
}

func (app *Application) Close() {
	log.Println("Bye!")
	app.Storage.Close()
	app.Database.Close()
	app.Cache.Close()
}
