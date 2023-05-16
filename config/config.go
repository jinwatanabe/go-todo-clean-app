package config

import (
	"os"
	"strconv"
)

var config *Config

type DatabaseConfig struct {
	User string
	Password string
	Host string
	Port int
}

type Config struct {
    Database DatabaseConfig
}

func InitConfig() {
    config = &Config{
        Database: getDataConfig(),
    }
}

func GetConfig() *Config {
    return config
}

func getDataConfig() DatabaseConfig {
    port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
    return DatabaseConfig{
        User: os.Getenv("DB_USER"),
        Password: os.Getenv("DB_PASSWORD"),
        Host: os.Getenv("DB_HOST"),
        Port: port,
    }
}