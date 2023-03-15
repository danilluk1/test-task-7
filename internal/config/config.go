package config

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	PostgresUrl   string `required:"true" default:"postgres://test:test@localhost:5432/test_task_7" envconfig:"POSTGRES_URL"`
	TelegramToken string `required:"true" default:""`
	AppEnv        string `required:"true" default:"development" envconfig:"APP_ENV"`
}

func New() (*Config, error) {
	var newCfg Config
	var err error

	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	wd = filepath.Join(wd, "..", "..")
	envPath := filepath.Join(wd, ".env")
	_ = godotenv.Load(envPath)

	if err = envconfig.Process("", &newCfg); err != nil {
		return nil, err
	}

	return &newCfg, nil
}
