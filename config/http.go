package config

import "os"

type HTTPConfig struct {
	Mode       string
	Host       string
	Port       string
	ExposePort string
}

func LoadHTTPConfig() *HTTPConfig {
	return &HTTPConfig{
		Mode:       os.Getenv("GIN_MODE"),
		Host:       os.Getenv("HOST"),
		Port:       os.Getenv("PORT"),
		ExposePort: os.Getenv("EXPOSE_PORT"),
	}
}
