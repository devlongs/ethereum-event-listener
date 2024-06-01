package config

import (
	"os"

	"io/ioutil"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
)

type Config struct {
	DBType      string `yaml:"db_type"`
	DBHost      string `yaml:"db_host"`
	DBPort      string `yaml:"db_port"`
	DBUser      string `yaml:"db_user"`
	DBPassword  string `yaml:"db_password"`
	DBName      string `yaml:"db_name"`
	EthereumURL string `yaml:"ethereum_url"`
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	config := &Config{
		DBType:      os.Getenv("DB_TYPE"),
		DBHost:      os.Getenv("DB_HOST"),
		DBPort:      os.Getenv("DB_PORT"),
		DBUser:      os.Getenv("DB_USER"),
		DBPassword:  os.Getenv("DB_PASSWORD"),
		DBName:      os.Getenv("DB_NAME"),
		EthereumURL: os.Getenv("ETHEREUM_URL"),
	}

	return config, nil
}

func LoadYAMLConfig(filepath string) (*Config, error) {
	var config Config
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
