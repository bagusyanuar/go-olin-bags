package config

import (
	"errors"
	"fmt"

	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
)

type Config struct {
	Host string `env:"SERVER_HOST,default=localhost"`
	Port string `env:"SERVER_PORT"`
}

func NewConfig(env string) (*Config, error) {
	_ = godotenv.Load(env)

	var config Config
	if err := envdecode.Decode(&config); err != nil {
		message := fmt.Sprintf("error load %s file", env)
		return nil, errors.New(message)
	}
	return &config, nil
}
