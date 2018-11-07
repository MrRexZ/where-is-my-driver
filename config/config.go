package config

import "os"

type Config struct {
	ServerCfg *ServerCfg  `json:"serverCfg"`
	MongoCfg  *MongoDBCfg `json:"mongoCfg"`
}

type ServerCfg struct {
	HostName string `json:"hostName"`
	Port     string `json:"port"`
}

type MongoDBCfg struct {
	HostName string `json:"hostName"`
	DbName   string `json:"dbName"`
}

func GetConfig() *Config {
	return &Config{
		ServerCfg: &ServerCfg{
			HostName: envValElseDefault("gojek_1st:server:hostname", "http://localhost"),
			Port:     envValElseDefault("gojek_1st:server:port", "8080")},
		MongoCfg: &MongoDBCfg{
			HostName: envValElseDefault("gojek_1st:mongo:hostname", "mongodb://mongo:27017"),
			DbName:   envValElseDefault("gojek_1st:mongo:dbname", "gojek-db")},
	}

}

func envValElseDefault(envKey string, defaultVal string) string {
	value := os.Getenv(envKey)
	if value == "" {
		return defaultVal
	}

	return value
}
