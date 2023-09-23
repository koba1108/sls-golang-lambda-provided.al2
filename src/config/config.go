package config

import "github.com/jinzhu/configor"

type AppConfig struct {
	Env string `yaml:"env"`
}

func LoadConfig() (*AppConfig, error) {
	var cfg AppConfig
	err := configor.Load(&cfg, "src/config/config.yml")
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
