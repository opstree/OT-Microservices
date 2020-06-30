package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// Configuration is struct for configuration file
type Configuration struct {
	Elasticsearch Elasticsearch `yaml:"elasticsearch"`
	Management    Management    `yaml:"management"`
	Attendance    Attendance    `yaml:"attendance"`
	MySQL         MySQL         `yaml:"mysql"`
	Salary        Salary        `yaml:"salary"`
}

// Elasticsearch is struct for elasticsearch configuration
type Elasticsearch struct {
	Enabled  bool   `yaml:"enabled"`
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

// Management is struct for management API config
type Management struct {
	API     string `yaml:"api_endpoint"`
	APIPort string `yaml:"api_port"`
}

// Attendance is struct for management API config
type Attendance struct {
	API     string `yaml:"api_endpoint"`
	APIPort string `yaml:"api_port"`
}

// Salary is struct for management API config
type Salary struct {
	API     string `yaml:"api_endpoint"`
	APIPort string `yaml:"api_port"`
}

// MySQL will the struct for MYSQL Config
type MySQL struct {
	Enabled  bool   `yaml:"enabled"`
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"db_name"`
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
