package common

import (
	"fmt"
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
	// If config file is not found in the specified path, try to find it in the default path
	if _, err := os.Stat(*ConfigPath); os.IsNotExist(err) {
		basePath := path.Join(GetHomeDir(), ".config")
		var configPath2 string
		if isServer {
			configPath2 = path.Join(basePath, "go-public-server.yaml")
		} else {
			configPath2 = path.Join(basePath, "go-public-client.yaml")
		}
		if _, err := os.Stat(configPath2); err == nil {
			configPath = configPath2
		}
	}
	return configPath
}

func InitConfigFile(isServer bool) {
	configPath := getConfigPath(isServer)
	if _, err := os.Stat(configPath); err == nil {
		fmt.Println("Config file already exists at: " + configPath)
		os.Exit(1)
	}
	if !isServer {
		fmt.Print("Where do you want to save the config file? (default: ~/.config) ")
		var input string
		_, _ = fmt.Scanln(&input)
		if input == "" {
			input = path.Join(GetHomeDir(), ".config")
		}
		if _, err := os.Stat(input); os.IsNotExist(err) {
			err := os.Mkdir(input, 0700)
			if err != nil {
				fmt.Println("Failed to create directory: " + input)
				os.Exit(1)
			}
		}
		configPath = path.Join(input, "go-public-client.yaml")
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
			println("Generating a new config file...")
			InitConfigFile(true)
			configBytes, _ = os.ReadFile(configPath)
		} else {
			println("Use `go-public init client` to initialize a config file.")
			os.Exit(1)
		}
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
