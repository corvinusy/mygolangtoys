package core

type ConfigurationDatabase struct {
	Host     string `json:"host"`
	Dbname   string `json:"dbname"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Sslmode  string `json:"sslmode"`
}

type ConfigurationStorage struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	Db       int    `json:"db"`
}

type ConfigurationCache struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	Db       int    `json:"db"`
}

type Configuration struct {
	Secret     string `json:"secret"`
	ApiVersion string `json:"ApiVersion"`
	Database   ConfigurationDatabase
	Storage    ConfigurationStorage
	Cache      ConfigurationCache
}
