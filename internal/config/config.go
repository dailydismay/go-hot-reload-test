package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	Port string `yaml:"port"`
}

func NewConfig() (*Config, error) {
	var c Config

	err := cleanenv.ReadConfig("config/config.yml", &c)

	if err != nil {
		return nil, err
	}

	return &c, nil
}
