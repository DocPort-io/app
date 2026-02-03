package config

import (
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
	Host string `mapstructure:"host" validate:"required"`
	Bind string `mapstructure:"bind" validate:"required"`
	Port string `mapstructure:"port" validate:"required"`
}

type DatabaseConfig struct {
	DSN string `mapstructure:"dsn" validate:"required"`
}

type AuthConfig struct {
	Issuer       string `mapstructure:"issuer" validate:"required"`
	ClientId     string `mapstructure:"clientId" validate:"required"`
	ClientSecret string `mapstructure:"clientSecret" validate:"required"`
}

type StorageConfig struct {
	Provider string `mapstructure:"provider" validate:"required"`
	Path     string `mapstructure:"path" validate:"required"`
}

func Load() (Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("/etc/docport/")
	viper.AddConfigPath("$HOME/.docport")
	viper.AddConfigPath(".")

	viper.SetEnvPrefix("docport")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return Config{}, fmt.Errorf("error reading config file: %w", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return Config{}, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	validate := validator.New()

	if err := validate.Struct(&cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
