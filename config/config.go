package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
	"os"
)

type config struct {
	MongoUrlConnection string `yaml:"mongo_url"`
	Addr               string `yaml:"port"`
}

func NewConfigStruct() *config {
	return &config{}
}

func (c *config) LoadConfig(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("Failed open config file :%v", err)
	}
	defer file.Close()
	read, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("Failed reading file :%v", err)
	}
	err = yaml.Unmarshal(read, c)
	if err != nil {
		return fmt.Errorf("Yaml unmarshalling error :%v", err)
	}

	return nil
}
