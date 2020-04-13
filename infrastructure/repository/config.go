package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Config struct for config
type Config struct {
	DatabaseInfo *DatabaseInfo `json:"DatabaseInfo"`
}

// DatabaseInfo struct for database info
type DatabaseInfo struct {
	Host                   string `json:"Host"`
	EmployerDatabaseName   string `json:"EmployerDatabaseName"`
	AgencyDatabaseName     string `json:"AgencyDatabaseName"`
	AuctionDatabaseName    string `json:"AuctionDatabaseName"`
	EngagementDatabaseName string `json:"EngagementDatabaseName"`
	Port                   string `json:"Port"`
	Username               string `json:"Username"`
	Password               string `json:"Password"`
}

// NewConfig return new instance of Config
func NewConfig() *Config {
	localFile := "./config.json"

	if fileExists(localFile) {
		var config *Config
		value, err := ioutil.ReadFile(localFile)

		if err != nil {
			fmt.Println(err)
		}

		json.Unmarshal(value, &config)

		return config
	}

	configfile := "/run/secrets/config.json"

	if fileExists(configfile) {
		var config *Config
		value, err := ioutil.ReadFile(configfile)

		if err != nil {
			fmt.Println(err)
		}

		json.Unmarshal(value, &config)

		return config
	}
	return nil
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)

	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
