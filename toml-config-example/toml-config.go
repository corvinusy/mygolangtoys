package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"code.dm.com/CloudOffice-Server/ls/sl-common"

	"github.com/BurntSushi/toml"
	"github.com/jinzhu/gorm"
	"github.com/t-k/fluent-logger-golang/fluent"
)

type Email struct {
	Server   string `toml:"server"`
	Port     string `toml:"port"`
	Name     string `toml:"name"`
	Address  string `toml:"address"`
	Password string `toml:"password"`
	Title    string `toml:"title"`
}

type appConfig struct {
	Secret       string `toml:"secret"`
	Version      string `toml:"version"`
	Port         string `toml:"port"`
	Db           string `toml:"db"`
	Dsn          string `toml:"dsn"`
	TestDb       string `toml:"test_db"`
	TestDsn      string `toml:"test_dsn"`
	LeaseTimeAdd string `toml:"lease_time_add"`
	Swagger      struct {
		URL         string `toml:"url"`
		APIPath     string `toml:"api_path"`
		SwaggerPath string `toml:"swagger_path"`
		SwaggerDist string `toml:"swagger_dist"`
	} `toml:"swagger"`
	Fluent struct {
		Port  string `toml:"port"`
		Host  string `toml:"host"`
		Tag   string `toml:"tag"`
		DbTag string `toml:"db_tag"`
	} `toml:"fluent"`
	Email sl.Email `toml:"email"`
}

type appEnvironment struct {
	Config *appConfig
	DB     *gorm.DB
	Logger *fluent.Fluent
}

const (
	cfgFileName = "./default.config.toml"
)

func main() {
	// read config file and unmarshal it
	tomlData, err := ioutil.ReadFile(cfgFileName)
	if err != nil {
		panic(err)
	}

	var config appConfig
	metaData, err := toml.Decode(string(tomlData[:]), &config)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", config)
	fmt.Printf("%+v\n", metaData.Undecoded())
	os.Exit(1)
}
