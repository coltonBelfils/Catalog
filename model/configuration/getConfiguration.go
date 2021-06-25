package configuration

import (
	"Catalog/niceErrors"
	"encoding/json"
	"io/ioutil"
)

type Configuration struct {
	Sql struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Address  string `json:"address"`
		Port     string `json:"port"`
		DbName   string `json:"db_name"`
	}
	Environment string `json:"environment"`
	Version     string `json:"version"`
	Google      struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
}

var conf *Configuration

func GetConfiguration() (Configuration, *niceErrors.NiceErrors) {
	if conf == nil {
		jsonData, err := ioutil.ReadFile("configuration.json")
		if err != nil {
			return Configuration{}, niceErrors.FromErrorFull(err, "Cannot read configuration.json", "-", niceErrors.ConfigurationError, niceErrors.FATAL)
		}

		conf = &Configuration{}
		err = json.Unmarshal(jsonData, conf)
		if err != nil {
			return Configuration{}, niceErrors.FromErrorFull(err, "Cannot parse configuration.json", "-", niceErrors.ConfigurationError, niceErrors.FATAL)
		}
	}
	return *conf, nil
}
