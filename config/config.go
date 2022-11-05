package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"os"
)

type conf struct {
	Addr  string `yaml:"port"`
	DbUrl string `yaml:"db_url"`
}

func NewConfigStruct() *conf {
	return &conf{}
}

func (c *conf) LoadConfig(fileName string) error {
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
