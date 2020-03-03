package config

import "os"

type Config struct {
	Server  string
	Subject string
}

func GetConfig() (Config, error) {
	cfg := Config{}

	cfg.Server = os.Getenv("NATS_SERVER")
	if cfg.Server == "" {
		cfg.Server = "nats://127.0.0.1:4222"
	}

	cfg.Subject = os.Getenv("NATS_SUBJECT")
	if cfg.Subject == "" {
		cfg.Subject = ">"
	}

	return cfg, nil
}
