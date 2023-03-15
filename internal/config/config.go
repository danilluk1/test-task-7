package config

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	PostgresUrl   string `required:"true" default:"postgres://postgres:admin@localhost:5432/test_task_7" envconfig:"POSTGRES_URL"`
	TelegramToken string `required:"true" default:"" envconfig:"TELEGRAM_TOKEN"`
	AppEnv        string `required:"true" default:"development" envconfig:"APP_ENV"`
	WeatherApiKey string `required:"true" default:"" envconfig:"WEATHER_API_KEY"`
}

func New() (*Config, error) {
	var newCfg Config
	var err error

	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	envPath := filepath.Join(wd, ".env")
	_ = godotenv.Load(envPath)

	if err = envconfig.Process("", &newCfg); err != nil {
		return nil, err
	}

	return &newCfg, nil
}
