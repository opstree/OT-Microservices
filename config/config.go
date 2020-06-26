package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// Configuration is struct for configuration file
type Configuration struct {
	Elasticsearch Elasticsearch `yaml:"elasticsearch"`
}

// Elasticsearch is struct for elasticsearch configuration
type Elasticsearch struct {
	Enabled  bool   `yaml:"enabled"`
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

// ParseFile parses config from a file.
func ParseFile(path string) (Configuration, error) {
	var config Configuration
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return Configuration{}, err
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return Configuration{}, err
	}
	return config, nil
}
