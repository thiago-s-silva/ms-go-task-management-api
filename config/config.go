package config

import (
	"fmt"

	"github.com/gofor-little/env"
)

type Config struct {
	SERVER_HOST      string
	API_V1_BASE_PATH string
	API_V2_BASE_PATH string
}

func Load() (*Config, error) {
	// Load an .env file and set the key-value pairs as environment variables.
	if err := env.Load(".env"); err != nil {
		return nil, err
	}

	// load server config variable values
	serverIp := env.Get("SERVER_IP", "localhost")
	serverPort := env.Get("SERVER_PORT", "8080")

	c := Config{
		SERVER_HOST:      fmt.Sprintf("%s:%s", serverIp, serverPort),
		API_V1_BASE_PATH: "/api/v1",
		API_V2_BASE_PATH: "/api/v2",
	}

	return &c, nil
}
