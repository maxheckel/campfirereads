package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Port         int32
	GoogleAPIKey string `envconfig:"GOOGLE_API_KEY" required:"true"`
}

func Load(prefix string) (*Config, error) {
	cfg := &Config{}
	err := envconfig.Process(prefix, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
