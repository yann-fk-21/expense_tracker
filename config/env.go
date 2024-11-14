package config

import "os"

type Env struct {
	DBUser     string
	DBPass     string
	DBName     string
	DBAddr     string
	ServerPort string
}

func InitConfig() *Env {
	return &Env{
		DBUser:     getEnv("DB_USER", "root"),
		DBPass:     getEnv("DB_PASS", "root"),
		DBName:     getEnv("DB_PASS", "Go_expense"),
		DBAddr:     getEnv("localhost", "127.0.0.1:3306"),
		ServerPort: getEnv("PORT", ":8080"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
