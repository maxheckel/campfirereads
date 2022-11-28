package config

import "github.com/kelseyhightower/envconfig"

var SmokeCostPerOrder = int64(1000)
var StandardShippingCost = int64(499)

type Config struct {
	Port                int32  `envconfig:"SERVER_PORT" required:"true"`
	GoogleAPIKey        string `envconfig:"GOOGLE_API_KEY" required:"true"`
	StripePublicAPIKey  string `envconfig:"STRIPE_PUBLIC_KEY" required:"true"`
	StripePrivateAPIKey string `envconfig:"STRIPE_PRIVATE_KEY" required:"true"`
	NYTAPIKey           string `envconfig:"NYT_API_KEY" required:"true"`
	CacheDriver         string `envconfig:"CACHE_DRIVER" required:"true"`
	CacheAddress        string `envconfig:"CACHE_ADDRESS"`
	FrontendHost        string `envconfig:"FRONTEND_HOST" default:"http://localhost:4200"`
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
