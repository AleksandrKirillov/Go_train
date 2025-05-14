package config

import "os"

type Config struct {
	Key string
}

func NewConfig() *Config {
	key := os.Getenv("KEY")
	if key == "" {
		panic("KEY not set in environment variables")
	}

	return &Config{
		Key: key,
	}
}
