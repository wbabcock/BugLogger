package configuration

import (
	"encoding/json"
	"os"
)

// Database - Struct for the config file
type Database struct {
	Type     string `json:"type"`
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
	Port     int    `json:"port"`
	Database string `json:"database"`
}

// Config - Struct for the config file
type Config struct {
	API struct {
		Port            int      `json:"port"`
		AdminUser       string   `json:"adminUser"`
		AdminPassword   string   `json:"adminPassword"`
		AllowAllOrigins bool     `json:"allowAllOrigins"`
		AllowOrigins    []string `json:"allowOrigins"`
		AllowMethods    []string `json:"allowMethods"`
		AllowHeaders    []string `json:"alloHeaders"`
	}
	Databases []Database
}

var _config Config

// LoadConfig - Get the configuration data for the API
func LoadConfig(file string) (Config, error) {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		return config, err
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	_config = config
	return config, err
}

// GetConfig - Get the configuration
func GetConfig() Config {
	return _config
}

// GetDatabaseConfig - Get the configuration
func GetDatabaseConfig(name string) Database {
	var database Database

	for _, dbConfig := range _config.Databases {
		if name == dbConfig.Database {
			database = dbConfig
			return database
		}
	}

	return database
}
