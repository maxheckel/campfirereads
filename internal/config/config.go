package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Port         int32  `envconfig:"SERVER_PORT" required:"true"`
	GoogleAPIKey string `envconfig:"GOOGLE_API_KEY" required:"true"`
	NYTAPIKey    string `envconfig:"NYT_API_KEY" required:"true"`
	CacheDriver  string `envconfig:"CACHE_DRIVER" required:"true"`
	CacheAddress string `envconfig:"CACHE_ADDRESS"`
}

func Load(prefix string) (*Config, error) {
	cfg := &Config{
		Port: 8080,
	}
	err := envconfig.Process(prefix, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
