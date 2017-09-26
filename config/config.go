package config

import (
	"io/ioutil"
	"github.com/olebedev/config"
	"fmt"
)

var m_Config *Config = nil

type Config struct {
	mainConfig *config.Config
}

func NewConfig() *Config {
	configObj := new(Config)

	json, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println("Config not found")
		return nil
	} else {
		var err error
		configObj.mainConfig, err = config.ParseJson(string(json))
		if err != nil {
			fmt.Println("config.json : " + err.Error())
			return nil
		}
	}
	return configObj
}


func GetConfigInstance() *Config {
	if (m_Config != nil) {
		return m_Config
	} else {
		m_Config = NewConfig()
		return m_Config
	}
}

func (conf *Config) GetRootWWW(def string) string {
	root, err := conf.mainConfig.String("server.root")
	if (err != nil) {
		return def
	}
	return root
}