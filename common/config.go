package common

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
}

var CONFIG = Config{}

func initConfigFile() {
	if _, err := os.Stat(*ConfigFile); err == nil {
		println("Config file already exists.")
		os.Exit(1)
	}
	defaultConfig := Config{
		Host:     "localhost",
		Port:     6971,
		Password: "123456",
	}
	defaultConfigBytes, err := yaml.Marshal(defaultConfig)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	err = os.WriteFile(*ConfigFile, defaultConfigBytes, 0644)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	println("Config file created successfully, file path: " + *ConfigFile)
}

func loadConfigFile() {
	configBytes, err := os.ReadFile(*ConfigFile)
	if err != nil {
		println("Config file `" + *ConfigFile + "` not found, use subcommand `init` to create a new one.")
		os.Exit(1)
	}
	err = yaml.Unmarshal(configBytes, &CONFIG)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
}
