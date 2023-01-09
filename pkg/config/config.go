package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

var ServerConfig Config

type Config struct {
	// Version is the current version of the application
	Version string `yaml:"version"`
}

func (c *Config) GetConfig() *Config {

	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		panic(err)
	}

	return c
}

func Init() error {
	ServerConfig = Config{}
	ServerConfig.GetConfig()
	return nil
}
