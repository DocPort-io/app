package config

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig   `mapstructure:"server" validate:"required"`
	Database DatabaseConfig `mapstructure:"database" validate:"required"`
	Auth     AuthConfig     `mapstructure:"auth" validate:"required"`
	Storage  StorageConfig  `mapstructure:"storage" validate:"required"`
}

type ServerConfig struct {
	Bind string `mapstructure:"bind" validate:"required"`
	Port int    `mapstructure:"port" validate:"required,min=1,max=65535"`
}

type DatabaseConfig struct {
	DSN string `mapstructure:"dsn" validate:"required"`
}

type AuthConfig struct {
	JWKSUrl string   `mapstructure:"jwks_url" validate:"required"`
	Scopes  []string `mapstructure:"scopes"`
}

type StorageConfig struct {
	Provider string `mapstructure:"provider" validate:"required"`
	Path     string `mapstructure:"path" validate:"required"`
}

func Load() (Config, error) {
	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("toml")
	v.AddConfigPath("/etc/docport/")
	v.AddConfigPath("$HOME/.docport")
	v.AddConfigPath(".")

	v.SetDefault("server.bind", "0.0.0.0")
	v.SetDefault("server.port", 8080)
	v.SetDefault("database.dsn", "postgres://postgres:postgres@localhost:5432/docport?sslmode=disable")
	v.SetDefault("auth.jwks_url", "https://keycloak/realms/docport/protocol/openid-connect/certs")
	v.SetDefault("auth.scopes", []string{})
	v.SetDefault("storage.provider", "filesystem")
	v.SetDefault("storage.path", "./storage")

	v.SetEnvPrefix("docport")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		var notFound viper.ConfigFileNotFoundError
		if !errors.As(err, &notFound) {
			return Config{}, fmt.Errorf("error reading config file: %w", err)
		}
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return Config{}, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	validate := validator.New()

	if err := validate.Struct(&cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
