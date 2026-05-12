package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort string
	DatabaseUrl string
	JWTSecret string
	JWTExpiresHour int
}

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func Load() Config {
	_ = godotenv.Load()

	expiresHour, err := strconv.Atoi(getEnv("EXPIRES_HOUR", "24"))
	if err != nil {
		expiresHour = 24
	}

	return Config{
		AppPort: getEnv("PORT", "3000"),
		DatabaseUrl: getEnv("DATABASE_URL", ""),
		JWTSecret: getEnv("JWT_SECRET", "SAUKIGANTENG"),
		JWTExpiresHour: expiresHour,
	}
}