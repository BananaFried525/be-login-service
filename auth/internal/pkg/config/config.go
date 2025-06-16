package config

import (
	"os"

	"github.com/joho/godotenv"
)

type ServiceConfig struct {
	Name      string
	Port      string
	ApiKey    string
	JWTSecret string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

type RedisConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DB       string
}

func Init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func GetDBConfig() *DBConfig {
	return &DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}
}

func GetRedisConfig() *RedisConfig {
	return &RedisConfig{
		Host:     os.Getenv("REDIS_HOST"),
		Port:     os.Getenv("REDIS_PORT"),
		User:     os.Getenv("REDIS_USER"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       os.Getenv("REDIS_DB"),
	}
}

func GetServiceConfig() *ServiceConfig {
	return &ServiceConfig{
		Name:      os.Getenv("SERVICE_NAME"),
		Port:      os.Getenv("SERVICE_PORT"),
		ApiKey:    os.Getenv("SERVICE_API_KEY"),
		JWTSecret: os.Getenv("JWT_SECRET"),
	}
}
