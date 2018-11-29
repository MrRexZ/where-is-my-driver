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
			HostName: envValElseDefault(SERVER_HOST_ENV_KEY, "http://localhost"),
			Port:     envValElseDefault(SERVER_PORT_ENV_KEY, "8080")},
		MongoCfg: &MongoDBCfg{
			HostName: envValElseDefault(MONGO_HOST_ENV_KEY, "mongodb://localhost:27017"),
			DbName:   envValElseDefault(MONGO_DB_ENV_KEY, "gojek-db")},
	}

}

func envValElseDefault(envKey string, defaultVal string) string {
	value := os.Getenv(envKey)
	if value == "" {
		return defaultVal
	}

	return value
}
