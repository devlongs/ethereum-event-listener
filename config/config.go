package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
    DBHost     string
    DBPort     string
    DBUser     string
    DBPassword string
    DBName     string
    EthereumURL string
}

func LoadConfig() (*Config, error) {
    err := godotenv.Load()
    if err != nil {
        return nil, err
    }

    config := &Config{
        DBHost:     os.Getenv("DB_HOST"),
        DBPort:     os.Getenv("DB_PORT"),
        DBUser:     os.Getenv("DB_USER"),
        DBPassword: os.Getenv("DB_PASSWORD"),
        DBName:     os.Getenv("DB_NAME"),
        EthereumURL: os.Getenv("ETHEREUM_URL"),
    }

    return config, nil
}