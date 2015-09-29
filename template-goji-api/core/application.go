package core

import (
	"database/sql"
	"encoding/json"

	"github.com/golang/glog"
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

func (application *Application) Init(filename *string) {

	data, err := ioutil.ReadFile(*filename)

	if err != nil {
		glog.Fatalf("Can't read configuration file: %s", err)
		panic(err)
	}

	application.Configuration = &Configuration{}

	err = json.Unmarshal(data, &application.Configuration)

	if err != nil {
		glog.Fatalf("Can't parse configuration file: %s", err)
		panic(err)
	}

	application.ApiVersion = application.Configuration.ApiVersion

	//	application.Store = sessions.NewCookieStore([]byte(application.Configuration.Secret))
}

func (application *Application) ConnectToDatabase() {
	var (
		err error
		dsn string
		ver string
	)
	//postgresql://[user[:password]@][netloc][:port][/dbname][?param1=value1&...]
	dsn = "postgres://" + application.Configuration.Database.User +
		":" + application.Configuration.Database.Password +
		"@" + application.Configuration.Database.Host +
		":" + application.Configuration.Database.Port +
		"/" + application.Configuration.Database.Dbname +
		"?sslmode=" + application.Configuration.Database.Sslmode

	application.Database, err = sql.Open("postgres", dsn)
	if err != nil {
		glog.Fatalf("Can't connect to the database: %v", err)
		panic("application.ConnectToDatabase()\n" + err.Error())
	}

	qtst := "SELECT constant_value FROM msc.constants WHERE constant_name = 'SCHEMA_VERSION';"
	err = application.Database.QueryRow(qtst).Scan(&ver)
	if err != nil {
		glog.Fatalf("Can't query schema version: %v", err)
		panic("application.ConnectToDatabase()\n" + err.Error())
	}
	application.DbVersion = ver
}

func (application *Application) ConnectToStorage() {
	var err error
	application.Storage = redis.NewTCPClient(&redis.Options{
		Addr:     application.Configuration.Storage.Addr,
		Password: application.Configuration.Storage.Password,
		DB:       int64(application.Configuration.Storage.Db)})
	_, err = application.Storage.Ping().Result()
	if err != nil {
		glog.Fatalf("Can't connect to the storage: %v", err)
		panic("application.ConnectToStorage()\n" + err.Error())
	}
}

func (application *Application) ConnectToCache() {
	var err error
	application.Cache = redis.NewTCPClient(&redis.Options{
		Addr:     application.Configuration.Cache.Addr,
		Password: application.Configuration.Cache.Password,
		DB:       int64(application.Configuration.Cache.Db)})
	_, err = application.Storage.Ping().Result()
	if err != nil {
		glog.Fatalf("Can't connect to the cache: %v", err)
		panic("application.ConnectToCache()\n" + err.Error())
	}
}

func (application *Application) Close() {
	glog.Info("Bye!")
	application.Storage.Close()
	application.Database.Close()
	application.Cache.Close()
}
