package config

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string `yaml:"env" env-required:"true"`
	DatabaseUrl string `yaml:"database_url" env-required:"true"`
}

func MustNew() *Config {
	var cfg Config

	cfgPath := os.Getenv("CONFIG_PATH")
	if cfgPath == "" {
		panic("variable CONFIG_PATH is not set")
	}

	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		panic(fmt.Sprintf("config file %s is not exists", cfgPath))
	}

	if err := cleanenv.ReadConfig(cfgPath, &cfg); err != nil {
		panic(err)
	}

	return &cfg
}
