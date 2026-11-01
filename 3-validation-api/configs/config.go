package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Db          DbConfig    `json:"db"`
	Auth        AuthConfig  `json:"auth"`
	EmailConfig EmailConfig `json:"email_config"`
}

type AuthConfig struct {
	Secret string `json:"token"`
}

type EmailConfig struct {
	Email    string `json:"verify"`
	Password string `json:"password"`
	Address  string `json:"address"`
}
type DbConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
	DSN  string `json:"dsn"`
}

func GetConfig() Config {
	// godotenv.Load() без аргументов автоматически ищет файл .env в текущей директории
	err := godotenv.Load()
	if err != nil {
		// Не игнорируйте ошибку! Как минимум, запишите её в лог.
		// Программа продолжит работу, пытаясь взять переменные из среды ОС.
		log.Println("⚠️ Warning: .env file not found, using system environment variables or empty values")
	}

	return Config{
		Db: DbConfig{
			Host: os.Getenv("DB_HOST"),
			Port: os.Getenv("DB_PORT"),
			DSN:  os.Getenv("DB_DSN"),
		},
		Auth: AuthConfig{
			Secret: os.Getenv("AUTH_TOKEN"),
		},
		EmailConfig: EmailConfig{
			Email:    os.Getenv("EMAIL_ADDRESS"),
			Password: os.Getenv("EMAIL_PASSWORD"),
			Address:  os.Getenv("EMAIL_ADDRESS"),
		},
	}
}
