package config

import "os"

type Config struct {
	AddrStr string `json:"AddrStr"`
}

func New() *Config {
	return &Config{}
}

func (c *Config) Read() {
	httpAddr, exists := os.LookupEnv("HTTP_ADDR")
	if exists {
		c.AddrStr = httpAddr
	}
}
