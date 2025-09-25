package config

import (
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

const defaultConfigPath = "./configs/default.yaml"

type Config struct {
	Host           string        `yaml:"host"`
	Port           string        `yaml:"port"`
	ReadTimeout    time.Duration `yaml:"read_timeout"`
	WriteTimeout   time.Duration `yaml:"write_timeout"`
	ContextTimeout time.Duration `yaml:"context_timeout"`
}

func MustLoadConfig() Config {
	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		configPath = defaultConfigPath
	}

	var config Config

	cleanenv.ReadConfig(configPath, &config)

	return config
}
