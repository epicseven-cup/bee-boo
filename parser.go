package beeboo

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const (
	configFilePath = "config/config.yaml"
	data           = "data/"
)

type Config struct {
	service map[string]interface{} `yaml:service`
}

type ServiceConfig struct {
	name string
}

func parser() {
	file, err := os.ReadFile(configFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	cfg := Config{}
	err = yaml.Unmarshal(file, &cfg)

	// Going through each service
	for serviceConfig := range cfg.service {
		port, ok := serviceConfig["port"]
	}
}

type Service struct {
	Port    int
	Streams map[string]string
}
