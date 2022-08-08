package config

import "fmt"

type Config struct {
	Environment     string `env:"ENVIRONMENT"`
	LogLevel        string `env:"LOG_LEVEL"`
	ServerIP        string `env:"SERVER_IP"`
	GRPCPort        string `env:"GRPC_PORT"`
	HTTPPort        string `env:"HTTP_PORT"`
	ServiceName     string `env:"SERVICE_NAME"`
	OpenApiURL      string `env:"OPEN_API_URL"`
	GRPCPostService string `env:"GRPC_POST_SERVICE"`

	PostgresHost     string `env:"POSTGRES_HOST"`
	PostgresPort     string `env:"POSTGRES_PORT"`
	PostgresDatabase string `env:"POSTGRES_DATABASE"`
	PostgresUser     string `env:"POSTGRES_USER"`
	PostgresPassword string `env:"POSTGRES_PASSWORD"`
	PostgresSSLMode  string `env:"POSTGRES_SSLMODE"`
}

func (c *Config) PostgresURI() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.PostgresHost,
		c.PostgresPort,
		c.PostgresUser,
		c.PostgresPassword,
		c.PostgresDatabase,
		c.PostgresSSLMode)
}
