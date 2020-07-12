package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// Configuration is struct for configuration file
type Configuration struct {
	Notification Notification `yaml:"notification"`
	SMTP         SMTP         `yaml:"smtp"`
}

// SMTP is struct for smtp configuration
type SMTP struct {
	From       string `yaml:"from"`
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
	SMTPServer string `yaml:"smtp_server"`
	SMTPPort   string `yaml:"smtp_port"`
}

// Notification is struct for Notification API config
type Notification struct {
	APIPort string `yaml:"api_port"`
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
