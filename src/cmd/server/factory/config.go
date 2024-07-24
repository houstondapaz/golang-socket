package factory

import "github.com/caarlos0/env/v11"

type ServerConfig struct {
	Port      string `env:"API_PORT" envDefault:"3000"`
	BaseRoute string `env:"API_BASE_ROUTE", envDefault:"/"`
}

func NewApiConfig() (*ServerConfig, error) {
	config := &ServerConfig{}

	if err := env.Parse(config); err != nil {
		return nil, err
	}

	return config, nil
}
