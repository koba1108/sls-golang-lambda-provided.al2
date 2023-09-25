package config

import (
	"github.com/caarlos0/env/v9"
	"github.com/jinzhu/configor"
)

type AppConfig struct {
	AppEnv    string `env:"APP_ENV"`
	NestedEnv string `env:"NESTED_NAME"`
}

func LoadConfig() (*AppConfig, error) {
	var cfg AppConfig
	err := configor.Load(&cfg, "src/config/config.yml")
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func LoadConfigFromENV() (*AppConfig, error) {
	var conf AppConfig
	err := env.ParseWithOptions(&conf, env.Options{
		RequiredIfNoDef: true,
	})
	if err != nil {
		return nil, err
	}
	return &conf, nil
}
