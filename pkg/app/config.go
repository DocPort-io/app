package app

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

// Config represents the application configuration loaded via Viper.
type Config struct {
	Database struct {
		Driver string `mapstructure:"driver" validate:"required"`
		URL    string `mapstructure:"url" validate:"required"`
	} `mapstructure:"database"`

	Server struct {
		Host string `mapstructure:"host"`
		Bind string `mapstructure:"bind" validate:"required"`
		Port string `mapstructure:"port" validate:"required"`
	} `mapstructure:"server"`

	Storage struct {
		Provider string `mapstructure:"provider" validate:"required"`
		Path     string `mapstructure:"path" validate:"required"`
	} `mapstructure:"storage"`
}

// LoadConfig reads configuration using Viper, unmarshals into a struct, and validates it.
func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("/etc/docport/")
	viper.AddConfigPath("$HOME/.docport")
	viper.AddConfigPath(".")

	viper.SetEnvPrefix("docport")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	if err := validateConfig(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func validateConfig(cfg *Config) error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(cfg); err != nil {
		return fmt.Errorf("invalid configuration: %w", err)
	}

	return nil
}
