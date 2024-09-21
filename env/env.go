package env

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func GetEnv(key string) string {
	return os.Getenv(key)
}

func GetDatabaseURL() string {
	env := GetEnv("DATABASE_URL")

	if env == "" {
		panic("DATABASE_URL is not set")
	}

	return env
}
