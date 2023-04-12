package common

import (
	"gopkg.in/yaml.v3"
	"os"
	"path"
)

type serverConfig struct {
	Port      int      `yaml:"port"`
	Token     string   `yaml:"token"`
	Whitelist []string `yaml:"whitelist"`
}

type clientConfig struct {
	Host  string `yaml:"host"`
	Port  int    `yaml:"port"`
	Token string `yaml:"token"`
}

var ServerConfig = serverConfig{}
var ClientConfig = clientConfig{}

func getConfigPath(isServer bool) string {
	var configPath string
	if isServer {
		configPath = path.Join(*ConfigPath, "go-public-server.yaml")
	} else {
		configPath = path.Join(*ConfigPath, "go-public-client.yaml")
	}
	return configPath
}

func InitConfigFile(isServer bool) {
	configPath := getConfigPath(isServer)
	if _, err := os.Stat(configPath); err == nil {
		println("Config file already exists.")
		os.Exit(1)
	}
	defaultServerConfig := serverConfig{
		Port:  6871,
		Token: GenerateToken(),
	}
	defaultClientConfig := clientConfig{
		Host:  "replace_this_with_your_server_host",
		Port:  6871,
		Token: "replace_this_with_your_server_token",
	}
	var defaultConfigBytes []byte
	var err error
	if isServer {
		defaultConfigBytes, err = yaml.Marshal(defaultServerConfig)
	} else {
		defaultConfigBytes, err = yaml.Marshal(defaultClientConfig)
	}
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	err = os.WriteFile(configPath, defaultConfigBytes, 0644)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	println("Config file initialized at: " + configPath)
}

func LoadConfigFile(isServer bool) {
	configPath := getConfigPath(isServer)
	configBytes, err := os.ReadFile(configPath)
	if err != nil {
		println("Config file `" + configPath + "` not found.")
		if isServer {
			println("Use `go-public init server` to initialize a config file.")
		} else {
			println("Use `go-public init client` to initialize a config file.")
		}
		os.Exit(1)
	}
	if isServer {
		err = yaml.Unmarshal(configBytes, &ServerConfig)
	} else {
		err = yaml.Unmarshal(configBytes, &ClientConfig)
	}
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
}
