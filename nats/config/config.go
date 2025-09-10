package config

type Config struct {
	NatsURL string
}

func NewConfig() *Config {
	return &Config{
		NatsURL: "nats://localhost:4222",
	}
}
